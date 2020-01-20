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

// EnterListInitExpression is called when production listInitExpression is entered.
func (s *BaseZZListener) EnterListInitExpression(ctx *ListInitExpressionContext) {}

// ExitListInitExpression is called when production listInitExpression is exited.
func (s *BaseZZListener) ExitListInitExpression(ctx *ListInitExpressionContext) {}

// EnterAExp_bracketExpression is called when production aExp_bracketExpression is entered.
func (s *BaseZZListener) EnterAExp_bracketExpression(ctx *AExp_bracketExpressionContext) {}

// ExitAExp_bracketExpression is called when production aExp_bracketExpression is exited.
func (s *BaseZZListener) ExitAExp_bracketExpression(ctx *AExp_bracketExpressionContext) {}

// EnterAExp_FloatLiteral is called when production aExp_FloatLiteral is entered.
func (s *BaseZZListener) EnterAExp_FloatLiteral(ctx *AExp_FloatLiteralContext) {}

// ExitAExp_FloatLiteral is called when production aExp_FloatLiteral is exited.
func (s *BaseZZListener) ExitAExp_FloatLiteral(ctx *AExp_FloatLiteralContext) {}

// EnterAExp_multiplicativeExpression is called when production aExp_multiplicativeExpression is entered.
func (s *BaseZZListener) EnterAExp_multiplicativeExpression(ctx *AExp_multiplicativeExpressionContext) {
}

// ExitAExp_multiplicativeExpression is called when production aExp_multiplicativeExpression is exited.
func (s *BaseZZListener) ExitAExp_multiplicativeExpression(ctx *AExp_multiplicativeExpressionContext) {
}

// EnterAExp_additiveExpression is called when production aExp_additiveExpression is entered.
func (s *BaseZZListener) EnterAExp_additiveExpression(ctx *AExp_additiveExpressionContext) {}

// ExitAExp_additiveExpression is called when production aExp_additiveExpression is exited.
func (s *BaseZZListener) ExitAExp_additiveExpression(ctx *AExp_additiveExpressionContext) {}

// EnterAExp_Identifier is called when production aExp_Identifier is entered.
func (s *BaseZZListener) EnterAExp_Identifier(ctx *AExp_IdentifierContext) {}

// ExitAExp_Identifier is called when production aExp_Identifier is exited.
func (s *BaseZZListener) ExitAExp_Identifier(ctx *AExp_IdentifierContext) {}

// EnterAExp_IntergerLiteral is called when production aExp_IntergerLiteral is entered.
func (s *BaseZZListener) EnterAExp_IntergerLiteral(ctx *AExp_IntergerLiteralContext) {}

// ExitAExp_IntergerLiteral is called when production aExp_IntergerLiteral is exited.
func (s *BaseZZListener) ExitAExp_IntergerLiteral(ctx *AExp_IntergerLiteralContext) {}

// EnterAExp_listElementExpression is called when production aExp_listElementExpression is entered.
func (s *BaseZZListener) EnterAExp_listElementExpression(ctx *AExp_listElementExpressionContext) {}

// ExitAExp_listElementExpression is called when production aExp_listElementExpression is exited.
func (s *BaseZZListener) ExitAExp_listElementExpression(ctx *AExp_listElementExpressionContext) {}

// EnterAExprList is called when production aExprList is entered.
func (s *BaseZZListener) EnterAExprList(ctx *AExprListContext) {}

// ExitAExprList is called when production aExprList is exited.
func (s *BaseZZListener) ExitAExprList(ctx *AExprListContext) {}

// EnterBExp is called when production bExp is entered.
func (s *BaseZZListener) EnterBExp(ctx *BExpContext) {}

// ExitBExp is called when production bExp is exited.
func (s *BaseZZListener) ExitBExp(ctx *BExpContext) {}

// EnterIntegerExpression is called when production integerExpression is entered.
func (s *BaseZZListener) EnterIntegerExpression(ctx *IntegerExpressionContext) {}

// ExitIntegerExpression is called when production integerExpression is exited.
func (s *BaseZZListener) ExitIntegerExpression(ctx *IntegerExpressionContext) {}

// EnterListElementIndex is called when production listElementIndex is entered.
func (s *BaseZZListener) EnterListElementIndex(ctx *ListElementIndexContext) {}

