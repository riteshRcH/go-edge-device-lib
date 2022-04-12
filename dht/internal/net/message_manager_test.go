package net

import (
	"context"
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	"github.com/riteshRcH/go-edge-device-lib/core/protocol"

	bhost "github.com/riteshRcH/go-edge-device-lib/p2p/host/basic"
	swarmt "github.com/riteshRcH/go-edge-device-lib/swarm/testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidMessageSenderTracking(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	foo := peer.ID("asdasd")

	h, err := bhost.NewHost(swarmt.GenSwarm(t, swarmt.OptDisableReuseport), new(bhost.HostOpts))
	require.NoError(t, err)
	defer h.Close()

	msgSender := NewMessageSenderImpl(h, []protocol.ID{"/test/kad/1.0.0"}).(*messageSenderImpl)

	_, err = msgSender.messageSenderForPeer(ctx, foo)
	require.Error(t, err, "should have failed to find message sender")

	msgSender.smlk.Lock()
	mscnt := len(msgSender.strmap)
	msgSender.smlk.Unlock()

	if mscnt > 0 {
		t.Fatal("should have no message senders in map")
	}
}
