package swarm

import (
	"fmt"
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/core/test"
	ma "github.com/riteshRcH/go-edge-device-lib/multiaddr"

	"github.com/stretchr/testify/require"
)

func TestIsFdConsuming(t *testing.T) {
	tcs := map[string]struct {
		addr          string
		isFdConsuming bool
	}{
		"tcp": {
			addr:          "/ip4/127.0.0.1/tcp/20",
			isFdConsuming: true,
		},
		"addr-without-registered-transport": {
			addr:          "/ip4/127.0.0.1/tcp/20/ws",
			isFdConsuming: true,
		},
		"relay-tcp": {
			addr:          fmt.Sprintf("/ip4/127.0.0.1/tcp/20/p2p-circuit/p2p/%s", test.RandPeerIDFatal(t)),
			isFdConsuming: true,
		},
		"relay-without-serveraddr": {
			addr:          fmt.Sprintf("/p2p-circuit/p2p/%s", test.RandPeerIDFatal(t)),
			isFdConsuming: true,
		},
		"relay-without-registered-transport-server": {
			addr:          fmt.Sprintf("/ip4/127.0.0.1/tcp/20/ws/p2p-circuit/p2p/%s", test.RandPeerIDFatal(t)),
			isFdConsuming: true,
		},
	}

	for name := range tcs {
		maddr, err := ma.NewMultiaddr(tcs[name].addr)
		require.NoError(t, err, name)
		require.Equal(t, tcs[name].isFdConsuming, isFdConsumingAddr(maddr), name)
	}
}
