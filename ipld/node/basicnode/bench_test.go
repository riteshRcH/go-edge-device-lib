package basicnode_test

import (
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/ipld/node/basicnode"
	"github.com/riteshRcH/go-edge-device-lib/ipld/node/tests"
)

func BenchmarkSpec_Walk_Map3StrInt(b *testing.B) {
	tests.BenchmarkSpec_Walk_Map3StrInt(b, basicnode.Prototype.Any)
}

func BenchmarkSpec_Walk_MapNStrMap3StrInt(b *testing.B) {
	tests.BenchmarkSpec_Walk_MapNStrMap3StrInt(b, basicnode.Prototype.Any)
}
