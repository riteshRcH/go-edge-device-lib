// qp helps to quickly build IPLD nodes.
//
// It contains top-level Build funcs, such as BuildMap and BuildList, which
// return the final node as well as an error.
//
// Underneath, one can use a number of Assemble functions to construct basic
// nodes, such as String or Int.
//
// Finally, functions like MapEntry and ListEntry allow inserting into maps and
// lists.
//
// These all use the same IPLD datamodel interfaces such as NodePrototype and
// NodeAssembler, but with some magic to reduce verbosity.
package qp

import (
	"fmt"

	"github.com/riteshRcH/go-edge-device-lib/ipld/datamodel"
)

type Assemble = func(datamodel.NodeAssembler)

func BuildMap(np datamodel.NodePrototype, sizeHint int64, fn func(datamodel.MapAssembler)) (_ datamodel.Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			if rerr, ok := r.(error); ok {
				err = rerr
			} else {
				// A reasonable fallback, for e.g. strings.
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	nb := np.NewBuilder()
	Map(sizeHint, fn)(nb)
	return nb.Build(), nil
}

type mapParams struct {
	sizeHint int64
	fn       func(datamodel.MapAssembler)
}

func (mp mapParams) Assemble(na datamodel.NodeAssembler) {
	ma, err := na.BeginMap(mp.sizeHint)
	if err != nil {
		panic(err)
	}
	mp.fn(ma)
	if err := ma.Finish(); err != nil {
		panic(err)
	}
}

func Map(sizeHint int64, fn func(datamodel.MapAssembler)) Assemble {
	return mapParams{sizeHint, fn}.Assemble
}

func MapEntry(ma datamodel.MapAssembler, k string, fn Assemble) {
	na, err := ma.AssembleEntry(k)
	if err != nil {
		panic(err)
	}
	fn(na)
}

func BuildList(np datamodel.NodePrototype, sizeHint int64, fn func(datamodel.ListAssembler)) (_ datamodel.Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			if rerr, ok := r.(error); ok {
				err = rerr
			} else {
				// A reasonable fallback, for e.g. strings.
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	nb := np.NewBuilder()
	List(sizeHint, fn)(nb)
	return nb.Build(), nil
}

type listParams struct {
	sizeHint int64
	fn       func(datamodel.ListAssembler)
}

func (lp listParams) Assemble(na datamodel.NodeAssembler) {
	la, err := na.BeginList(lp.sizeHint)
	if err != nil {
		panic(err)
	}
	lp.fn(la)
	if err := la.Finish(); err != nil {
		panic(err)
	}
}

func List(sizeHint int64, fn func(datamodel.ListAssembler)) Assemble {
	return listParams{sizeHint, fn}.Assemble
}

func ListEntry(la datamodel.ListAssembler, fn Assemble) {
	fn(la.AssembleValue())
}

type nullParam struct{}

func (s nullParam) Assemble(na datamodel.NodeAssembler) {
	if err := na.AssignNull(); err != nil {
		panic(err)
	}
}

func Null() Assemble {
	return nullParam{}.Assemble
}

type boolParam bool

func (s boolParam) Assemble(na datamodel.NodeAssembler) {
	if err := na.AssignBool(bool(s)); err != nil {
		panic(err)
	}
}

func Bool(b bool) Assemble {
	return boolParam(b).Assemble
}

type intParam int64

func (i intParam) Assemble(na datamodel.NodeAssembler) {
	if err := na.AssignInt(int64(i)); err != nil {
		panic(err)
	}
}

func Int(i int64) Assemble {
	return intParam(i).Assemble
}

type floatParam float64

func (f floatParam) Assemble(na datamodel.NodeAssembler) {
	if err := na.AssignFloat(float64(f)); err != nil {
		panic(err)
	}
}

func Float(f float64) Assemble {
	return intParam(f).Assemble
}

type stringParam string

func (s stringParam) Assemble(na datamodel.NodeAssembler) {
	if err := na.AssignString(string(s)); err != nil {
		panic(err)
	}
}

func String(s string) Assemble {
	return stringParam(s).Assemble
}

type bytesParam []byte

func (p bytesParam) Assemble(na datamodel.NodeAssembler) {
	if err := na.AssignBytes([]byte(p)); err != nil {
		panic(err)
	}
}

func Bytes(p []byte) Assemble {
	return bytesParam(p).Assemble
}

type linkParam struct {
	x datamodel.Link
}

func (l linkParam) Assemble(na datamodel.NodeAssembler) {
	if err := na.AssignLink(l.x); err != nil {
		panic(err)
	}
}

func Link(l datamodel.Link) Assemble {
	return linkParam{l}.Assemble
}

type nodeParam struct {
	x datamodel.Node
}

func (n nodeParam) Assemble(na datamodel.NodeAssembler) {
	if err := na.AssignNode(n.x); err != nil {
		panic(err)
	}
}

func Node(n datamodel.Node) Assemble {
	return nodeParam{n}.Assemble
}
