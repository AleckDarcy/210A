package ast

var SimpleTypeSpecifier1 = &SimpleTypeSpecifier{name: "int"} // int

var SimpleTypeSpecifier2 = &SimpleTypeSpecifier{name: "float64"} // float

var AExprSimple1 = &AExprSimple{e: &IntegerLiteral{value: 1}}

var AExprAdd1 = &AExprArith{ // 2 + 3
	e1: &AExprSimple{e: &IntegerLiteral{value: 2}},
	e2: &AExprSimple{e: &IntegerLiteral{value: 3}},
	op: AExprArithAdd,
}

var AExprAdd2 = &AExprArith{ // 1 + (2 + 3)
	e1: &AExprSimple{e: &IntegerLiteral{value: 1}},
	e2: AExprAdd1,
	op: AExprArithAdd,
}

var BExprCompare1 = &BExprCompare{ // 2 + 3 == 5
	e1: AExprAdd1,
	e2: &AExprSimple{e: &IntegerLiteral{value: 5}},
	op: BExprCompareEQ,
}

var BExprCompare2 = &BExprCompare{ // 2 == 2
	e1: &AExprSimple{e: &IntegerLiteral{value: 2}},
	e2: &AExprSimple{e: &IntegerLiteral{value: 2}},
	op: BExprCompareEQ,
}

var BExprBinary1 = &BExprBinary{ // (2 + 3 == 5) == (2 == 2)
	e1: BExprCompare1,
	e2: BExprCompare2,
	op: BExprBinaryEQ,
}

var BExprBinary2 = &BExprBinary{ // true == (2 == 2)
	e1: &BinaryLiteral{value: true},
	e2: BExprCompare2,
	op: BExprBinaryEQ,
}

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

var ListTypeSpecifier3 = &ListTypeSpecifier{ // [][]int
	elem: ListElementTypeSpecifier1,
}

var ListElementExpr1 = &ListElementExpr{ // a[1]
	name: &Identifier{name: "a"},
	list: []*ListElementIndex{
		{e: &AExprSimple{e: &IntegerLiteral{value: 1}}},
	},
}

var ListElementExpr2 = &ListElementExpr{ // b[2 + 3][3]
	name: &Identifier{name: "b"},
	list: []*ListElementIndex{
		{e: AExprAdd1},
		{e: &IntegerLiteral{value: 3}},
	},
}

var ListElementTypeSpecifier1 = &ListElementTypeSpecifier{ // [][]int
	elem: &ListElementTypeSpecifier{
		elem: &SimpleTypeSpecifier{name: "int"},
		typ:  ListElementTypeSpecifierSimple,
	},
	typ: ListElementTypeSpecifierNested,
}

var ListInitExpr1 = &ListInitExpr{ // list([][]int, 2 + 3)
	typeSpecifier: ListTypeSpecifier3,
	size:          AExprAdd1,
}

var ListInitExpr2 = &ListInitExpr{ // list([]int, 4)
	typeSpecifier: &ListTypeSpecifier{
		elem: &ListElementTypeSpecifier{
			elem: SimpleTypeSpecifier1,
		},
	},
	size: &AExprSimple{e: &IntegerLiteral{value: 4}},
}

