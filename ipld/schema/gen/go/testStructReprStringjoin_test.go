package gengo

import (
	"runtime"
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/ipld/node/tests"
)

func TestStructReprStringjoin(t *testing.T) {
	if runtime.GOOS != "darwin" { // TODO: enable parallelism on macos
		t.Parallel()
	}

	engine := &genAndCompileEngine{prefix: "struct-str-join"}
	tests.SchemaTestStructReprStringjoin(t, engine)
}
