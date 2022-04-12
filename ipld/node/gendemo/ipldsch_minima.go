package gendemo

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	"fmt"

	"github.com/riteshRcH/go-edge-device-lib/ipld/datamodel"
	"github.com/riteshRcH/go-edge-device-lib/ipld/schema"
)

const (
	midvalue  = schema.Maybe(4)
	allowNull = schema.Maybe(5)
)

type maState uint8

const (
	maState_initial maState = iota
	maState_midKey
	maState_expectValue
	maState_midValue
	maState_finished
)

type laState uint8

const (
	laState_initial laState = iota
	laState_midValue
	laState_finished
)

type _ErrorThunkAssembler struct {
	e error
}

func (ea _ErrorThunkAssembler) BeginMap(_ int64) (datamodel.MapAssembler, error)   { return nil, ea.e }
func (ea _ErrorThunkAssembler) BeginList(_ int64) (datamodel.ListAssembler, error) { return nil, ea.e }
func (ea _ErrorThunkAssembler) AssignNull() error                                  { return ea.e }
func (ea _ErrorThunkAssembler) AssignBool(bool) error                              { return ea.e }
func (ea _ErrorThunkAssembler) AssignInt(int64) error                              { return ea.e }
func (ea _ErrorThunkAssembler) AssignFloat(float64) error                          { return ea.e }
func (ea _ErrorThunkAssembler) AssignString(string) error                          { return ea.e }
func (ea _ErrorThunkAssembler) AssignBytes([]byte) error                           { return ea.e }
func (ea _ErrorThunkAssembler) AssignLink(datamodel.Link) error                    { return ea.e }
func (ea _ErrorThunkAssembler) AssignNode(datamodel.Node) error                    { return ea.e }
func (ea _ErrorThunkAssembler) Prototype() datamodel.NodePrototype {
	panic(fmt.Errorf("cannot get prototype from error-carrying assembler: already derailed with error: %w", ea.e))
}
