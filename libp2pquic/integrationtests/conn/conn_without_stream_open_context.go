//go:build stream_open_no_context
// +build stream_open_no_context

package conn

import (
	"context"

	"github.com/riteshRcH/go-edge-device-lib/core/mux"
	tpt "github.com/riteshRcH/go-edge-device-lib/core/transport"
)

func OpenStream(_ context.Context, c tpt.CapableConn) (mux.MuxedStream, error) {
	return c.OpenStream()
}
