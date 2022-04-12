//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/riteshRcH/go-edge-device-lib/ipld/node/bindnode"
	schemadmt "github.com/riteshRcH/go-edge-device-lib/ipld/schema/dmt"
)

func main() {
	f, err := os.Create("types.go")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(f, "package schemadmt\n\n")
	if err := bindnode.ProduceGoTypes(f, schemadmt.InternalTypeSystem()); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}
