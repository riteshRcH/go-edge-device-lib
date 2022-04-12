package gengo

import (
	"runtime"
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/ipld/node/tests"
)

func TestStructReprTuple(t *testing.T) {
	if runtime.GOOS != "darwin" { // TODO: enable parallelism on macos
		t.Parallel()
	}

	engine := &genAndCompileEngine{prefix: "struct-tuple"}
	tests.SchemaTestStructReprTuple(t, engine)
}
