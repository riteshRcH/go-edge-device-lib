package peerstream_multiplex

import (
	"net"

	"github.com/riteshRcH/go-edge-device-lib/core/network"

	mp "github.com/riteshRcH/go-edge-device-lib/mplex"
)

// DefaultTransport has default settings for Transport
var DefaultTransport = &Transport{}

var _ network.Multiplexer = &Transport{}

// Transport implements mux.Multiplexer that constructs
// mplex-backed muxed connections.
type Transport struct{}

func (t *Transport) NewConn(nc net.Conn, isServer bool, scope network.PeerScope) (network.MuxedConn, error) {
	m, err := mp.NewMultiplex(nc, isServer, scope)
	if err != nil {
		return nil, err
	}
	return (*conn)(m), nil
}
