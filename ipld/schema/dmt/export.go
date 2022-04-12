//go:build schemadmtgen
// +build schemadmtgen

package schemadmt

import "github.com/riteshRcH/go-edge-device-lib/ipld/schema"

func InternalTypeSystem() *schema.TypeSystem {
	return &schemaTypeSystem
}
