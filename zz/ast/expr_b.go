package ast

import "fmt"

type BExprCompareOpType int64

const (
	BExprCompareEQ BExprCompareOpType = iota
	BExprCompareLT
)

func (t BExprCompareOpType) toString(indent string) string {
	switch t {
	case BExprCompareEQ:
		return indent + "Equals"
	case BExprCompareLT:
		return indent + "Less Than"
	default:
		return indent + "undefined"
	}
}

func (t BExprCompareOpType) String() string {
	return t.toString("")
}

type BExprCompare struct {
	e1, e2 AExpr
	op     BExprCompareOpType
}

var BExprCompareHelper *BExprCompare

func (e *BExprCompare) New(e1, e2 AExpr, op BExprCompareOpType) *BExprCompare {
	return &BExprCompare{e1: e1, e2: e2, op: op}
}

func (e *BExprCompare) bExpr() {}

func (e *BExprCompare) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sBExprCompare {\n"+
		"%s..E1:\n"+
		"%s\n"+
		"%s..E2:\n"+
		"%s\n"+
		"%s..Op: %s\n"+
		"%s}",

		ident, ident, e.e1.toString("...."+ident),
		ident, e.e2.toString("...."+ident),
		ident, e.op,
		ident)
}

func (e *BExprCompare) String() string {
	return e.toString("")
}

type BExprBinaryOpType int64

const (
	BExprBinaryEQ BExprBinaryOpType = iota
)

func (t BExprBinaryOpType) toString(indent string) string {
	switch t {
	case BExprBinaryEQ:
		return indent + "Equals"
	default:
		return indent + "undefined"
	}
}

func (t BExprBinaryOpType) String() string {
	return t.toString("")
}

type BExprBinary struct {
	e1, e2 BExpr
	op     BExprBinaryOpType
}

var BExprBinaryHelper *BExprBinary

func (e *BExprBinary) New(e1, e2 BExpr, op BExprBinaryOpType) *BExprBinary {
	return &BExprBinary{e1: e1, e2: e2, op: op}
}

func (e *BExprBinary) bExpr() {}

func (e *BExprBinary) funcReturnParaer() {}

func (e *BExprBinary) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sBExprBinary {\n"+
		"%s..E1:\n"+
		"%s\n"+
		"%s..E2:\n"+
		"%s\n"+
		"%s..Op: %s\n"+
		"%s}",
		ident, ident, e.e1.toString("...."+ident),
		ident, e.e2.toString("...."+ident),
		ident, e.op,
		ident,
	)
}

func (e *BExprBinary) String() string {
	return e.toString("")
}
