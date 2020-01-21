package ast

import (
	"fmt"
	"testing"
)

var ListElementTypeSpecifier1 = &ListElementTypeSpecifier{
	elem: &ListElementTypeSpecifier{
		elem: &SimpleTypeSpecifier{name: "int"},
		typ:  ListElementTypeSpecifierSimple,
	},
	typ: ListElementTypeSpecifierNested,
}

var ListTypeSpecifier1 = &ListTypeSpecifier{
	elem: ListElementTypeSpecifier1,
}

var ListElementExpr1 = &ListElementExpr{
	name: &Identifier{name: "b"},
	list: []*ListElementIndex{
		{e: aExprAdd1},
		{e: &IntegerLiteral{value: 3}},
	},
}

var ListInitExpr1 = &ListInitExpr{
	typeSpecifier: ListTypeSpecifier1,
	size: aExprAdd1,
}

var ListInitExpr2 = &ListInitExpr{
	typeSpecifier: &ListTypeSpecifier{
		elem: &ListElementTypeSpecifier{
			elem: &SimpleTypeSpecifier{name: "int"},
		},
	},
	size: &IntegerLiteral{value: 3},
}

func TestListElementIndex_String(t *testing.T) {
	i := &ListElementIndex{e: aExprAdd1}
	fmt.Println(i)
}

func TestListElementExpr_String(t *testing.T) {
	e := ListElementExpr1
	e.declaratorer()
	fmt.Println(e)
}

func TestListInitExpr_String(t *testing.T) {
	e := ListInitExpr1
	e.assignIniter()
	fmt.Println(e)
}