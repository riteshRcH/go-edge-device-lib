package dagcbor

import (
	"bytes"
	"io"
	"testing"

	qt "github.com/frankban/quicktest"

	cid "github.com/riteshRcH/go-edge-device-lib/cid"
	"github.com/riteshRcH/go-edge-device-lib/ipld/datamodel"
	"github.com/riteshRcH/go-edge-device-lib/ipld/linking"
	cidlink "github.com/riteshRcH/go-edge-device-lib/ipld/linking/cid"
	"github.com/riteshRcH/go-edge-device-lib/ipld/node/basicnode"
	nodetests "github.com/riteshRcH/go-edge-device-lib/ipld/node/tests"
)

func TestRoundtripCidlink(t *testing.T) {
	lp := cidlink.LinkPrototype{Prefix: cid.Prefix{
		Version:  1,
		Codec:    0x71,
		MhType:   0x13,
		MhLength: 4,
	}}
	lsys := cidlink.DefaultLinkSystem()

	buf := bytes.Buffer{}
	lsys.StorageWriteOpener = func(lnkCtx linking.LinkContext) (io.Writer, linking.BlockWriteCommitter, error) {
		return &buf, func(lnk datamodel.Link) error { return nil }, nil
	}
	lsys.StorageReadOpener = func(lnkCtx linking.LinkContext, lnk datamodel.Link) (io.Reader, error) {
		return bytes.NewReader(buf.Bytes()), nil
	}

	lnk, err := lsys.Store(linking.LinkContext{}, lp, n)
	qt.Assert(t, err, qt.IsNil)

	n2, err := lsys.Load(linking.LinkContext{}, lnk, basicnode.Prototype.Any)
	qt.Assert(t, err, qt.IsNil)
	qt.Check(t, n2, nodetests.NodeContentEquals, nSorted)
}
