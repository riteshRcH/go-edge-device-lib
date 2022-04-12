package basicnode_test

import (
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/ipld/node/basicnode"
	"github.com/riteshRcH/go-edge-device-lib/ipld/node/tests"
)

func TestString(t *testing.T) {
	tests.SpecTestString(t, basicnode.Prototype__String{})
}
