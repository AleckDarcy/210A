package ast

import (
	"fmt"
	"testing"
)

var ParaDeclarator1 = &ParaDeclarator{name: &Identifier{name: "para"}}

var FuncIdentifier1 = &FuncIdentifier{name: IdentifierHelper.New("function")}

func TestParaDeclarator_String(t *testing.T) {
	d := ParaDeclarator1
	fmt.Println(d)
}

func TestParaDeclaratorWithIdentity_String(t *testing.T) {
	d := ParaDeclaratorWithIdentity1
	fmt.Println(d)
}

func TestFuncTypeSpecifier_String(t *testing.T) {
	s := FuncTypeSpecifier1
	s.typeSpecifierer()
	fmt.Println(s)
}

func TestFuncInitExpr_String(t *testing.T) {
	e := FuncInitExpr1
	e.assignIniter()
	fmt.Println(e)
}

func TestFuncIdentifier_String(t *testing.T) {
	i := FuncIdentifier1
	fmt.Println(i)
}

func TestFuncTypeSpecifierWithName_String(t *testing.T) {
	s := FuncTypeSpecifierWithName2
	fmt.Println(s)
}

func TestFuncDefinition_String(t *testing.T) {
	d := FuncDefinition2
	d.definitioner()
	fmt.Println(d)
}

func TestFuncReturnStatement_String(t *testing.T) {
	s := FuncReturnStatement1
	s.funcStatementer()
	fmt.Println(s)
}

func TestFuncExecutePara_String(t *testing.T) {
	p := FuncExecutePara1
	fmt.Println(p)
}

func TestFuncExecuteExpression_String(t *testing.T) {
	e := FuncExecuteExpression1
	e.assignIniter()
	fmt.Println(e)
}

func TestFuncExecuteStatement_String(t *testing.T) {
	e := FuncExecuteStatement1
	e.funcStatementer()
	fmt.Println(e)
}
