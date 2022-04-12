package routinghelpers

import (
	"context"
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	"github.com/riteshRcH/go-edge-device-lib/core/routing"

	cid "github.com/riteshRcH/go-edge-device-lib/cid"
)

func TestNull(t *testing.T) {
	var n Null
	ctx := context.Background()
	if err := n.PutValue(ctx, "anything", nil); err != routing.ErrNotSupported {
		t.Fatal(err)
	}
	if _, err := n.GetValue(ctx, "anything", nil); err != routing.ErrNotFound {
		t.Fatal(err)
	}
	if err := n.Provide(ctx, cid.Cid{}, false); err != routing.ErrNotSupported {
		t.Fatal(err)
	}
	if _, ok := <-n.FindProvidersAsync(ctx, cid.Cid{}, 10); ok {
		t.Fatal("expected no values")
	}
	if _, err := n.FindPeer(ctx, peer.ID("thing")); err != routing.ErrNotFound {
		t.Fatal(err)
	}
	if err := n.Close(); err != nil {
		t.Fatal(err)
	}
}
