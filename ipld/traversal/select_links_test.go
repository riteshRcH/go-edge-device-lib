package traversal_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/riteshRcH/go-edge-device-lib/ipld/datamodel"
	"github.com/riteshRcH/go-edge-device-lib/ipld/traversal"
)

func TestSelectLinks(t *testing.T) {

	t.Run("Scalar", func(t *testing.T) {
		lnks, _ := traversal.SelectLinks(leafAlpha)
		qt.Check(t, lnks, deepEqualsAllowAllUnexported, []datamodel.Link(nil))
	})
	t.Run("DeepMap", func(t *testing.T) {
		lnks, _ := traversal.SelectLinks(middleMapNode)
		qt.Check(t, lnks, deepEqualsAllowAllUnexported, []datamodel.Link{leafAlphaLnk})
	})
	t.Run("List", func(t *testing.T) {
		lnks, _ := traversal.SelectLinks(rootNode)
		qt.Check(t, lnks, deepEqualsAllowAllUnexported, []datamodel.Link{leafAlphaLnk, middleMapNodeLnk, middleListNodeLnk})
	})
}
