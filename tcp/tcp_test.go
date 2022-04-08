package tcp

import (
	"context"
	"errors"
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/core/network"

	"github.com/riteshRcH/go-edge-device-lib/core/crypto"
	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	"github.com/riteshRcH/go-edge-device-lib/core/sec"
	"github.com/riteshRcH/go-edge-device-lib/core/sec/insecure"
	"github.com/riteshRcH/go-edge-device-lib/core/transport"

	csms "github.com/riteshRcH/go-edge-device-lib/conn-security-multistream"
	mplex "github.com/riteshRcH/go-edge-device-lib/mplex"
	mocknetwork "github.com/riteshRcH/go-edge-device-lib/testing/mocks/network"
	ttransport "github.com/riteshRcH/go-edge-device-lib/testing/suites/transport"
	tptu "github.com/riteshRcH/go-edge-device-lib/upgrader"

	ma "github.com/riteshRcH/go-edge-device-lib/multiaddr"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestTcpTransport(t *testing.T) {
	for i := 0; i < 2; i++ {
		peerA, ia := makeInsecureMuxer(t)
		_, ib := makeInsecureMuxer(t)

		ua, err := tptu.New(ia, new(mplex.Transport))
		require.NoError(t, err)
		ta, err := NewTCPTransport(ua, nil)
		require.NoError(t, err)
		ub, err := tptu.New(ib, new(mplex.Transport))
		require.NoError(t, err)
		tb, err := NewTCPTransport(ub, nil)
		require.NoError(t, err)

		zero := "/ip4/127.0.0.1/tcp/0"
		ttransport.SubtestTransport(t, ta, tb, zero, peerA)

		envReuseportVal = false
	}
	envReuseportVal = true
}

func TestResourceManager(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	peerA, ia := makeInsecureMuxer(t)
	_, ib := makeInsecureMuxer(t)

	ua, err := tptu.New(ia, new(mplex.Transport))
	require.NoError(t, err)
	ta, err := NewTCPTransport(ua, nil)
	require.NoError(t, err)
	ln, err := ta.Listen(ma.StringCast("/ip4/127.0.0.1/tcp/0"))
	require.NoError(t, err)
	defer ln.Close()

	ub, err := tptu.New(ib, new(mplex.Transport))
	require.NoError(t, err)
	rcmgr := mocknetwork.NewMockResourceManager(ctrl)
	tb, err := NewTCPTransport(ub, rcmgr)
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		scope := mocknetwork.NewMockConnManagementScope(ctrl)
		rcmgr.EXPECT().OpenConnection(network.DirOutbound, true).Return(scope, nil)
		scope.EXPECT().SetPeer(peerA)
		scope.EXPECT().PeerScope().Return(network.NullScope).AnyTimes() // called by the upgrader
		conn, err := tb.Dial(context.Background(), ln.Multiaddr(), peerA)
		require.NoError(t, err)
		scope.EXPECT().Done()
		defer conn.Close()
	})

	t.Run("connection denied", func(t *testing.T) {
		rerr := errors.New("nope")
		rcmgr.EXPECT().OpenConnection(network.DirOutbound, true).Return(nil, rerr)
		_, err = tb.Dial(context.Background(), ln.Multiaddr(), peerA)
		require.ErrorIs(t, err, rerr)
	})

	t.Run("peer denied", func(t *testing.T) {
		scope := mocknetwork.NewMockConnManagementScope(ctrl)
		rcmgr.EXPECT().OpenConnection(network.DirOutbound, true).Return(scope, nil)
		rerr := errors.New("nope")
		scope.EXPECT().SetPeer(peerA).Return(rerr)
		scope.EXPECT().Done()
		_, err = tb.Dial(context.Background(), ln.Multiaddr(), peerA)
		require.ErrorIs(t, err, rerr)
	})
}

func TestTcpTransportCantDialDNS(t *testing.T) {
	for i := 0; i < 2; i++ {
		dnsa, err := ma.NewMultiaddr("/dns4/example.com/tcp/1234")
		require.NoError(t, err)

		var u transport.Upgrader
		tpt, err := NewTCPTransport(u, nil)
		require.NoError(t, err)

		if tpt.CanDial(dnsa) {
			t.Fatal("shouldn't be able to dial dns")
		}

		envReuseportVal = false
	}
	envReuseportVal = true
}

func TestTcpTransportCantListenUtp(t *testing.T) {
	for i := 0; i < 2; i++ {
		utpa, err := ma.NewMultiaddr("/ip4/127.0.0.1/udp/0/utp")
		require.NoError(t, err)

		var u transport.Upgrader
		tpt, err := NewTCPTransport(u, nil)
		require.NoError(t, err)

		_, err = tpt.Listen(utpa)
		require.Error(t, err, "shouldn't be able to listen on utp addr with tcp transport")

		envReuseportVal = false
	}
	envReuseportVal = true
}

func makeInsecureMuxer(t *testing.T) (peer.ID, sec.SecureMuxer) {
	t.Helper()
	priv, _, err := crypto.GenerateKeyPair(crypto.Ed25519, 256)
	require.NoError(t, err)
	id, err := peer.IDFromPrivateKey(priv)
	require.NoError(t, err)
	var secMuxer csms.SSMuxer
	secMuxer.AddTransport(insecure.ID, insecure.NewWithIdentity(id, priv))
	return id, &secMuxer
}
