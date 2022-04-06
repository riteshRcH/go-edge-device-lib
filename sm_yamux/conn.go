package sm_yamux

import (
	"context"

	"github.com/riteshRcH/go-edge-device-lib/core/network"

	"github.com/riteshRcH/go-edge-device-lib/yamux"
)

// conn implements mux.MuxedConn over yamux.Session.
type conn yamux.Session

var _ network.MuxedConn = &conn{}

// Close closes underlying yamux
func (c *conn) Close() error {
	return c.yamux().Close()
}

// IsClosed checks if yamux.Session is in closed state.
func (c *conn) IsClosed() bool {
	return c.yamux().IsClosed()
}

// OpenStream creates a new stream.
func (c *conn) OpenStream(ctx context.Context) (network.MuxedStream, error) {
	s, err := c.yamux().OpenStream(ctx)
	if err != nil {
		return nil, err
	}

	return (*stream)(s), nil
}

// AcceptStream accepts a stream opened by the other side.
func (c *conn) AcceptStream() (network.MuxedStream, error) {
	s, err := c.yamux().AcceptStream()
	return (*stream)(s), err
}

func (c *conn) yamux() *yamux.Session {
	return (*yamux.Session)(c)
}
