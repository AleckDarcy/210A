package ast

import "fmt"

type BExprCompareOpType int64

const (
	BExprCompareEQ BExprCompareOpType = iota
	BExprCompareLT
	BExprCompareGT
	BExprCompareLEQ
	BExprCompareGEQ
	BExprCompareNEQ
)

func (t BExprCompareOpType) toString(indent string) string {
	switch t {
	case BExprCompareEQ:
		return indent + "Equals"
	case BExprCompareLT:
		return indent + "Less Than"
	case BExprCompareGT:
		return indent + "Greater Than"
	case BExprCompareLEQ:
		return indent + "Less or Equal"
	case BExprCompareGEQ:
		return indent + "Greater or Equal"
	case BExprCompareNEQ:
		return indent + "Not Equals"
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

func (e *BExprCompare) funcReturnParaer() {}

func (e *BExprCompare) funcExecuteParaer() {}

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

func (e *BExprCompare) E1() AExpr {
	return e.e1
}

func (e *BExprCompare) E2() AExpr {
	return e.e2
}

func (e *BExprCompare) Op() BExprCompareOpType {
	return e.op
}

type BExprBinaryOpType int64

const (
	BExprBinaryEQ BExprBinaryOpType = iota
	BExprBinaryAND
	BExprBinaryOR
	BExprBinaryNEQ
	BExprBinaryNOT
)

func (t BExprBinaryOpType) toString(indent string) string {
	switch t {
	case BExprBinaryEQ:
		return indent + "Equals"
	case BExprBinaryAND:
		return indent + "And"
	case BExprBinaryOR:
		return indent + "Or"
	case BExprBinaryNEQ:
		return indent + "Not Equals"
	case BExprBinaryNOT:
		return indent + "Not"
	default:
		return indent + "undefined"
	}
}

func PriorLevel(t interface{}) int {
	switch t.(type) {
	case BExprCompare:
		return 1
	case BExprBinary:
		return 0
	default:
		panic("undefined")
	}
}

func PriorTo(t interface{}, a interface{}) bool {
	level1 := PriorLevel(t)
	level2 := PriorLevel(a)

	return level1 > level2
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

func (e *BExprBinary) funcExecuteParaer() {}

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

func (e *BExprBinary) E1() BExpr {
	return e.e1
}

func (e *BExprBinary) E2() BExpr {
	return e.e2
}

func (e *BExprBinary) Op() BExprBinaryOpType {
	return e.op
}
