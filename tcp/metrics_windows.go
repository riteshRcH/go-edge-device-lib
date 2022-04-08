//go:build windows
// +build windows

package tcp

import manet "github.com/riteshRcH/go-edge-device-lib/multiaddr/net"

func newTracingConn(c manet.Conn, _ bool) (manet.Conn, error) { return c, nil }
func newTracingListener(l manet.Listener) manet.Listener      { return l }
