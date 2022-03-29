package itest

import (
	"context"
	"testing"

	libp2p "github.com/riteshRcH/go-edge-device-lib"
	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	"github.com/riteshRcH/go-edge-device-lib/core/peerstore"

	"github.com/stretchr/testify/require"
)

func createEchos(t *testing.T, count int, makeOpts ...func(int) libp2p.Option) []*Echo {
	result := make([]*Echo, 0, count)

	for i := 0; i < count; i++ {
		opts := make([]libp2p.Option, 0, len(makeOpts))
		for _, makeOpt := range makeOpts {
			opts = append(opts, makeOpt(i))
		}

		h, err := libp2p.New(opts...)
		if err != nil {
			t.Fatal(err)
		}

		e := NewEcho(h)
		result = append(result, e)
	}

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if i == j {
				continue
			}

			result[i].Host.Peerstore().AddAddrs(result[j].Host.ID(), result[j].Host.Addrs(), peerstore.PermanentAddrTTL)
		}
	}

	return result
}

func closeEchos(echos []*Echo) {
	for _, e := range echos {
		e.Host.Close()
	}
}

func checkEchoStatus(t *testing.T, e *Echo, expected EchoStatus) {
	t.Helper()
	require.Equal(t, expected, e.Status())
}

func TestEcho(t *testing.T) {
	echos := createEchos(t, 2)
	defer closeEchos(echos)

	if err := echos[0].Host.Connect(context.TODO(), peer.AddrInfo{ID: echos[1].Host.ID()}); err != nil {
		t.Fatal(err)
	}

	if err := echos[0].Echo(echos[1].Host.ID(), "hello libp2p"); err != nil {
		t.Fatal(err)
	}

	checkEchoStatus(t, echos[1], EchoStatus{
		StreamsIn: 1,
		EchosIn:   1,
		EchosOut:  1,
	})
}
