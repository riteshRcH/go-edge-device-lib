package identify

import (
	"context"
	"testing"
	"time"

	"github.com/riteshRcH/go-edge-device-lib/core/network"
	"github.com/riteshRcH/go-edge-device-lib/core/peer"

	blhost "github.com/riteshRcH/go-edge-device-lib/blankhost"
	swarmt "github.com/riteshRcH/go-edge-device-lib/swarm/testing"

	"github.com/stretchr/testify/require"
)

func TestMakeApplyDelta(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h1 := blhost.NewBlankHost(swarmt.GenSwarm(t))
	defer h1.Close()
	ids1, err := NewIDService(h1)
	require.NoError(t, err)
	ph := newPeerHandler(h1.ID(), ids1)
	ph.start(ctx, func() {})
	defer ph.stop()

	m1 := ph.nextDelta()
	require.NotNil(t, m1)
	// We haven't changed anything since creating the peer handler
	require.Empty(t, m1.AddedProtocols)

	h1.SetStreamHandler("p1", func(network.Stream) {})
	m2 := ph.nextDelta()
	require.Len(t, m2.AddedProtocols, 1)
	require.Contains(t, m2.AddedProtocols, "p1")
	require.Empty(t, m2.RmProtocols)

	h1.SetStreamHandler("p2", func(network.Stream) {})
	h1.SetStreamHandler("p3", func(stream network.Stream) {})
	m3 := ph.nextDelta()
	require.Len(t, m3.AddedProtocols, 2)
	require.Contains(t, m3.AddedProtocols, "p2")
	require.Contains(t, m3.AddedProtocols, "p3")
	require.Empty(t, m3.RmProtocols)

	h1.RemoveStreamHandler("p3")
	m4 := ph.nextDelta()
	require.Empty(t, m4.AddedProtocols)
	require.Len(t, m4.RmProtocols, 1)
	require.Contains(t, m4.RmProtocols, "p3")

	h1.RemoveStreamHandler("p2")
	h1.RemoveStreamHandler("p1")
	m5 := ph.nextDelta()
	require.Empty(t, m5.AddedProtocols)
	require.Len(t, m5.RmProtocols, 2)
	require.Contains(t, m5.RmProtocols, "p2")
	require.Contains(t, m5.RmProtocols, "p1")
}

func TestHandlerClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h1 := blhost.NewBlankHost(swarmt.GenSwarm(t))
	defer h1.Close()
	ids1, err := NewIDService(h1)
	require.NoError(t, err)
	ph := newPeerHandler(h1.ID(), ids1)
	closedCh := make(chan struct{}, 2)
	ph.start(ctx, func() {
		closedCh <- struct{}{}
	})

	require.NoError(t, ph.stop())
	select {
	case <-closedCh:
	case <-time.After(time.Second):
		t.Fatal("expected the handler to close")
	}

	require.NoError(t, ph.stop())
	select {
	case <-closedCh:
		t.Fatal("expected only one close event")
	case <-time.After(10 * time.Millisecond):
	}
}

func TestPeerSupportsProto(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h1 := blhost.NewBlankHost(swarmt.GenSwarm(t))
	defer h1.Close()
	ids1, err := NewIDService(h1)
	require.NoError(t, err)

	rp := peer.ID("test")
	ph := newPeerHandler(rp, ids1)
	require.NoError(t, h1.Peerstore().AddProtocols(rp, "test"))
	require.True(t, ph.peerSupportsProtos(ctx, []string{"test"}))
	require.False(t, ph.peerSupportsProtos(ctx, []string{"random"}))

	// remove support for protocol and check
	require.NoError(t, h1.Peerstore().RemoveProtocols(rp, "test"))
	require.False(t, ph.peerSupportsProtos(ctx, []string{"test"}))
}
