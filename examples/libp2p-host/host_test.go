package main

import (
	"testing"

	"github.com/riteshRcH/go-edge-device-lib/examples/testutils"
)

func TestMain(t *testing.T) {
	var h testutils.LogHarness
	h.ExpectPrefix("Hello World, my hosts ID is ")
	h.ExpectPrefix("Hello World, my second hosts ID is ")
	h.Run(t, run)
}
