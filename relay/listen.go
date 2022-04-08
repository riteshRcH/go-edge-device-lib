package relay

import (
	"net"

	pb "github.com/riteshRcH/go-edge-device-lib/relay/pb"

	ma "github.com/riteshRcH/go-edge-device-lib/multiaddr"
	manet "github.com/riteshRcH/go-edge-device-lib/multiaddr/net"
)

var _ manet.Listener = (*RelayListener)(nil)

type RelayListener Relay

func (l *RelayListener) Relay() *Relay {
	return (*Relay)(l)
}

func (r *Relay) Listener() *RelayListener {
	// TODO: Only allow one!
	return (*RelayListener)(r)
}

func (l *RelayListener) Accept() (manet.Conn, error) {
	for {
		select {
		case c := <-l.incoming:
			err := l.Relay().writeResponse(c.stream, pb.CircuitRelay_SUCCESS)
			if err != nil {
				log.Debugf("error writing relay response: %s", err.Error())
				c.stream.Reset()
				continue
			}

			// TODO: Pretty print.
			log.Infof("accepted relay connection: %q", c)

			c.tagHop()
			return c, nil
		case <-l.ctx.Done():
			return nil, l.ctx.Err()
		}
	}
}

func (l *RelayListener) Addr() net.Addr {
	return &NetAddr{
		Relay:  "any",
		Remote: "any",
	}
}

func (l *RelayListener) Multiaddr() ma.Multiaddr {
	return circuitAddr
}

func (l *RelayListener) Close() error {
	// TODO: noop?
	return nil
}
