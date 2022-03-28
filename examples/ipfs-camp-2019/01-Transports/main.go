package main

import (
	libp2p "github.com/riteshRcH/go-edge-device-lib"
)

func main() {
	// TODO: add some libp2p.Transport options to this chain!
	transports := libp2p.ChainOptions()

	host, err := libp2p.New(transports)
	if err != nil {
		panic(err)
	}

	host.Close()
}
