package dht

import (
	"github.com/riteshRcH/go-edge-device-lib/core/protocol"
)

var (
	// ProtocolDHT is the default DHT protocol.
	ProtocolDHT protocol.ID = "/ipfs/kad/1.0.0"
	// DefaultProtocols spoken by the DHT.
	DefaultProtocols = []protocol.ID{ProtocolDHT}
)
