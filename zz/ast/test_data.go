package ast

var SimpleTypeSpecifier1 = &SimpleTypeSpecifier{name: "int"} // int

var SimpleTypeSpecifier2 = &SimpleTypeSpecifier{name: "float64"} // float

var AExprSimple1 = AExprSimpleHelper.New(&IntegerLiteral{value: 1})

var AExprTranspose1 = ArithTransposeHelper.New(&AExprSimple{e: &IntegerLiteral{value: 2}})

var AExprAdd1 = AExprArithHelper.New( // 2 + 3
	&AExprSimple{e: &IntegerLiteral{value: 2}},
	&AExprSimple{e: &IntegerLiteral{value: 3}},
	AExprArithAdd)

//var AExprAdd1 = &AExprArith{ // 2 + 3
//	e1: &AExprSimple{e: &IntegerLiteral{value: 2}},
//	e2: &AExprSimple{e: &IntegerLiteral{value: 3}},
//	op: AExprArithAdd,
//}

var AExprAdd2 = &AExprArith{ // 1 + (2 + 3)
	e1: &AExprSimple{e: &IntegerLiteral{value: 1}},
	e2: AExprAdd1,
	op: AExprArithAdd,
}

var AExprSub = &AExprArith{ // 2 - (1 + (2 + 3))
	e1: &AExprSimple{e: &IntegerLiteral{value: 2}},
	e2: AExprAdd2,
	op: AExprArithSub,
}

var AExprMul = &AExprArith{ //  2 * (2 - (1 + (2 + 3)))
	e1: &AExprSimple{e: &IntegerLiteral{value: 2}},
	e2: AExprSub,
	op: AExprArithMul,
}

var AExprDiv = &AExprArith{ //  2 * (2 - (1 + (2 + 3))) / 2
	e1: AExprMul,
	e2: &AExprSimple{e: &IntegerLiteral{value: 2}},
	op: AExprArithDiv,
}

var AExprDiv2 = &AExprArith{ // 2 * (2 - (1 + (2 + 3))) / 2.2
	e1: AExprAdd4,
	e2: AExprAdd4,
	op: AExprArithDiv,
}

var AExprAdd3 = &AExprArith{ //  d[1][1] + d[1][1]
	e1: ListElementExpr5,
	e2: ListElementExpr5,
	op: AExprArithAdd,
}

var AExprAdd4 = &AExprArith{ //  2.2 + 2.2
	e1: &AExprSimple{e: &FloatLiteral{value: 2.2}},
	e2: &AExprSimple{e: &FloatLiteral{value: 2.2}},
	op: AExprArithAdd,
}

var BExprCompare1 = BExprCompareHelper.New( // 2 + 3 == 5
	AExprAdd1,
	&AExprSimple{e: &IntegerLiteral{value: 5}},
	BExprCompareEQ)

var BExprCompare2 = &BExprCompare{ // 2 == 2
	e1: &AExprSimple{e: &IntegerLiteral{value: 2}},
	e2: &AExprSimple{e: &IntegerLiteral{value: 2}},
	op: BExprCompareEQ,
}

var BExprCompare3 = &BExprCompare{ // 2 < 3
	e1: &AExprSimple{e: &IntegerLiteral{value: 2}},
	e2: &AExprSimple{e: &IntegerLiteral{value: 3}},
	op: BExprCompareLT,
}

var BExprCompare4 = &BExprCompare{ // 6 > 3
	e1: &AExprSimple{e: &IntegerLiteral{value: 6}},
	e2: &AExprSimple{e: &IntegerLiteral{value: 3}},
	op: BExprCompareGT,
}

var BExprCompare5 = &BExprCompare{ // 3 <= 3
	e1: &AExprSimple{e: &IntegerLiteral{value: 3}},
	e2: &AExprSimple{e: &IntegerLiteral{value: 3}},
	op: BExprCompareLEQ,
}

var BExprCompare6 = &BExprCompare{ // 5 >= 3
	e1: &AExprSimple{e: &IntegerLiteral{value: 5}},
	e2: &AExprSimple{e: &IntegerLiteral{value: 3}},
	op: BExprCompareGEQ,
}

var BExprCompare7 = &BExprCompare{ // 2 * (2 - (1 + (2 + 3))) / 2 != 2 * (2 - (1 + (2 + 3)))
	e1: AExprDiv,
	e2: AExprMul,
	op: BExprCompareNEQ,
}

