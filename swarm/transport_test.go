package swarm_test

import (
	"context"
	"testing"

	swarm "github.com/riteshRcH/go-edge-device-lib/swarm"
	swarmt "github.com/riteshRcH/go-edge-device-lib/swarm/testing"

	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	"github.com/riteshRcH/go-edge-device-lib/core/transport"

	ma "github.com/riteshRcH/go-edge-device-lib/multiaddr"

	"github.com/stretchr/testify/require"
)

type dummyTransport struct {
	protocols []int
	proxy     bool
	closed    bool
}

func (dt *dummyTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	panic("unimplemented")
}

func (dt *dummyTransport) CanDial(addr ma.Multiaddr) bool {
	panic("unimplemented")
}

func (dt *dummyTransport) Listen(laddr ma.Multiaddr) (transport.Listener, error) {
	panic("unimplemented")
}

func (dt *dummyTransport) Proxy() bool {
	return dt.proxy
}

func (dt *dummyTransport) Protocols() []int {
	return dt.protocols
}
func (dt *dummyTransport) Close() error {
	dt.closed = true
	return nil
}

func TestUselessTransport(t *testing.T) {
	s := swarmt.GenSwarm(t)
	require.Error(t, s.AddTransport(new(dummyTransport)), "adding a transport that supports no protocols should have failed")
}

func TestTransportClose(t *testing.T) {
	s := swarmt.GenSwarm(t)
	tpt := &dummyTransport{protocols: []int{1}}
	require.NoError(t, s.AddTransport(tpt))
	_ = s.Close()
	if !tpt.closed {
		t.Fatal("expected transport to be closed")
	}
}

func TestTransportAfterClose(t *testing.T) {
	s := swarmt.GenSwarm(t)
	s.Close()

	tpt := &dummyTransport{protocols: []int{1}}
	if err := s.AddTransport(tpt); err != swarm.ErrSwarmClosed {
		t.Fatal("expected swarm closed error, got: ", err)
	}
}
