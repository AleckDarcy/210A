package transformer

import (
	"fmt"

	"github.com/AleckDarcy/210A/zz/ast"
)

type FileInfo struct {
	name string
	file *ast.File
}

type Transformer struct {
	files map[string]*FileInfo

	stack *StoreStack
}

func NewTransformer() *Transformer {
	return &Transformer{
		files: map[string]*FileInfo{},
		stack: &StoreStack{KVs: map[string]*StoreMeta{}},
	}
}

func (t *Transformer) AddFile(name string, file *FileInfo) {
	t.files[name] = file
}

func (t *Transformer) EnterBlock() {
	t.stack = &StoreStack{
		KVs:    map[string]*StoreMeta{},
		Parent: t.stack,
	}
}

func (t *Transformer) ExitBlock() {
	t.stack = t.stack.Parent
}

func (t *Transformer) Walk(noder ast.BasicNoder) string {
	switch node := noder.(type) {
	case *ast.AExprArith:
		return t.WalkAExprArith(node)
	case *ast.ListInitExpr:
		return t.WalkListInitExpr(node)
	case *ast.AssignStmt:
		return t.WalkAssignStmt(node)
	case *ast.FuncDefinition:
		return t.WalkFuncDefinition(node)
	case *ast.File:
		return t.WalkFile(node)
	}

	panic("todo Walk")
}

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
	case *ast.ListElementExpr:
		return t.WalkListElementExpr(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkAExprArith(node *ast.AExprArith) string {
	var opStr string
	switch node.Op() {
	case ast.AExprArithAdd:
		opStr = "+"
	case ast.AExprArithMul:
		opStr = "*"
	default:
		panic("todo")
	}

	str1, str2 := t.WalkAExpr(node.E1()), t.WalkAExpr(node.E2())
	if e, ok := node.E1().(*ast.AExprArith); ok && node.Op().PriorTo(e.Op()) {
		str1 = fmt.Sprintf("(%s)", str1)
	}
	if e, ok := node.E2().(*ast.AExprArith); ok && node.Op().PriorTo(e.Op()) {
		str1 = fmt.Sprintf("(%s)", str2)
	}

	return fmt.Sprintf("%s %s %s", str1, opStr, str2)
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
	default:
		panic("todo")
	}

	str1, str2 := t.WalkAExpr(node.E1()), t.WalkAExpr(node.E2())

	return fmt.Sprintf("%s %s %s", str1, opStr, str2)
}

func (t *Transformer) WalkBExprBinary(node *ast.BExprBinary) string {
	var opStr string
	switch node.Op() {
	case ast.BExprBinaryEQ:
		opStr = "=="
	default:
		panic("todo")
	}

	str1, str2 := t.WalkBExpr(node.E1()), t.WalkBExpr(node.E2())

	return fmt.Sprintf("%s %s %s", str1, opStr, str2)
}

func (t *Transformer) WalkTypeSpecifierer(noder ast.TypeSpecifierer) string {
	switch node := noder.(type) {
	case *ast.SimpleTypeSpecifier:
		return node.Name()
	case *ast.ListTypeSpecifier:
		return t.WalkListTypeSpecifier(node)
	case *ast.FuncTypeSpecifier:
		return t.WalkFuncTypeSpecifier(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkListElementExpr(node *ast.ListElementExpr) string {
	var listStr string
	list := node.List()
	for _, index := range list {
		listStr += fmt.Sprintf("[%s]", t.WalkListElementIndex(index))
	}

	return fmt.Sprintf("%s%s", t.WalkIdentifier(node.Name()), listStr)
}

func (t *Transformer) WalkListElementIndex(node *ast.ListElementIndex) string {
	return t.WalkAExpr(node.E())
}

func (t *Transformer) WalkListInitExpr(node *ast.ListInitExpr) string {
	return fmt.Sprintf("make(%s, %s)",
		t.WalkListTypeSpecifier(node.TypeSpecifier()), t.WalkAExpr(node.Size()))
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
	case *ast.ListElementExpr:
		return t.WalkListElementExpr(node)
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
	case *ast.ListInitExpr:
		return t.WalkListInitExpr(node)
	case *ast.FuncInitExpr:
		return t.WalkFuncInitExpr(node)
	case *ast.FuncExecuteExpression:
		return t.WalkFuncExecuteExpression(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkAssignStmt(node *ast.AssignStmt) string {
	if node == nil {
		return ""
	}

	declListStr := t.WalkDeclaratorerList(node.DeclList())
	initListStr := t.WalkAssignIniterList(node.InitList())
	if node.Flag().IsInit() {
		if node.Flag().IsGlobal() {
			return fmt.Sprintf("var %s = %s", declListStr, initListStr)
		}

		return fmt.Sprintf("%s := %s", declListStr, initListStr)
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
	return fmt.Sprintf(""+
		"for %s; %s; %s {\n"+
		"%s\n"+
		"}",
		t.WalkAssignStmt(node.InitStmt()),
		t.WalkBExpr(node.BinExpr()),
		t.WalkAssignStmt(node.IncreStmt()),
		t.WalkFuncStatementerList(node.StmtList()),
	)
}

func (t *Transformer) WalkFuncStatementer(noder ast.FuncStatementer) string {
	switch node := noder.(type) {
	case *ast.AssignStmt:
		return t.WalkAssignStmt(node)
	case *ast.SelectionStmt:
		return t.WalkSelectionStmt(node)
	case *ast.IterationStmt:
		return t.WalkIterationStmt(node)
	case *ast.FuncReturnStatement:
		return t.WalkFuncReturnStatement(node)
	case *ast.FuncExecuteStatement:
		return t.WalkFuncExecuteStatement(node)
	default:
		panic("unreachable code")
	}
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

func (t *Transformer) WalkFuncTypeSpecifier(node *ast.FuncTypeSpecifier) string {
	returnList := node.ReturnList()
	if len(returnList) == 0 {
		return fmt.Sprintf("(%s)",
			t.WalkParaDeclaratorWithIdentityList(node.ParaList()),
		)
	}

	return fmt.Sprintf("(%s) (%s)",
		t.WalkParaDeclaratorWithIdentityList(node.ParaList()),
		t.WalkTypeSpecifiererList(node.ReturnList()),
	)
}

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

func (t *Transformer) WalkFuncInitExpr(node *ast.FuncInitExpr) string {
	return fmt.Sprintf(""+
		"func %s {\n"+
		"%s\n"+
		"}\n",
		t.WalkFuncTypeSpecifier(node.TypeSpecifier()),
		t.WalkFuncStatementerList(node.StmtList()),
	)
}

func (t *Transformer) WalkFuncDefinition(node *ast.FuncDefinition) string {
	result := fmt.Sprintf(""+
		"func %s {\n"+
		"%s\n"+
		"}\n",
		t.WalkFuncTypeSpecifierWithName(node.TypeSpecifier()),
		t.WalkFuncStatementerList(node.StmtList()),
	)

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
		return t.WalkAssignStmt(node)
	case *ast.FuncDefinition:
		return t.WalkFuncDefinition(node)
	default:
		panic("unreachable code")
	}
}

func (t *Transformer) WalkFile(node *ast.File) string {
	return fmt.Sprintf("%s\n", t.WalkDefinitionerList(node.DefList()))
}