// ExitListElementIndex is called when production listElementIndex is exited.
func (s *BaseZZListener) ExitListElementIndex(ctx *ListElementIndexContext) {}

// EnterListElementIndexList is called when production listElementIndexList is entered.
func (s *BaseZZListener) EnterListElementIndexList(ctx *ListElementIndexListContext) {}

// ExitListElementIndexList is called when production listElementIndexList is exited.
func (s *BaseZZListener) ExitListElementIndexList(ctx *ListElementIndexListContext) {}

// EnterListElementExpression is called when production listElementExpression is entered.
func (s *BaseZZListener) EnterListElementExpression(ctx *ListElementExpressionContext) {}

// ExitListElementExpression is called when production listElementExpression is exited.
func (s *BaseZZListener) ExitListElementExpression(ctx *ListElementExpressionContext) {}

// EnterAssignInit is called when production assignInit is entered.
func (s *BaseZZListener) EnterAssignInit(ctx *AssignInitContext) {}

// ExitAssignInit is called when production assignInit is exited.
func (s *BaseZZListener) ExitAssignInit(ctx *AssignInitContext) {}

// EnterAssignInitList is called when production assignInitList is entered.
func (s *BaseZZListener) EnterAssignInitList(ctx *AssignInitListContext) {}

// ExitAssignInitList is called when production assignInitList is exited.
func (s *BaseZZListener) ExitAssignInitList(ctx *AssignInitListContext) {}

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

// EnterParaDeclaratorList is called when production paraDeclaratorList is entered.
func (s *BaseZZListener) EnterParaDeclaratorList(ctx *ParaDeclaratorListContext) {}

// ExitParaDeclaratorList is called when production paraDeclaratorList is exited.
func (s *BaseZZListener) ExitParaDeclaratorList(ctx *ParaDeclaratorListContext) {}

// EnterParaDeclaratorWithIdentity is called when production paraDeclaratorWithIdentity is entered.
func (s *BaseZZListener) EnterParaDeclaratorWithIdentity(ctx *ParaDeclaratorWithIdentityContext) {}

// ExitParaDeclaratorWithIdentity is called when production paraDeclaratorWithIdentity is exited.
func (s *BaseZZListener) ExitParaDeclaratorWithIdentity(ctx *ParaDeclaratorWithIdentityContext) {}

// EnterParaDeclaratorWithIdentityList is called when production paraDeclaratorWithIdentityList is entered.
func (s *BaseZZListener) EnterParaDeclaratorWithIdentityList(ctx *ParaDeclaratorWithIdentityListContext) {
}

// ExitParaDeclaratorWithIdentityList is called when production paraDeclaratorWithIdentityList is exited.
func (s *BaseZZListener) ExitParaDeclaratorWithIdentityList(ctx *ParaDeclaratorWithIdentityListContext) {
}

// EnterFuncTypeSpecifier is called when production funcTypeSpecifier is entered.
func (s *BaseZZListener) EnterFuncTypeSpecifier(ctx *FuncTypeSpecifierContext) {}

// ExitFuncTypeSpecifier is called when production funcTypeSpecifier is exited.
func (s *BaseZZListener) ExitFuncTypeSpecifier(ctx *FuncTypeSpecifierContext) {}

// EnterFuncTypeSpecifierWithName is called when production funcTypeSpecifierWithName is entered.
func (s *BaseZZListener) EnterFuncTypeSpecifierWithName(ctx *FuncTypeSpecifierWithNameContext) {}

// ExitFuncTypeSpecifierWithName is called when production funcTypeSpecifierWithName is exited.
func (s *BaseZZListener) ExitFuncTypeSpecifierWithName(ctx *FuncTypeSpecifierWithNameContext) {}

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

// EnterFuncInitExpression is called when production funcInitExpression is entered.
func (s *BaseZZListener) EnterFuncInitExpression(ctx *FuncInitExpressionContext) {}

// ExitFuncInitExpression is called when production funcInitExpression is exited.
func (s *BaseZZListener) ExitFuncInitExpression(ctx *FuncInitExpressionContext) {}

// EnterFuncExpression is called when production funcExpression is entered.
func (s *BaseZZListener) EnterFuncExpression(ctx *FuncExpressionContext) {}

// ExitFuncExpression is called when production funcExpression is exited.
func (s *BaseZZListener) ExitFuncExpression(ctx *FuncExpressionContext) {}
