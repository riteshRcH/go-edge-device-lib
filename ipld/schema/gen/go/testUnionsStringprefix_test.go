package gengo

import (
	"runtime"
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/ipld/node/tests"
	"github.com/riteshRcH/go-edge-device-lib/ipld/schema"
)

func TestUnionStringprefix(t *testing.T) {
	if runtime.GOOS != "darwin" { // TODO: enable parallelism on macos
		t.Parallel()
	}

	for _, engine := range []*genAndCompileEngine{
		{
			subtestName: "union-using-embed",
			prefix:      "union-stringprefix-using-embed",
			adjCfg: AdjunctCfg{
				CfgUnionMemlayout: map[schema.TypeName]string{"WheeUnion": "embedAll"},
			},
		},
		{
			subtestName: "union-using-interface",
			prefix:      "union-stringprefix-using-interface",
			adjCfg: AdjunctCfg{
				CfgUnionMemlayout: map[schema.TypeName]string{"WheeUnion": "interface"},
			},
		},
	} {
		t.Run(engine.subtestName, func(t *testing.T) {
			tests.SchemaTestUnionStringprefix(t, engine)
		})
	}
}
