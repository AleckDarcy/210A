package ast

import "fmt"

type AExprArithOpType int64

const (
	AExprArithAdd AExprArithOpType = iota
	AExprArithSub
	AExprArithMul
	AExprArithDiv
)

func (t AExprArithOpType) toString(ident string) string {
	switch t {
	case AExprArithAdd:
		return ident + "Add"
	case AExprArithSub:
		return ident + "Sub"
	case AExprArithMul:
		return ident + "Mul"
	case AExprArithDiv:
		return ident + "Div"
	default:
		return ident + "undefined"
	}
}

func (t AExprArithOpType) String() string {
	return t.toString("")
}

func (t AExprArithOpType) PriorLevel() int {
	switch t {
	case AExprArithAdd:
		return 0
	case AExprArithSub:
		return 0
	case AExprArithMul:
		return 1
	case AExprArithDiv:
		return 1
	default:
		panic("undefined")
	}

}

func (t AExprArithOpType) PriorTo(a AExprArithOpType) bool {
	level1 := t.PriorLevel()
	level2 := a.PriorLevel()

	return level1 > level2
}

type AExprSimple struct {
	e AExpr
}

var AExprSimpleHelper *AExprSimple

func (e *AExprSimple) New(exp AExpr) *AExprSimple {
	return &AExprSimple{e: exp}
}

func (e *AExprSimple) aExpr() {}

func (e *AExprSimple) assignIniter() {}

func (e *AExprSimple) funcReturnParaer() {}

func (e *AExprSimple) funcExecuteParaer() {}

func (e *AExprSimple) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sAExprSimple {\n"+
		"%s..E:\n"+
		"%s\n"+
		"%s}",
		ident, ident, e.e.toString(ident+"...."),
		ident,
	)
}

func (e *AExprSimple) String() string {
	return e.toString("")
}

func (e *AExprSimple) E() AExpr {
	return e.e
}

type AExprArith struct {
	e1, e2 AExpr
	op     AExprArithOpType
}

var AExprArithHelper *AExprArith

func (e *AExprArith) New(e1, e2 AExpr, op AExprArithOpType) *AExprArith {
	return &AExprArith{e1: e1, e2: e2, op: op}
}

func (e *AExprArith) aExpr() {}

func (e *AExprArith) assignIniter() {}

func (e *AExprArith) funcReturnParaer() {}

func (e *AExprArith) funcExecuteParaer() {}

func (e *AExprArith) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sAExprArith {\n"+
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

func (e *AExprArith) String() string {
	return e.toString("")
}

func (e *AExprArith) E1() AExpr {
	return e.e1
}

func (e *AExprArith) E2() AExpr {
	return e.e2
}

func (e *AExprArith) Op() AExprArithOpType {
	return e.op
}

type ArithTranspose struct {
	e AExpr
}

var ArithTransposeHelper *ArithTranspose

func (t *ArithTranspose) aExpr() {}

func (t *ArithTranspose) assignIniter() {}

func (t *ArithTranspose) New(e AExpr) *ArithTranspose {
	return &ArithTranspose{e: e}
}

func (t *ArithTranspose) toString(indent string) string {
	return fmt.Sprintf(""+
		"%sArithTranspose {\n"+
		"%s..Expr:\n"+
		"%s\n"+
		"%s}\n",
		indent, indent, t.e.toString(indent+"...."),
		indent,
	)
}

func (t *ArithTranspose) String() string {
	return t.toString("")
}

func (t *ArithTranspose) E() AExpr {
	return t.e
}
