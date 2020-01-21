package ast

import (
	"fmt"
	"testing"
)

var FuncStmtList1 = []FuncStatementer{
	AssignStmt1,
}

var ParaDeclaratorWithIdentity1 = &ParaDeclaratorWithIdentity{
	declList: []*Identifier{
		{name: "a"}, {name: "b"},
	},
	typeSpecifier: ListTypeSpecifier1,
}

var ParaDeclaratorWithIdentity2 = &ParaDeclaratorWithIdentity{
	declList: []*Identifier{
		{name: "c"}, {name: "d"},
	},
	typeSpecifier: &SimpleTypeSpecifier{name: "int"},
}

var FuncTypeSpecifier1 = &FuncTypeSpecifier{
	paraList: []*ParaDeclaratorWithIdentity{
		ParaDeclaratorWithIdentity1,
		ParaDeclaratorWithIdentity2,
	},
	typeSpecifierList: []TypeSpecifier{
		ListTypeSpecifier1,
		&SimpleTypeSpecifier{name: "int"},
	},
}

func TestParaDeclaratorWithIdentity_String(t *testing.T) {
	d := ParaDeclaratorWithIdentity1
	fmt.Println(d)
}

func TestFuncTypeSpecifier_String(t *testing.T) {
	s := FuncTypeSpecifier1
	s.typeSpecifier()
	fmt.Println(s)
}

func TestFuncInitExpr_String(t *testing.T) {
	e := &FuncInitExpr{
		typeSpecifier: FuncTypeSpecifier1,
		stmtList:      FuncStmtList1,
	}
	e.assignIniter()
	fmt.Println(e)
}
