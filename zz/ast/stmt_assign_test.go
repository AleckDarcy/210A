package ast

import (
	"fmt"
	"testing"
)

var AssignStmt1 = &AssignStmt{
	declList: []Declaratorer{
		&Identifier{name: "a"},
		&listElementExpr{
			name: &Identifier{name: "b"},
			list: []*listElementIndex{
				{
					e: &AExprArith{
						e1: &IntegerLiteral{value: 1},
						e2: &IntegerLiteral{value: 2},
					},
				},
				{
					e: &IntegerLiteral{value: 3},
				},
			},
		},
		&Identifier{name: "c"},
	},
	initList: []AssignIniter{
		&AExprArith{
			e1: &IntegerLiteral{value: 1},
			e2: &IntegerLiteral{value: 2},
		},
		&ListInitExpr{
			typeSpecifier: &ListTypeSpecifier{
				elem: &ListElementTypeSpecifier{
					elem: &ListElementTypeSpecifier{
						elem: &SimpleTypeSpecifier{name: "int"},
						typ:  ListElementTypeSpecifierSimple,
					},
					typ: ListElementTypeSpecifierNested,
				},
			},
			size: &AExprArith{
				e1: &IntegerLiteral{value: 1},
				e2: &IntegerLiteral{value: 2},
			},
		},
		&ListInitExpr{
			typeSpecifier: &ListTypeSpecifier{
				elem: &ListElementTypeSpecifier{
					elem: &SimpleTypeSpecifier{name: "int"},
				},
			},
			size: &IntegerLiteral{value: 3},
		},
	},
}

func TestAssignStmt_String(t *testing.T) {
	s := AssignStmt1

	fmt.Println(s.String())

	s.statementer()
}
