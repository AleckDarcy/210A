package ast

var SimpleTypeSpecifier1 = &SimpleTypeSpecifier{name: "int"} // int

var SimpleTypeSpecifier2 = &SimpleTypeSpecifier{name: "float"} // float

var ListTypeSpecifier1 = &ListTypeSpecifier{ // []int
	elem: &ListElementTypeSpecifier{
		elem: SimpleTypeSpecifier1,
		typ:  ListElementTypeSpecifierSimple,
	},
}

var ListTypeSpecifier2 = &ListTypeSpecifier{ // []float
	elem: &ListElementTypeSpecifier{
		elem: SimpleTypeSpecifier2,
		typ:  ListElementTypeSpecifierSimple,
	},
}

var ListElementExpr1 = &ListElementExpr{ // a[1]
	name: &Identifier{name: "a"},
	list: []*ListElementIndex{
		{e: &AExprSimple{e: &IntegerLiteral{value: 1}}},
	},
}

var AssignStmt1 = &AssignStmt{ // a = list([]int, 4)
	declList: []Declaratorer{
		&Identifier{name: "a"},
	},
	initList: []AssignIniter{
		&ListInitExpr{
			typeSpecifier: ListTypeSpecifier1,
			size:          &AExprSimple{e: &IntegerLiteral{value: 4}},
		},
	},
}

var AssignStmt2 = &AssignStmt{ // a[1], b = 2, 3
	declList: []Declaratorer{
		ListElementExpr1,
		&Identifier{name: "b"},
	},
	initList: []AssignIniter{
		&AExprSimple{e: &IntegerLiteral{value: 2}},
		&AExprSimple{e: &IntegerLiteral{value: 3}},
	},
}

var AssignStmt3 = &AssignStmt{ // b = x + y
	declList: []Declaratorer{
		&Identifier{name: "b"},
	},
	initList: []AssignIniter{
		&AExprArith{
			e1: &AExprSimple{e: &Identifier{name: "x"}},
			e2: &AExprSimple{e: &Identifier{name: "y"}},
			op: AExprArithAdd,
		},
	},
}

/*
	if a[1] == b {
		a[1] = a[1] + 1
	}
*/
var IFExpr1 = &IfExpr{
	binExpr: &BExprCompare{ // a[1] == b
		e1: ListElementExpr1,
		e2: &AExprSimple{e: &Identifier{name: "b"}},
		op: BExprCompareEQ,
	},
	stmtList: []FuncStatementer{ // a[1] = a[1] + 1
		&AssignStmt{
			declList: []Declaratorer{
				ListElementExpr1,
			},
			initList: []AssignIniter{
				&AExprArith{
					e1: ListElementExpr1,
					e2: &AExprSimple{e: &IntegerLiteral{value: 1}},
					op: AExprArithAdd,
				},
			},
		},
	},
}

/*
	if a[1] == b + 1 {
		a[1] = a[1] + 2
	}
*/
var IFExpr2 = &IfExpr{
	binExpr: &BExprCompare{ // a[1] == b + 1
		e1: ListElementExpr1,
		e2: &AExprArith{
			e1: &AExprSimple{e: &Identifier{name: "b"}}, e2: &AExprSimple{e: &IntegerLiteral{value: 1}},
		},
		op: BExprCompareEQ,
	},
	stmtList: []FuncStatementer{ // a[1] = a[1] + 2
		&AssignStmt{
			declList: []Declaratorer{
				ListElementExpr1,
			},
			initList: []AssignIniter{
				&AExprArith{
					e1: ListElementExpr1,
					e2: &AExprSimple{e: &IntegerLiteral{value: 2}},
					op: AExprArithAdd,
				},
			},
		},
	},
}

