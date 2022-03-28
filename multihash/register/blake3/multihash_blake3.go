/*
	This package has no purpose except to register the blake3 hash function.

	It is meant to be used as a side-effecting import, e.g.

		import (
			_ ""github.com/riteshRcH/go-edge-device-lib/multihash/register/blake3"
		)
*/
package blake3

import (
	"hash"

	"lukechampine.com/blake3"

	multihash "github.com/riteshRcH/go-edge-device-lib/multihash/core"
)

func init() {
	multihash.Register(multihash.BLAKE3, func() hash.Hash { h := blake3.New(32, nil); return h })

}
