package upgrader_test

import (
	"sync"

	"github.com/riteshRcH/go-edge-device-lib/core/connmgr"
	"github.com/riteshRcH/go-edge-device-lib/core/control"
	"github.com/riteshRcH/go-edge-device-lib/core/network"
	"github.com/riteshRcH/go-edge-device-lib/core/peer"

	ma "github.com/riteshRcH/go-edge-device-lib/multiaddr"
)

type testGater struct {
	sync.Mutex

	blockAccept, blockSecured bool
}

var _ connmgr.ConnectionGater = (*testGater)(nil)

func (t *testGater) BlockAccept(block bool) {
	t.Lock()
	defer t.Unlock()

	t.blockAccept = block
}

func (t *testGater) BlockSecured(block bool) {
	t.Lock()
	defer t.Unlock()

	t.blockSecured = block
}

func (t *testGater) InterceptPeerDial(p peer.ID) (allow bool) {
	panic("not implemented")
}

func (t *testGater) InterceptAddrDial(id peer.ID, multiaddr ma.Multiaddr) (allow bool) {
	panic("not implemented")
}

func (t *testGater) InterceptAccept(multiaddrs network.ConnMultiaddrs) (allow bool) {
	t.Lock()
	defer t.Unlock()

	return !t.blockAccept
}

func (t *testGater) InterceptSecured(direction network.Direction, id peer.ID, multiaddrs network.ConnMultiaddrs) (allow bool) {
	t.Lock()
	defer t.Unlock()

	return !t.blockSecured
}

func (t *testGater) InterceptUpgraded(conn network.Conn) (allow bool, reason control.DisconnectReason) {
	panic("not implemented")
}
