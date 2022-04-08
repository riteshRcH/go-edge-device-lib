package upgrader

import (
	"fmt"

	"github.com/riteshRcH/go-edge-device-lib/core/mux"
	"github.com/riteshRcH/go-edge-device-lib/core/network"
	"github.com/riteshRcH/go-edge-device-lib/core/transport"
)

type transportConn struct {
	mux.MuxedConn
	network.ConnMultiaddrs
	network.ConnSecurity
	transport transport.Transport
	scope     network.ConnManagementScope
	stat      network.ConnStats
}

var _ transport.CapableConn = &transportConn{}

func (t *transportConn) Transport() transport.Transport {
	return t.transport
}

func (t *transportConn) String() string {
	ts := ""
	if s, ok := t.transport.(fmt.Stringer); ok {
		ts = "[" + s.String() + "]"
	}
	return fmt.Sprintf(
		"<stream.Conn%s %s (%s) <-> %s (%s)>",
		ts,
		t.LocalMultiaddr(),
		t.LocalPeer(),
		t.RemoteMultiaddr(),
		t.RemotePeer(),
	)
}

func (t *transportConn) Stat() network.ConnStats {
	return t.stat
}

func (t *transportConn) Scope() network.ConnScope {
	return t.scope
}

func (t *transportConn) Close() error {
	defer t.scope.Done()
	return t.MuxedConn.Close()
}