/*
	else {
		a[1] = 0
	}
*/
var ElseExpr1 = &ElseExpr{
	stmtList: []FuncStatementer{ // a[1] = 0
		&AssignStmt{
			declList: []Declaratorer{
				ListElementExpr1,
			},
			initList: []AssignIniter{
				&AExprSimple{e: &IntegerLiteral{value: 0}},
			},
		},
	},
}

/*
	if a[1] == b {
		a[1] = a[1] + 1
	} else a[1] == b + 1 {
		a[1] = a[1] + 2
	} else {
		a[1] = 0
	}
*/
var SelectionStmt1 = &SelectionStmt{
	ifExprList: []*IfExpr{
		IFExpr1,
		IFExpr2,
	},
	elseExpr: ElseExpr1,
}

/*
	if a[1] == b {
		a[1] = a[1] + 1
	} else {
		a[1] = 0
	}

or
	a[1] == b ? a[1] = a[1] + 1 : a[1] = 0
*/
var SelectionStmt2 = &SelectionStmt{
	ifExprList: []*IfExpr{
		IFExpr1,
	},
	elseExpr: ElseExpr1,
}

/*
	if a[1] == b {
		a[1] = a[1] + 1
	}
*/
var SelectionStmt3 = &SelectionStmt{
	ifExprList: []*IfExpr{
		IFExpr1,
	},
}

/*
	for ;; {}
*/
var IterationStmt1 = &IterationStmt{
	binExpr: &BinaryLiteral{value: true},
}

/*
	for i = 1; i < 5; i = i + 1 {}
*/
var IterationStmt2 = &IterationStmt{
	initStmt: &AssignStmt{
		declList: []Declaratorer{&Identifier{name: "i"}},
		initList: []AssignIniter{&AExprSimple{e: &IntegerLiteral{value: 1}}},
	},
	binExpr: &BExprCompare{
		e1: &AExprSimple{e: &Identifier{name: "i"}},
		e2: &AExprSimple{e: &IntegerLiteral{value: 5}},
		op: BExprCompareLT,
	},
	increStmt: &AssignStmt{
		declList: []Declaratorer{&Identifier{name: "i"}},
		initList: []AssignIniter{
			&AExprArith{
				e1: &AExprSimple{e: &Identifier{name: "i"}},
				e2: &AExprSimple{e: &IntegerLiteral{value: 1}},
				op: AExprArithAdd,
			},
		},
	},
}

/*
	func function() {}
*/
var FuncDefinition1 = &FuncDefinition{
	typeSpecifier: &FuncTypeSpecifierWithName{
		name: &FuncIdentifier{name: &Identifier{name: "function"}},
	},
}

/*
	func function1(x, y int, z []float) ([]int, float, int) {
		a = list([]int, 4)
		b = x + y
		return a, b + 1, 1
	}
*/
var FuncDefinition2 = &FuncDefinition{
	typeSpecifier: &FuncTypeSpecifierWithName{
		name: &FuncIdentifier{name: &Identifier{name: "function1"}},
		paraList: []*ParaDeclaratorWithIdentity{
			{
				declList: []*Identifier{
					{name: "x"}, {name: "y"},
				},
				typeSpecifier: SimpleTypeSpecifier1,
			},
			{
				declList:      []*Identifier{{name: "z"}},
				typeSpecifier: ListTypeSpecifier2,
			},
		},
		returnList: []TypeSpecifierer{
			ListTypeSpecifier1,
			SimpleTypeSpecifier2,
			SimpleTypeSpecifier1,
		},
	},
	stmtList: []FuncStatementer{
		AssignStmt1,
		AssignStmt3,
		&FuncReturnStatement{
			returnList: []FuncReturnParaer{
				&AExprSimple{e: &Identifier{name: "a"}},
				&AExprArith{
					e1: &AExprSimple{e: &Identifier{name: "b"}},
					e2: &AExprSimple{e: &IntegerLiteral{value: 1}},
					op: AExprArithAdd,
				},
				&AExprSimple{e: &IntegerLiteral{value: 1}},
			},
		},
	},
}
