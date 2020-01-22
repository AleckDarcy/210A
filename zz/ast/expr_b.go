package ast

import "fmt"

type BExprCompareOpType int64

const (
	BExprCompareEQ BExprCompareOpType = iota
)

func (t BExprCompareOpType) toString(ident string) string {
	switch t {
	case BExprCompareEQ:
		return ident + "Equals"
	default:
		return ident + "undefined"
	}
}

func (t BExprCompareOpType) String() string {
	return t.toString("")
}

type BExprCompare struct {
	e1, e2 AExpr
	op     BExprCompareOpType
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

func (t BExprBinaryOpType) toString(ident string) string {
	switch t {
	case BExprBinaryEQ:
		return ident + "Equals"
	default:
		return ident + "undefined"
	}
}

func (t BExprBinaryOpType) String() string {
	return t.toString("")
}

type BExprBinaryBool struct {
	value bool
}

type BExprBinary struct {
	e1, e2 BExpr
	op     BExprBinaryOpType
}

func (e *BExprBinary) bExpr() {}

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
