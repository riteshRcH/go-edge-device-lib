package basicnode_test

import (
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/ipld/node/basicnode"
	"github.com/riteshRcH/go-edge-device-lib/ipld/node/tests"
)

func TestList(t *testing.T) {
	tests.SpecTestListString(t, basicnode.Prototype.List)
}
