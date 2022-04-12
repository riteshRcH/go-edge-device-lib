package codec_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	_ "github.com/riteshRcH/go-edge-device-lib/ipld/codec/cbor"
	_ "github.com/riteshRcH/go-edge-device-lib/ipld/codec/dagcbor"
	_ "github.com/riteshRcH/go-edge-device-lib/ipld/codec/dagjson"
	_ "github.com/riteshRcH/go-edge-device-lib/ipld/codec/json"
	mcregistry "github.com/riteshRcH/go-edge-device-lib/ipld/multicodec"
	basicnode "github.com/riteshRcH/go-edge-device-lib/ipld/node/basic"
	"github.com/riteshRcH/go-edge-device-lib/multicodec"
)

func TestDecodeZero(t *testing.T) {
	for _, code := range []multicodec.Code{
		multicodec.Cbor,
		multicodec.DagCbor,
		multicodec.Json,
		multicodec.DagJson,
	} {
		t.Run(code.String(), func(t *testing.T) {
			nb := basicnode.Prototype.Any.NewBuilder()
			decode, err := mcregistry.LookupDecoder(uint64(code))
			if err != nil {
				t.Fatal(err)
			}

			err = decode(nb, strings.NewReader(""))
			if !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
