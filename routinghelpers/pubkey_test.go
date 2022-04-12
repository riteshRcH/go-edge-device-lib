package routinghelpers

import (
	"context"
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/core/routing"
	"github.com/riteshRcH/go-edge-device-lib/core/test"
)

func TestGetPublicKey(t *testing.T) {
	d := Parallel{
		Routers: []routing.Routing{
			Parallel{
				Routers: []routing.Routing{
					&Compose{
						ValueStore: &LimitedValueStore{
							ValueStore: new(dummyValueStore),
							Namespaces: []string{"other"},
						},
					},
				},
			},
			Tiered{
				Routers: []routing.Routing{
					&Compose{
						ValueStore: &LimitedValueStore{
							ValueStore: new(dummyValueStore),
							Namespaces: []string{"pk"},
						},
					},
				},
			},
			&Compose{
				ValueStore: &LimitedValueStore{
					ValueStore: new(dummyValueStore),
					Namespaces: []string{"other", "pk"},
				},
			},
			&struct{ Compose }{Compose{ValueStore: &LimitedValueStore{ValueStore: Null{}}}},
			&struct{ Compose }{},
		},
	}

	pid, _ := test.RandPeerID()

	ctx := context.Background()
	if _, err := d.GetPublicKey(ctx, pid); err != routing.ErrNotFound {
		t.Fatal(err)
	}
}
