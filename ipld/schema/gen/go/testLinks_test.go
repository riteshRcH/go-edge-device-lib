package gengo

import (
	"runtime"
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/ipld/node/tests"
)

func TestLinks(t *testing.T) {
	if runtime.GOOS != "darwin" { // TODO: enable parallelism on macos
		t.Parallel()
	}

	engine := &genAndCompileEngine{prefix: "links"}
	tests.SchemaTestLinks(t, engine)
}
