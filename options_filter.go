package libp2p

import (
	"github.com/riteshRcH/core/connmgr"
	"github.com/riteshRcH/core/control"
	"github.com/riteshRcH/core/network"
	"github.com/riteshRcH/core/peer"

	ma "github.com/riteshRcH/multiaddr"
)

// filtersConnectionGater is an adapter that turns multiaddr.Filter into a
// connmgr.ConnectionGater.
type filtersConnectionGater ma.Filters

var _ connmgr.ConnectionGater = (*filtersConnectionGater)(nil)

func (f *filtersConnectionGater) InterceptAddrDial(_ peer.ID, addr ma.Multiaddr) (allow bool) {
	return !(*ma.Filters)(f).AddrBlocked(addr)
}

func (f *filtersConnectionGater) InterceptPeerDial(p peer.ID) (allow bool) {
	return true
}

func (f *filtersConnectionGater) InterceptAccept(connAddr network.ConnMultiaddrs) (allow bool) {
	return !(*ma.Filters)(f).AddrBlocked(connAddr.RemoteMultiaddr())
}

func (f *filtersConnectionGater) InterceptSecured(_ network.Direction, _ peer.ID, connAddr network.ConnMultiaddrs) (allow bool) {
	return !(*ma.Filters)(f).AddrBlocked(connAddr.RemoteMultiaddr())
}

func (f *filtersConnectionGater) InterceptUpgraded(_ network.Conn) (allow bool, reason control.DisconnectReason) {
	return true, 0
}
