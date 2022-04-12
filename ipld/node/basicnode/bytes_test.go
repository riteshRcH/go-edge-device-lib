package basicnode_test

import (
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/ipld/node/basicnode"
	"github.com/riteshRcH/go-edge-device-lib/ipld/node/tests"
)

func TestBytes(t *testing.T) {
	tests.SpecTestBytes(t, basicnode.Prototype__Bytes{})
}
