package relay

import (
	"context"
	"fmt"

	"github.com/riteshRcH/go-edge-device-lib/core/network"

	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	"github.com/riteshRcH/go-edge-device-lib/core/transport"
	ma "github.com/riteshRcH/go-edge-device-lib/multiaddr"
)

func (d *RelayTransport) Dial(ctx context.Context, a ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	c, err := d.Relay().Dial(ctx, a, p)
	if err != nil {
		return nil, err
	}
	c.tagHop()
	scope, _ := network.NullResourceManager.OpenConnection(network.DirOutbound, false)
	return d.upgrader.Upgrade(ctx, d, c, network.DirOutbound, p, scope)
}

func (r *Relay) Dial(ctx context.Context, a ma.Multiaddr, p peer.ID) (*Conn, error) {
	// split /a/p2p-circuit/b into (/a, /p2p-circuit/b)
	relayaddr, destaddr := ma.SplitFunc(a, func(c ma.Component) bool {
		return c.Protocol().Code == ma.P_CIRCUIT
	})

	// If the address contained no /p2p-circuit part, the second part is nil.
	if destaddr == nil {
		return nil, fmt.Errorf("%s is not a relay address", a)
	}

	if relayaddr == nil {
		return nil, fmt.Errorf(
			"can't dial a p2p-circuit without specifying a relay: %s",
			a,
		)
	}

	// Strip the /p2p-circuit prefix from the destaddr.
	_, destaddr = ma.SplitFirst(destaddr)

	dinfo := &peer.AddrInfo{ID: p, Addrs: []ma.Multiaddr{}}
	if destaddr != nil {
		dinfo.Addrs = append(dinfo.Addrs, destaddr)
	}

	var rinfo *peer.AddrInfo
	rinfo, err := peer.AddrInfoFromP2pAddr(relayaddr)
	if err != nil {
		return nil, fmt.Errorf("error parsing multiaddr '%s': %s", relayaddr.String(), err)
	}

	return r.DialPeer(ctx, *rinfo, *dinfo)
}