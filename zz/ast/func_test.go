package ast

import (
	"fmt"
	"testing"
)

var ListTypeSpecifier1 = &ListTypeSpecifier{
	elem: &ListElementTypeSpecifier{
		elem: &ListElementTypeSpecifier{
			elem: &SimpleTypeSpecifier{name: "int"},
			typ:  ListElementTypeSpecifierSimple,
		},
		typ: ListElementTypeSpecifierNested,
	},
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

	fmt.Println(d.String())
}

func TestFuncTypeSpecifier_String(t *testing.T) {
	s := FuncTypeSpecifier1

	fmt.Println(s.String())
}

func TestFuncInitExpr_String(t *testing.T) {
	e := &FuncInitExpr{
		typeSpecifier: FuncTypeSpecifier1,
		stmtList: []FuncStatementer{
			AssignStmt1,
		},
	}

	fmt.Println(e.String())
}
