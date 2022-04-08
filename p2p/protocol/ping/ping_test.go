package ping_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	bhost "github.com/riteshRcH/go-edge-device-lib/p2p/host/basic"
	"github.com/riteshRcH/go-edge-device-lib/p2p/protocol/ping"
	swarmt "github.com/riteshRcH/go-edge-device-lib/swarm/testing"
)

func TestPing(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	h1, err := bhost.NewHost(swarmt.GenSwarm(t), nil)
	require.NoError(t, err)
	h2, err := bhost.NewHost(swarmt.GenSwarm(t), nil)
	require.NoError(t, err)

	err = h1.Connect(ctx, peer.AddrInfo{
		ID:    h2.ID(),
		Addrs: h2.Addrs(),
	})
	require.NoError(t, err)

	ps1 := ping.NewPingService(h1)
	ps2 := ping.NewPingService(h2)

	testPing(t, ps1, h2.ID())
	testPing(t, ps2, h1.ID())
}

func testPing(t *testing.T, ps *ping.PingService, p peer.ID) {
	pctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ts := ps.Ping(pctx, p)

	for i := 0; i < 5; i++ {
		select {
		case res := <-ts:
			require.NoError(t, res.Error)
			t.Log("ping took: ", res.RTT)
		case <-time.After(time.Second * 4):
			t.Fatal("failed to receive ping")
		}
	}

}