var BExprBinary1 = BExprBinaryHelper.New( // (2 + 3 == 5) == (2 == 2)
	BExprCompare1,
	BExprCompare2,
	BExprBinaryEQ)

var BExprBinary2 = &BExprBinary{ // true == (2 == 2)
	e1: &BinaryLiteral{value: true},
	e2: BExprCompare2,
	op: BExprBinaryEQ,
}

var BExprBinary3 = &BExprBinary{ // (true == (2 == 2)) || (2 < 3)
	e1: BExprBinary2,
	e2: BExprCompare3,
	op: BExprBinaryOR,
}
var BExprBinary4 = &BExprBinary{ //  ((true == (2 == 2)) != ((true == (2 == 2)) || (2 < 3)))
	e1: BExprBinary2,
	e2: BExprBinary3,
	op: BExprBinaryNEQ,
}
var BExprBinary5 = &BExprBinary{ //  (true && true)
	e1: &BinaryLiteral{value: true},
	e2: &BinaryLiteral{value: true},
	op: BExprBinaryAND,
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

var CollectionElementIndex1 = CollectionElementIndexHelper.New(
	&AExprSimple{&IntegerLiteral{value: 1}}) // 1

var ListElementExpr1 = CollectionElementExprHelper.New( // a[1]
	&Identifier{name: "a"},
	[]*CollectionElementIndex{CollectionElementIndex1})

var ListElementExpr2 = &CollectionElementExpr{ // b[2 + 3][3]
	name: &Identifier{name: "b"},
	list: []*CollectionElementIndex{
		{e: AExprAdd1},
		{e: &IntegerLiteral{value: 3}},
	},
}

var ListElementExpr3 = CollectionElementExprHelper.New( // c[1]
	&Identifier{name: "c"},
	[]*CollectionElementIndex{CollectionElementIndex1},
)

var ListElementExpr4 = CollectionElementExprHelper.New( // d[1]
	&Identifier{name: "d"},
	[]*CollectionElementIndex{CollectionElementIndex1},
)

var ListElementExpr5 = CollectionElementExprHelper.New( // d[1][1]
	&Identifier{name: "d"},
	[]*CollectionElementIndex{CollectionElementIndex1, CollectionElementIndex1},
)

var ListElementExpr6 = CollectionElementExprHelper.New( // e[1]
	&Identifier{name: "e"},
	[]*CollectionElementIndex{CollectionElementIndex1},
)

var ListElementTypeSpecifier1 = &ListElementTypeSpecifier{ // [][]int
	elem: &ListElementTypeSpecifier{
		elem: &SimpleTypeSpecifier{name: "int"},
		typ:  ListElementTypeSpecifierSimple,
	},
	typ: ListElementTypeSpecifierNested,
}

var ListInitExpr1 = ListInitExprHelper.New( // list([][]int, 2 + 3)
	ListTypeSpecifier3,
	AExprAdd1)

var ListInitExpr2 = &ListInitExpr{ // list([]int, 4)
	typeSpecifier: &ListTypeSpecifier{
		elem: &ListElementTypeSpecifier{
			elem: SimpleTypeSpecifier1,
		},
	},
	size: &AExprSimple{e: &IntegerLiteral{value: 4}},
}

var MatrixInitExpr1 = MatrixInitExprHelper.New( // matrix(2,2)
	[]AExpr{&IntegerLiteral{value: 2},
		&IntegerLiteral{value: 2}})

var AssignStmt1 = &AssignStmt{ // a = list([]int, 4)
	declList: []Declaratorer{
		&Declarator{Declaratorer: &Identifier{name: "a"}},
	},
	initList: []AssignIniter{
		ListInitExpr2,
	},
}

var AssignStmt2 = &AssignStmt{ // a[1], b = 2, 3
	declList: []Declaratorer{
		&Declarator{Declaratorer: ListElementExpr1},
		&Declarator{Declaratorer: &Identifier{name: "b"}},
	},
	initList: []AssignIniter{
		&AExprSimple{e: &IntegerLiteral{value: 2}},
		&AExprSimple{e: &IntegerLiteral{value: 3}},
	},
}

var AssignStmt3 = AssignStmtHelper.New( // b := x + y
	AssignStmtFlagInit,
	[]Declaratorer{
		&Declarator{Declaratorer: &Identifier{name: "b"}}},
	[]AssignIniter{&AExprArith{
		e1: &AExprSimple{e: &Identifier{name: "x"}},
		e2: &AExprSimple{e: &Identifier{name: "y"}},
		op: AExprArithAdd}})

var AssignStmt4 = &AssignStmt{ // a := list([]int, 4)
	flag: AssignStmtFlagInit,
	declList: []Declaratorer{
		&Declarator{Declaratorer: &Identifier{name: "a"}},
	},
	initList: []AssignIniter{
		ListInitExpr2,
	},
}

var AssignStmt5 = &AssignStmt{ // b := x + y
	flag: AssignStmtFlagInit | AssignStmtFlagGlobal,
	declList: []Declaratorer{
		&Declarator{Declaratorer: &Identifier{name: "b"}},
	},
	initList: []AssignIniter{
		&AExprArith{
			e1: &AExprSimple{e: &Identifier{name: "x"}},
			e2: &AExprSimple{e: &Identifier{name: "y"}},
			op: AExprArithAdd,
		},
	},
}

var AssignStmt6 = &AssignStmt{ // c := list([]int, 4)
	flag: AssignStmtFlagInit | AssignStmtFlagGlobal,
	declList: []Declaratorer{
		&Declarator{Declaratorer: &Identifier{name: "c"}},
	},
	initList: []AssignIniter{
		ListInitExpr2,
	},
}

var AssignStmt7 = &AssignStmt{ // c[1] = 10
	declList: []Declaratorer{
		&Declarator{Declaratorer: ListElementExpr3},
	},
	initList: []AssignIniter{
		&AExprSimple{e: &IntegerLiteral{value: 10}},
	},
}

var AssignStmt8 = &AssignStmt{ // d := list([][]int, 2+3)
	flag: AssignStmtFlagInit,
	declList: []Declaratorer{
		&Declarator{Declaratorer: &Identifier{name: "d"}},
	},
	initList: []AssignIniter{
		ListInitExpr1,
	},
}

var AssignStmt9 = &AssignStmt{ // d[1] = list([]int, 4)
	declList: []Declaratorer{
		&Declarator{Declaratorer: ListElementExpr4},
	},
	initList: []AssignIniter{
		ListInitExpr2,
	},
}

var AssignStmt10 = &AssignStmt{ // d[1][1] = 10086
	declList: []Declaratorer{
		&Declarator{Declaratorer: ListElementExpr5},
	},
	initList: []AssignIniter{
		&AExprSimple{e: &IntegerLiteral{value: 10086}},
	},
}

var AssignStmt11 = &AssignStmt{ // e := matrix(2,2)
	flag: AssignStmtFlagInit,
	declList: []Declaratorer{
		&Declarator{Declaratorer: &Identifier{name: "e"}},
	},
	initList: []AssignIniter{
		MatrixInitExpr1,
	},
}

var AssignStmt12 = &AssignStmt{ // h := 12
	flag: AssignStmtFlagInit,
	declList: []Declaratorer{
		&Declarator{Declaratorer: &Identifier{name: "h"}},
	},
	initList: []AssignIniter{
		&AExprSimple{e: &IntegerLiteral{value: 12}},
	},
}

var AssignStmt13 = &AssignStmt{ // b13, c13 := 2, 3
	flag: AssignStmtFlagInit,
	declList: []Declaratorer{
		&Declarator{Declaratorer: &Identifier{name: "b13"}},
		&Declarator{Declaratorer: &Identifier{name: "c13"}},
	},
	initList: []AssignIniter{
		&AExprSimple{e: &IntegerLiteral{value: 2}},
		&AExprSimple{e: &IntegerLiteral{value: 3}},
	},
}

var PrintStatement1 = PrintStatementHelper.New(
	[]BasicNoder{AExprDiv2},
)

//var AssignStmt4 = &AssignStmt{
//	declList: []Declaratorer{
//		&Identifier{name: "f"},
//	},
//	initList: []AssignIniter{
//		&FuncInitExpr{
//			typeSpecifier: &FuncTypeSpecifier{},
//		},
//	},
//}
//
//var FuncTypeSpecifier1 = &FuncTypeSpecifier{
//	paraList:   FuncParaList1,
//	returnList: FuncReturnList1,
//}
//
//var FuncInitExpr1 = &FuncInitExpr{
//	typeSpecifier: FuncTypeSpecifier1,
//	stmtList:      FuncBody1,
//}
//
//var AssignStmt5 = &AssignStmt{
//	declList: []Declaratorer{
//		&Identifier{name: "f"},
//	},
//	initList: []AssignIniter{
//		FuncInitExpr1,
//	},
//}

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
				&Declarator{Declaratorer: ListElementExpr1},
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
				&Declarator{Declaratorer: ListElementExpr1},
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
				&Declarator{Declaratorer: ListElementExpr1},
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
	if 2 * (2 - (1 + (2 + 3))) / 2 != 2 * (2 - (1 + (2 + 3))) {}
*/
var IFExpr4 = &IfExpr{
	binExpr: BExprCompare7,
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
				&Declarator{Declaratorer: ListElementExpr1},
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
	if 2 * (2 - (1 + (2 + 3))) / 2 != 2 * (2 - (1 + (2 + 3))) {}
*/
var SelectionStmt5 = &SelectionStmt{
	ifExprList: []*IfExpr{
		IFExpr4,
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
		declList: []Declaratorer{&Declarator{Declaratorer: &Identifier{name: "i"}}},
		initList: []AssignIniter{&AExprSimple{e: &IntegerLiteral{value: 1}}},
	},
}

var IterationAssignStmt2 = &IterationAssignStmt{
	stmt: &AssignStmt{
		declList: []Declaratorer{&Declarator{Declaratorer: &Identifier{name: "i"}}},
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

var FuncTypeSpecifierWithName3 = &FuncTypeSpecifierWithName{
	name:     &Identifier{name: "function_matrix"},
	paraList: FuncParaList2,
}

var FuncTypeSpecifierWithName4 = &FuncTypeSpecifierWithName{
	name:       &Identifier{name: "function2"},
	paraList:   FuncParaList3,
	returnList: FuncReturnList2,
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

var FuncDefinition3 = &FuncDefinition{
	typeSpecifier: FuncTypeSpecifierWithName3,
}

var FuncDefinition4 = &FuncDefinition{
	typeSpecifier: FuncTypeSpecifierWithName4,
	stmtList:      FuncBody2,
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

var ParaDeclaratorWithIdentity3 = &ParaDeclaratorWithIdentity{
	declList:      []*Identifier{{name: "mat"}},
	typeSpecifier: SimpleTypeSpecifierHelper.New("matrix"),
}

var FuncParaList1 = []*ParaDeclaratorWithIdentity{
	ParaDeclaratorWithIdentity1,
	ParaDeclaratorWithIdentity2,
}

var FuncParaList2 = []*ParaDeclaratorWithIdentity{
	ParaDeclaratorWithIdentity3,
}

var FuncParaList3 = []*ParaDeclaratorWithIdentity{
	ParaDeclaratorWithIdentity1,
}

var FuncReturnList1 = []TypeSpecifierer{
	ListTypeSpecifier1,
	SimpleTypeSpecifier2,
	SimpleTypeSpecifier1,
}

var FuncReturnList2 = []TypeSpecifierer{
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
	AssignStmt4,
	AssignStmt3,
	FuncReturnStatement1,
}

var FuncBody2 = []FuncStatementer{
	AssignStmt3,
	&FuncReturnStatement{
		returnList: []FuncReturnParaer{
			&AExprSimple{e: &Identifier{name: "b"}},
		},
	},
}

var FuncExecutePara1 = &FuncExecutePara{ // 2 + 3
	para: AExprAdd1,
}

var FuncExecuteExpression1 = &FuncExecuteExpression{ // function2(2 + 3)
	name: &Identifier{"function2"},
	paraList: []FuncExecuteParaer{FuncExecutePara1.para,
		AExprSimpleHelper.New(IntegerLiteralHelper.New(1)),
	},
}

var FuncExecuteStatement1 = &FuncExecuteStatement{ // function2(2 + 3, 1)
	name: &Identifier{"function2"},
	paraList: []FuncExecuteParaer{FuncExecutePara1.para,
		AExprSimpleHelper.New(IntegerLiteralHelper.New(1)),
	},
}
