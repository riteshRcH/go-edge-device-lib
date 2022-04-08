//go:build new_transport_no_rcmgr
// +build new_transport_no_rcmgr

package transport

import (
	libp2pquic "github.com/riteshRcH/go-edge-device-lib/quic"

	"github.com/riteshRcH/go-edge-device-lib/core/crypto"
	"github.com/riteshRcH/go-edge-device-lib/core/transport"
)

func New(key crypto.PrivKey) (transport.Transport, error) {
	return libp2pquic.NewTransport(key, nil, nil)
}
