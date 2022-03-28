package main

import (
	"github.com/libp2p/go-tcp-transport"
	ws "github.com/libp2p/go-ws-transport"
	libp2p "github.com/riteshRcH/go-edge-device-lib"
)

func main() {
	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(ws.New),
	)

	// TODO: add some listen addresses with the libp2p.ListenAddrs or
	// libp2p.ListenAddrStrings configuration options.

	host, err := libp2p.New(transports)
	if err != nil {
		panic(err)
	}

	// TODO: with our host made, let's connect to our bootstrap peer

	host.Close()
}