var AssignStmt1 = &AssignStmt{ // a = list([]int, 4)
	declList: []Declaratorer{
		&Identifier{name: "a"},
	},
	initList: []AssignIniter{
		ListInitExpr2,
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

var AssignStmt4 = &AssignStmt{
	declList: []Declaratorer{
		&Identifier{name: "f"},
	},
	initList: []AssignIniter{
		&FuncInitExpr{
			typeSpecifier: &FuncTypeSpecifier{},
		},
	},
}

var FuncTypeSpecifier1 = &FuncTypeSpecifier{
	paraList:   FuncParaList1,
	returnList: FuncReturnList1,
}

var FuncInitExpr1 = &FuncInitExpr{
	typeSpecifier: FuncTypeSpecifier1,
	stmtList:      FuncBody1,
}

var AssignStmt5 = &AssignStmt{
	declList: []Declaratorer{
		&Identifier{name: "f"},
	},
	initList: []AssignIniter{
		FuncInitExpr1,
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
	if !(a[1] == b) {
		a[1] = a[1] + 1
	}
*/
var IFExpr3 = &IfExpr{
	binExpr: &BExprBinary{
		e1: &BExprCompare{ // !(a[1] == b)
			e1: ListElementExpr1,
			e2: &AExprSimple{e: &Identifier{name: "b"}},
			op: BExprCompareEQ,
		},
		e2: &BinaryLiteral{value: false},
		op: BExprBinaryEQ,
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
	if !(a[1] == b) {
		a[1] = a[1] + 1
	}
*/
var SelectionStmt4 = &SelectionStmt{
	ifExprList: []*IfExpr{
		IFExpr3,
	},
}

/*
	for ;; {}
*/
var IterationStmt1 = &IterationStmt{
	binExpr: &BinaryLiteral{value: true},
}

var IterationAssignStmt1 = &IterationAssignStmt{
	stmt: &AssignStmt{
		declList: []Declaratorer{&Identifier{name: "i"}},
		initList: []AssignIniter{&AExprSimple{e: &IntegerLiteral{value: 1}}},
	},
}

var IterationAssignStmt2 = &IterationAssignStmt{
	stmt: &AssignStmt{
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
	for i = 1; i < 5; i = i + 1 {}
*/
var IterationStmt2 = &IterationStmt{
	initStmt: IterationAssignStmt1.AssignStmt(),
	binExpr: &BExprCompare{
		e1: &AExprSimple{e: &Identifier{name: "i"}},
		e2: &AExprSimple{e: &IntegerLiteral{value: 5}},
		op: BExprCompareLT,
	},
	increStmt: IterationAssignStmt2.AssignStmt(),
}

/*
	func function() {}
*/
var FuncDefinition1 = &FuncDefinition{
	typeSpecifier: &FuncTypeSpecifierWithName{
		name: &Identifier{name: "function"},
	},
}

var FuncTypeSpecifierWithName2 = &FuncTypeSpecifierWithName{
	name:       &Identifier{name: "function1"},
	paraList:   FuncParaList1,
	returnList: FuncReturnList1,
}

/*
	func function1(x, y int, z []float) ([]int, float, int) {
		a = list([]int, 4)
		b = x + y
		return a, b + 1, 1
	}
*/
var FuncDefinition2 = &FuncDefinition{
	typeSpecifier: FuncTypeSpecifierWithName2,
	stmtList:      FuncBody1,
}

var ParaDeclaratorWithIdentity1 = &ParaDeclaratorWithIdentity{
	declList: []*Identifier{
		{name: "x"}, {name: "y"},
	},
	typeSpecifier: SimpleTypeSpecifier1,
}

var ParaDeclaratorWithIdentity2 = &ParaDeclaratorWithIdentity{
	declList:      []*Identifier{{name: "z"}},
	typeSpecifier: ListTypeSpecifier2,
}

var FuncParaList1 = []*ParaDeclaratorWithIdentity{
	ParaDeclaratorWithIdentity1,
	ParaDeclaratorWithIdentity2,
}

var FuncReturnList1 = []TypeSpecifierer{
	ListTypeSpecifier1,
	SimpleTypeSpecifier2,
	SimpleTypeSpecifier1,
}

var FuncReturnStatement1 = &FuncReturnStatement{
	returnList: []FuncReturnParaer{
		&AExprSimple{e: &Identifier{name: "a"}},
		&AExprArith{
			e1: &AExprSimple{e: &Identifier{name: "b"}},
			e2: &AExprSimple{e: &IntegerLiteral{value: 1}},
			op: AExprArithAdd,
		},
		&AExprSimple{e: &IntegerLiteral{value: 1}},
	},
}

var FuncBody1 = []FuncStatementer{
	AssignStmt1,
	AssignStmt3,
	FuncReturnStatement1,
}

var FuncExecutePara1 = &FuncExecutePara{ // 2 + 3
	para: AExprAdd1,
}

var FuncExecuteExpression1 = &FuncExecuteExpression{ // function2(2 + 3)
	name:     &Identifier{"function2"},
	paraList: []FuncExecuteParaer{FuncExecutePara1.para},
}

var FuncExecuteStatement1 = &FuncExecuteStatement{ // function2(2 + 3)
	name:     &Identifier{"function2"},
	paraList: []FuncExecuteParaer{FuncExecutePara1.para},
}
