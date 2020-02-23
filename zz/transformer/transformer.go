package transformer

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/AleckDarcy/210A/zz/ast"
)

type FileInfo struct {
	name string
	file *ast.File
}

type Transformer struct {
	files map[string]*FileInfo

	checker *checker
	//stack StoreStack
}

func NewTransformer() *Transformer {
	return &Transformer{
		files: map[string]*FileInfo{},

		checker: &checker{
			StoreStack: &StoreStack{variants: map[string]*VariantInfo{}},
		},
		//stack: StoreStack{variants: map[string]*VariantInfo{}},
	}
}

//func (t *Transformer) AddFile(name string, file *FileInfo) {
//	t.files[name] = file
//}

func (t *Transformer) WalkIdentifier(node *ast.Identifier) string {
	return node.Name()
}

func (t *Transformer) WalkIntegerLiteral(node *ast.IntegerLiteral) string {
	return fmt.Sprintf("%d", node.Value())
}

func (t *Transformer) WalkFloatLiteral(node *ast.FloatLiteral) string {
	return fmt.Sprintf("%f", node.Value())
}

func (t *Transformer) WalkBinaryLiteral(node *ast.BinaryLiteral) string {
	return fmt.Sprintf("%v", node.Value())
}

func (t *Transformer) WalkAExpr(noder ast.AExpr) string {
	switch node := noder.(type) {
	case *ast.Identifier:
		return t.WalkIdentifier(node)
	case *ast.IntegerLiteral:
		return t.WalkIntegerLiteral(node)
	case *ast.FloatLiteral:
		return t.WalkFloatLiteral(node)
	case *ast.AExprSimple:
		return t.WalkAExpr(node.E())
	case *ast.AExprArith:
		return t.WalkAExprArith(node)
	case *ast.ArithTranspose:
		return t.WalkArithTranspose(node)
	case *ast.CollectionElementExpr:
		return t.WalkCollectionElementExpr(node, false)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkAExprArith(node *ast.AExprArith) string {
	var opStr string
	switch node.Op() {
	case ast.AExprArithAdd:
		opStr = "+"
	case ast.AExprArithSub:
		opStr = "-"
	case ast.AExprArithMul:
		opStr = "*"
	case ast.AExprArithDiv:
		opStr = "/"
	default:
		panic("undefined")
	}

	// todo
	type1, _ := Checker.CheckAExprType(node.E1())
	type2, _ := Checker.CheckAExprType(node.E2())

	str1, str2 := t.WalkAExpr(node.E1()), t.WalkAExpr(node.E2())
	if type1.IsMatrix() && type2.IsMatrix() {
		return fmt.Sprintf("%s.MulMatrix(%s)", str1, str2)
	} else if type1.IsMatrix() {
		if type2.IsInterger() || type2.IsFloat() {
			return fmt.Sprintf("%s.MulFloat(runtime.FloatData(%s))", str1, str2)
		} else {
			panic("")
		}
	} else if type2.IsMatrix() {
		if type1.IsInterger() || type1.IsFloat() {
			return fmt.Sprintf("%s.MulFloat(runtime.FloatData(%s))", str2, str1)
		} else {
			panic("")
		}
	}

	if e, ok := node.E1().(*ast.AExprArith); ok && node.Op().PriorTo(e.Op()) {
		str1 = fmt.Sprintf("(%s)", str1)
	}
	if e, ok := node.E2().(*ast.AExprArith); ok && node.Op().PriorTo(e.Op()) {
		str2 = fmt.Sprintf("(%s)", str2)
	}

	return fmt.Sprintf("%s %s %s", str1, opStr, str2)
}

func (t *Transformer) WalkArithTranspose(node *ast.ArithTranspose) string {
	return fmt.Sprintf("%s.Transpose()", t.WalkAExpr(node.E()))
}

func (t *Transformer) WalkBExpr(noder ast.BExpr) string {
	switch node := noder.(type) {
	case *ast.BinaryLiteral:
		return t.WalkBinaryLiteral(node)
	case *ast.BExprCompare:
		return t.WalkBExprCompare(node)
	case *ast.BExprBinary:
		return t.WalkBExprBinary(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkBExprCompare(node *ast.BExprCompare) string {
	var opStr string
	switch node.Op() {
	case ast.BExprCompareEQ:
		opStr = "=="
	case ast.BExprCompareLT:
		opStr = "<"
	case ast.BExprCompareLEQ:
		opStr = "<="
	case ast.BExprCompareGT:
		opStr = ">"
	case ast.BExprCompareGEQ:
		opStr = ">="
	case ast.BExprCompareNEQ:
		opStr = "!="
	default:
		panic("undefined")
	}

	str1, str2 := t.WalkAExpr(node.E1()), t.WalkAExpr(node.E2())

	return fmt.Sprintf("%s %s %s", str1, opStr, str2)
}

func (t *Transformer) WalkBExprBinary(node *ast.BExprBinary) string {
	var opStr string
	switch node.Op() {
	case ast.BExprBinaryEQ:
		opStr = "=="
	case ast.BExprBinaryNEQ:
		opStr = "!="
	case ast.BExprBinaryAND:
		opStr = "&&"
	case ast.BExprBinaryOR:
		opStr = "||"
	default:
		panic("undefined")
	}

	str1, str2 := t.WalkBExpr(node.E1()), t.WalkBExpr(node.E2())

	if _, ok := node.E1().(*ast.BExprCompare); ok {
		str1 = fmt.Sprintf("(%s)", str1)
	}
	if _, ok := node.E2().(*ast.BExprCompare); ok {
		str2 = fmt.Sprintf("(%s)", str2)
	}
	return fmt.Sprintf("%s %s %s", str1, opStr, str2)
}

func (t *Transformer) WalkTypeSpecifierer(noder ast.TypeSpecifierer) string {
	switch node := noder.(type) {
	case *ast.SimpleTypeSpecifier:
		if node.Name() == "matrix" {
			return "*runtime.Data"
		}

		return node.Name()
	case *ast.ListTypeSpecifier:
		return t.WalkListTypeSpecifier(node)
	//case *ast.FuncTypeSpecifier:
	//	return t.WalkFuncTypeSpecifier(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkCollectionElementExpr(node *ast.CollectionElementExpr, assignFlag bool) string {
	var listStr string

	typ, _, err := Checker.CheckCollectionElementExprType(node)
	if err != nil {
		panic(err)
	}
	list := node.List()

	if typ.IsList() {
		for _, index := range list {
			listStr += fmt.Sprintf("[%s]", t.WalkCollectionElementIndex(index))
		}
	} else {
		for i, index := range list {
			if i == len(list)-1 && assignFlag {
				listStr += fmt.Sprintf(".Set(%s, %%s)", t.WalkCollectionElementIndex(index))
			} else {
				listStr += fmt.Sprintf(".Get(%s)", t.WalkCollectionElementIndex(index))
			}
		}
	}

	return fmt.Sprintf("%s%s", t.WalkIdentifier(node.Name()), listStr)
}

func (t *Transformer) WalkCollectionElementIndex(node *ast.CollectionElementIndex) string {
	return t.WalkAExpr(node.E())
}

func (t *Transformer) WalkListInitExpr(node *ast.ListInitExpr) string {
	return fmt.Sprintf("make(%s, %s)",
		t.WalkListTypeSpecifier(node.TypeSpecifier()), t.WalkAExpr(node.Size()))
}

func (t *Transformer) WalkMatrixInitExpr(node *ast.MatrixInitExpr) string {
	return fmt.Sprintf("runtime.New(%s)", t.WalkAExprList(node.Sizes()))
	//return fmt.Sprintf("make(%s, %s)",
	//	t.WalkListTypeSpecifier(node.TypeSpecifier()), t.WalkAExpr(node.Size()))
}

func (t *Transformer) WalkListTypeSpecifier(node *ast.ListTypeSpecifier) string {
	return t.WalkListElementTypeSpecifier(node.Elem())
}

func (t *Transformer) WalkListElementTypeSpecifier(node *ast.ListElementTypeSpecifier) string {
	return t.WalkListElementTypeSpecifierer(node.Elem())
}

func (t *Transformer) WalkListElementTypeSpecifierer(noder ast.ListElementTypeSpecifierer) string {
	switch node := noder.(type) {
	case *ast.SimpleTypeSpecifier:
		return fmt.Sprintf("[]%s", node.Name())
	case *ast.ListElementTypeSpecifier:
		return fmt.Sprintf("[]%s", t.WalkListElementTypeSpecifierer(node.Elem()))
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkDeclaratorer(noder ast.Declaratorer) string {
	switch node := noder.(type) {
	case *ast.Identifier:
		return t.WalkIdentifier(node)
	case *ast.CollectionElementExpr:
		return t.WalkCollectionElementExpr(node, true)
	case *ast.Declarator:
		return t.WalkDeclaratorer(node.Declaratorer)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkAssignIniter(noder ast.AssignIniter) string {
	switch node := noder.(type) {
	case *ast.IntegerLiteral:
		return t.WalkIntegerLiteral(node)
	case *ast.FloatLiteral:
		return t.WalkFloatLiteral(node)
	case *ast.AExprSimple:
		return t.WalkAExpr(node.E())
	case *ast.AExprArith:
		return t.WalkAExprArith(node)
	case *ast.ArithTranspose:
		return t.WalkArithTranspose(node)
	case *ast.ListInitExpr:
		return t.WalkListInitExpr(node)
	case *ast.MatrixInitExpr:
		return t.WalkMatrixInitExpr(node)
	//case *ast.FuncInitExpr:
	//	return t.WalkFuncInitExpr(node)
	case *ast.FuncExecuteExpression:
		return t.WalkFuncExecuteExpression(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) ProtectAssignStmt(node *ast.AssignStmt) string {
	var protectAssignee, protectAssignor string

	for i, decl := range node.DeclList() {
		name := decl.Identifier().Name()
		if i == 0 {
			protectAssignee = "_"
			protectAssignor = name
		} else {
			protectAssignee += ", _"
			protectAssignor += ", " + name
		}
	}

	return fmt.Sprintf("\n%s = %s", protectAssignee, protectAssignor)
}

func (t *Transformer) ProtectCurrentBlock() string {
	var protectAssignee, protectAssignor string
	protectCount := 0

	for name, info := range Checker.GetCurrentStack().variants {
		if !info.used {
			if protectCount == 0 {
				protectAssignee = "_"
				protectAssignor = name
			} else {
				protectAssignee += ", _"
				protectAssignor += ", " + name
			}

			protectCount++
		}
	}

	if protectCount > 0 {
		return fmt.Sprintf("%s = %s\n", protectAssignee, protectAssignor)
	}

	return ""
}

func (t *Transformer) WalkAssignStmt(node *ast.AssignStmt, protectFlag bool) string {
	if node == nil {
		return ""
	}

	declTypeList, initTypeList, err := Checker.CheckAssignStmtType(node)
	if err != nil {
		panic(err)
	}

	var protect string
	if protectFlag {
		protect = t.ProtectAssignStmt(node)
	}

	declList, initList := node.DeclList(), node.InitList()
	declListStr := t.WalkDeclaratorerList(declList)
	initListStr := t.WalkAssignIniterList(initList)

	if !reflect.DeepEqual(declTypeList, initTypeList) {
		var tmpDeclListStr, castTmpDeclListStr string
		for i, decl := range declList {
			if strings.Index(declListStr, "%s") != -1 { // todo: add flag to AssignStmt
				return fmt.Sprintf(declListStr, "runtime.FloatData("+initListStr+")")
			}

			castStr := "zz_tmp_" + decl.Identifier().Name()
			if i == 0 {
				tmpDeclListStr = castStr
			} else {
				tmpDeclListStr += ", " + castStr
			}

			if declTypeList[i] != initTypeList[i] {
				if declTypeList[i].IsInterger() {
					castStr = fmt.Sprintf("int64(%s)", castStr)
				} else if declTypeList[i].IsFloat() {
					castStr = fmt.Sprintf("float64(%s)", castStr)
				}
			}

			if i == 0 {
				castTmpDeclListStr = castStr
			} else {
				castTmpDeclListStr += ", " + castStr
			}
		}

		return fmt.Sprintf(""+
			"{\n"+
			"	%s := %s\n"+
			"	%s = %s\n"+
			"}", tmpDeclListStr, initListStr, declListStr, castTmpDeclListStr,
		)
	}

	if strings.Index(declListStr, "%s") != -1 { // todo: add flag to AssignStmt
		return fmt.Sprintf(declListStr, initListStr)
	}

	if node.Flag().IsInit() {
		if node.Flag().IsGlobal() {
			return fmt.Sprintf("var %s = %s", declListStr, initListStr)
		}

		return fmt.Sprintf(""+
			"%s := %s"+
			"%s", declListStr, initListStr, protect)
	}

	return fmt.Sprintf("%s = %s", declListStr, initListStr)
}

func (t *Transformer) WalkIfExpr(node *ast.IfExpr) string {
	return fmt.Sprintf(""+
		"if %s {\n"+
		"%s\n"+
		"}",
		t.WalkBExpr(node.BinExpr()),
		t.WalkFuncStatementerList(node.StmtList()),
	)
}

func (t *Transformer) WalkElseExpr(node *ast.ElseExpr) string {
	if node == nil {
		return ""
	}

	return fmt.Sprintf(""+
		" else {\n"+
		"%s\n"+
		"}",
		t.WalkFuncStatementerList(node.StmtList()),
	)
}

func (t *Transformer) WalkSelectionStmt(node *ast.SelectionStmt) string {
	return fmt.Sprintf("%s%s",
		t.WalkIfExprList(node.IfExprList()),
		t.WalkElseExpr(node.ElseExpr()),
	)
}

func (t *Transformer) WalkIterationStmt(node *ast.IterationStmt) string {
	Checker.EnterBlock()

	result := fmt.Sprintf(""+
		"for %s; %s; %s {\n"+
		"%s"+
		"%s\n"+
		"}",
		t.WalkAssignStmt(node.InitStmt(), false),
		t.WalkBExpr(node.BinExpr()),
		t.WalkAssignStmt(node.IncreStmt(), false),
		t.ProtectCurrentBlock(),
		t.WalkFuncStatementerList(node.StmtList()),
	)

	Checker.ExitBlock()

	return result
}

func (t *Transformer) WalkFuncStatementer(noder ast.FuncStatementer) string {
	switch node := noder.(type) {
	case *ast.AssignStmt:
		return t.WalkAssignStmt(node, true)
	case *ast.SelectionStmt:
		return t.WalkSelectionStmt(node)
	case *ast.IterationStmt:
		return t.WalkIterationStmt(node)
	case *ast.FuncReturnStatement:
		return t.WalkFuncReturnStatement(node)
	case *ast.FuncExecuteStatement:
		return t.WalkFuncExecuteStatement(node)
	case *ast.PrintStatement:
		return t.WalkPrintStatement(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkPrintStatement(node *ast.PrintStatement) string {
	var listStr string
	for i, basicNode := range node.List() {
		var tmp string
		switch elem := basicNode.(type) {
		case ast.AExpr:
			tmp = t.WalkAExpr(elem)
		case ast.BExpr:
			tmp = t.WalkBExpr(elem)
		default:
			panic("unreachable code")
		}

		if i == 0 {
			listStr = tmp
		} else {
			listStr += ", " + tmp
		}
	}

	return fmt.Sprintf("fmt.Println(%s)", listStr)
}

func (t *Transformer) WalkParaDeclaratorWithIdentity(node *ast.ParaDeclaratorWithIdentity) string {
	var declListStr string
	for i, decl := range node.DeclList() {
		if i == 0 {
			declListStr = t.WalkIdentifier(decl)
		} else {
			declListStr += fmt.Sprintf(", %s", t.WalkIdentifier(decl))
		}
	}

	return fmt.Sprintf("%s %s",
		declListStr,
		t.WalkTypeSpecifierer(node.TypeSpecifier()),
	)
}

//func (t *Transformer) WalkFuncTypeSpecifier(node *ast.FuncTypeSpecifier) string {
//	returnList := node.ReturnList()
//	if len(returnList) == 0 {
//		return fmt.Sprintf("(%s)",
//			t.WalkParaDeclaratorWithIdentityList(node.ParaList()),
//		)
//	}
//
//	return fmt.Sprintf("(%s) (%s)",
//		t.WalkParaDeclaratorWithIdentityList(node.ParaList()),
//		t.WalkTypeSpecifiererList(node.ReturnList()),
//	)
//}

func (t *Transformer) WalkFuncTypeSpecifierWithName(node *ast.FuncTypeSpecifierWithName) string {
	returnList := node.ReturnList()
	if len(returnList) == 0 {
		return fmt.Sprintf("%s(%s)",
			t.WalkIdentifier(node.Name()),
			t.WalkParaDeclaratorWithIdentityList(node.ParaList()),
		)
	}

	return fmt.Sprintf("%s(%s) (%s)",
		t.WalkIdentifier(node.Name()),
		t.WalkParaDeclaratorWithIdentityList(node.ParaList()),
		t.WalkTypeSpecifiererList(returnList),
	)
}

func (t *Transformer) WalkFuncReturnParaer(noder ast.FuncReturnParaer) string {
	switch node := noder.(type) {
	case *ast.Identifier:
		return t.WalkIdentifier(node)
	case *ast.AExprSimple:
		return t.WalkAExpr(node.E())
	case *ast.AExprArith:
		return t.WalkAExprArith(node)
	case *ast.BExprCompare:
		return t.WalkBExprCompare(node)
	case *ast.BExprBinary:
		return t.WalkBExprBinary(node)
	case *ast.FuncExecuteExpression:
		return t.WalkFuncExecuteExpression(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkFuncReturnStatement(node *ast.FuncReturnStatement) string {
	returnList := node.ReturnList()
	if len(returnList) == 0 {
		return "return"
	}

	return fmt.Sprintf("return %s", t.WalkFuncReturnParaerList(returnList))
}

//func (t *Transformer) WalkFuncInitExpr(node *ast.FuncInitExpr) string {
//	return fmt.Sprintf(""+
//		"func %s {\n"+
//		"%s\n"+
//		"}\n",
//		t.WalkFuncTypeSpecifier(node.TypeSpecifier()),
//		t.WalkFuncStatementerList(node.StmtList()),
//	)
//}

func (t *Transformer) WalkFuncDefinition(node *ast.FuncDefinition) string {
	Checker.EnterBlock()

	// generate code
	funcTypeStr := t.WalkFuncTypeSpecifierWithName(node.TypeSpecifier())
	stmtListStr := t.WalkFuncStatementerList(node.StmtList())

	result := fmt.Sprintf(""+
		"func %s {\n"+
		"%s\n"+
		"}\n",
		funcTypeStr,
		stmtListStr,
	)

	Checker.ExitBlock()

	return result
}

func (t *Transformer) WalkFuncExecuteParaer(noder ast.FuncExecuteParaer) string {
	switch node := noder.(type) {
	case *ast.AExprSimple:
		return t.WalkAExpr(node.E())
	case *ast.AExprArith:
		return t.WalkAExprArith(node)
	case *ast.BExprCompare:
		return t.WalkBExprCompare(node)
	case *ast.BExprBinary:
		return t.WalkBExprBinary(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkFuncExecuteExpression(node *ast.FuncExecuteExpression) string {
	return fmt.Sprintf("%s(%s)",
		t.WalkIdentifier(node.Name()),
		t.WalkFuncExecuteParaerList(node.ParaList()),
	)
}

func (t *Transformer) WalkFuncExecuteStatement(node *ast.FuncExecuteStatement) string {
	return fmt.Sprintf("%s(%s)",
		t.WalkIdentifier(node.Name()),
		t.WalkFuncExecuteParaerList(node.ParaList()),
	)
}

func (t *Transformer) WalkDefinitioner(noder ast.Definitioner) string {
	switch node := noder.(type) {
	case *ast.AssignStmt:
		return t.WalkAssignStmt(node, true)
	case *ast.FuncDefinition:
		return t.WalkFuncDefinition(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkFile(node *ast.File) string {
	return fmt.Sprintf("%s\n", t.WalkDefinitionerList(node.DefList()))
}
