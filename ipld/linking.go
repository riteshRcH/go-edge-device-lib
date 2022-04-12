package ipld

import (
	"github.com/riteshRcH/go-edge-device-lib/ipld/linking"
)

type (
	LinkSystem  = linking.LinkSystem
	LinkContext = linking.LinkContext
)

type (
	BlockReadOpener     = linking.BlockReadOpener
	BlockWriteOpener    = linking.BlockWriteOpener
	BlockWriteCommitter = linking.BlockWriteCommitter
	NodeReifier         = linking.NodeReifier
)
