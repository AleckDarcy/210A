// Code generated from zz/grammar/ZZ.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar // ZZ
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseZZListener is a complete listener for a parse tree produced by ZZParser.
type BaseZZListener struct{}

var _ ZZListener = &BaseZZListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZZListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZZListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZZListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZZListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseZZListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseZZListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterDeclaratorList is called when production declaratorList is entered.
func (s *BaseZZListener) EnterDeclaratorList(ctx *DeclaratorListContext) {}

// ExitDeclaratorList is called when production declaratorList is exited.
func (s *BaseZZListener) ExitDeclaratorList(ctx *DeclaratorListContext) {}

// EnterSimpleTypeSpecifier is called when production simpleTypeSpecifier is entered.
func (s *BaseZZListener) EnterSimpleTypeSpecifier(ctx *SimpleTypeSpecifierContext) {}

// ExitSimpleTypeSpecifier is called when production simpleTypeSpecifier is exited.
func (s *BaseZZListener) ExitSimpleTypeSpecifier(ctx *SimpleTypeSpecifierContext) {}

// EnterListElementTypeSpecifier is called when production listElementTypeSpecifier is entered.
func (s *BaseZZListener) EnterListElementTypeSpecifier(ctx *ListElementTypeSpecifierContext) {}

// ExitListElementTypeSpecifier is called when production listElementTypeSpecifier is exited.
func (s *BaseZZListener) ExitListElementTypeSpecifier(ctx *ListElementTypeSpecifierContext) {}

// EnterListTypeSpecifier is called when production listTypeSpecifier is entered.
func (s *BaseZZListener) EnterListTypeSpecifier(ctx *ListTypeSpecifierContext) {}

// ExitListTypeSpecifier is called when production listTypeSpecifier is exited.
func (s *BaseZZListener) ExitListTypeSpecifier(ctx *ListTypeSpecifierContext) {}

// EnterAExp is called when production aExp is entered.
func (s *BaseZZListener) EnterAExp(ctx *AExpContext) {}

// ExitAExp is called when production aExp is exited.
func (s *BaseZZListener) ExitAExp(ctx *AExpContext) {}

// EnterAExpList is called when production aExpList is entered.
func (s *BaseZZListener) EnterAExpList(ctx *AExpListContext) {}

// ExitAExpList is called when production aExpList is exited.
func (s *BaseZZListener) ExitAExpList(ctx *AExpListContext) {}

// EnterBExp is called when production bExp is entered.
func (s *BaseZZListener) EnterBExp(ctx *BExpContext) {}

// ExitBExp is called when production bExp is exited.
func (s *BaseZZListener) ExitBExp(ctx *BExpContext) {}

// EnterIntegerExpression is called when production integerExpression is entered.
func (s *BaseZZListener) EnterIntegerExpression(ctx *IntegerExpressionContext) {}

// ExitIntegerExpression is called when production integerExpression is exited.
func (s *BaseZZListener) ExitIntegerExpression(ctx *IntegerExpressionContext) {}

// EnterListElement is called when production listElement is entered.
func (s *BaseZZListener) EnterListElement(ctx *ListElementContext) {}

// ExitListElement is called when production listElement is exited.
func (s *BaseZZListener) ExitListElement(ctx *ListElementContext) {}

// EnterListElements is called when production listElements is entered.
func (s *BaseZZListener) EnterListElements(ctx *ListElementsContext) {}

// ExitListElements is called when production listElements is exited.
func (s *BaseZZListener) ExitListElements(ctx *ListElementsContext) {}

// EnterListElementExpression is called when production listElementExpression is entered.
func (s *BaseZZListener) EnterListElementExpression(ctx *ListElementExpressionContext) {}

// ExitListElementExpression is called when production listElementExpression is exited.
func (s *BaseZZListener) ExitListElementExpression(ctx *ListElementExpressionContext) {}

// EnterTupleSizes is called when production tupleSizes is entered.
func (s *BaseZZListener) EnterTupleSizes(ctx *TupleSizesContext) {}

// ExitTupleSizes is called when production tupleSizes is exited.
func (s *BaseZZListener) ExitTupleSizes(ctx *TupleSizesContext) {}

// EnterAssignStatement is called when production assignStatement is entered.
func (s *BaseZZListener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production assignStatement is exited.
func (s *BaseZZListener) ExitAssignStatement(ctx *AssignStatementContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *BaseZZListener) EnterSelectionStatement(ctx *SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *BaseZZListener) ExitSelectionStatement(ctx *SelectionStatementContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *BaseZZListener) EnterIterationStatement(ctx *IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *BaseZZListener) ExitIterationStatement(ctx *IterationStatementContext) {}

// EnterEntry is called when production entry is entered.
func (s *BaseZZListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseZZListener) ExitEntry(ctx *EntryContext) {}

// EnterEntryList is called when production entryList is entered.
func (s *BaseZZListener) EnterEntryList(ctx *EntryListContext) {}

// ExitEntryList is called when production entryList is exited.
func (s *BaseZZListener) ExitEntryList(ctx *EntryListContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BaseZZListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BaseZZListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterTypeSpecifierList is called when production typeSpecifierList is entered.
func (s *BaseZZListener) EnterTypeSpecifierList(ctx *TypeSpecifierListContext) {}

// ExitTypeSpecifierList is called when production typeSpecifierList is exited.
func (s *BaseZZListener) ExitTypeSpecifierList(ctx *TypeSpecifierListContext) {}

// EnterTypeSpecifierWithIdentity is called when production typeSpecifierWithIdentity is entered.
func (s *BaseZZListener) EnterTypeSpecifierWithIdentity(ctx *TypeSpecifierWithIdentityContext) {}

// ExitTypeSpecifierWithIdentity is called when production typeSpecifierWithIdentity is exited.
func (s *BaseZZListener) ExitTypeSpecifierWithIdentity(ctx *TypeSpecifierWithIdentityContext) {}

// EnterTypeSpecifierWithIdentityList is called when production typeSpecifierWithIdentityList is entered.
func (s *BaseZZListener) EnterTypeSpecifierWithIdentityList(ctx *TypeSpecifierWithIdentityListContext) {
}

// ExitTypeSpecifierWithIdentityList is called when production typeSpecifierWithIdentityList is exited.
func (s *BaseZZListener) ExitTypeSpecifierWithIdentityList(ctx *TypeSpecifierWithIdentityListContext) {
}

// EnterFuncSpecifier is called when production funcSpecifier is entered.
func (s *BaseZZListener) EnterFuncSpecifier(ctx *FuncSpecifierContext) {}

// ExitFuncSpecifier is called when production funcSpecifier is exited.
func (s *BaseZZListener) ExitFuncSpecifier(ctx *FuncSpecifierContext) {}

// EnterFuncSpecifierWithName is called when production funcSpecifierWithName is entered.
func (s *BaseZZListener) EnterFuncSpecifierWithName(ctx *FuncSpecifierWithNameContext) {}

// ExitFuncSpecifierWithName is called when production funcSpecifierWithName is exited.
func (s *BaseZZListener) ExitFuncSpecifierWithName(ctx *FuncSpecifierWithNameContext) {}

// EnterFuncReturnPara is called when production funcReturnPara is entered.
func (s *BaseZZListener) EnterFuncReturnPara(ctx *FuncReturnParaContext) {}

// ExitFuncReturnPara is called when production funcReturnPara is exited.
func (s *BaseZZListener) ExitFuncReturnPara(ctx *FuncReturnParaContext) {}

// EnterFuncReturnParaList is called when production funcReturnParaList is entered.
func (s *BaseZZListener) EnterFuncReturnParaList(ctx *FuncReturnParaListContext) {}

// ExitFuncReturnParaList is called when production funcReturnParaList is exited.
func (s *BaseZZListener) ExitFuncReturnParaList(ctx *FuncReturnParaListContext) {}

// EnterFuncStatement is called when production funcStatement is entered.
func (s *BaseZZListener) EnterFuncStatement(ctx *FuncStatementContext) {}

// ExitFuncStatement is called when production funcStatement is exited.
func (s *BaseZZListener) ExitFuncStatement(ctx *FuncStatementContext) {}

// EnterFuncStatementList is called when production funcStatementList is entered.
func (s *BaseZZListener) EnterFuncStatementList(ctx *FuncStatementListContext) {}

// ExitFuncStatementList is called when production funcStatementList is exited.
func (s *BaseZZListener) ExitFuncStatementList(ctx *FuncStatementListContext) {}

// EnterFuncBody is called when production funcBody is entered.
func (s *BaseZZListener) EnterFuncBody(ctx *FuncBodyContext) {}

// ExitFuncBody is called when production funcBody is exited.
func (s *BaseZZListener) ExitFuncBody(ctx *FuncBodyContext) {}

// EnterFuncExpression is called when production funcExpression is entered.
func (s *BaseZZListener) EnterFuncExpression(ctx *FuncExpressionContext) {}

// ExitFuncExpression is called when production funcExpression is exited.
func (s *BaseZZListener) ExitFuncExpression(ctx *FuncExpressionContext) {}

// EnterFuncExpressionWithName is called when production funcExpressionWithName is entered.
func (s *BaseZZListener) EnterFuncExpressionWithName(ctx *FuncExpressionWithNameContext) {}

// ExitFuncExpressionWithName is called when production funcExpressionWithName is exited.
func (s *BaseZZListener) ExitFuncExpressionWithName(ctx *FuncExpressionWithNameContext) {}
