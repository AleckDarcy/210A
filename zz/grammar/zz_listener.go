// Code generated from zz/grammar/ZZ.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar // ZZ
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ZZListener is a complete listener for a parse tree produced by ZZParser.
type ZZListener interface {
	antlr.ParseTreeListener

	// EnterDeclarator is called when entering the declarator production.
	EnterDeclarator(c *DeclaratorContext)

	// EnterDeclaratorList is called when entering the declaratorList production.
	EnterDeclaratorList(c *DeclaratorListContext)

	// EnterSimpleTypeSpecifier is called when entering the simpleTypeSpecifier production.
	EnterSimpleTypeSpecifier(c *SimpleTypeSpecifierContext)

	// EnterListElementTypeSpecifier is called when entering the listElementTypeSpecifier production.
	EnterListElementTypeSpecifier(c *ListElementTypeSpecifierContext)

	// EnterListTypeSpecifier is called when entering the listTypeSpecifier production.
	EnterListTypeSpecifier(c *ListTypeSpecifierContext)

	// EnterListInitExpression is called when entering the listInitExpression production.
	EnterListInitExpression(c *ListInitExpressionContext)

	// EnterAExp_bracketExpression is called when entering the aExp_bracketExpression production.
	EnterAExp_bracketExpression(c *AExp_bracketExpressionContext)

	// EnterAExp_FloatLiteral is called when entering the aExp_FloatLiteral production.
	EnterAExp_FloatLiteral(c *AExp_FloatLiteralContext)

	// EnterAExp_multiplicativeExpression is called when entering the aExp_multiplicativeExpression production.
	EnterAExp_multiplicativeExpression(c *AExp_multiplicativeExpressionContext)

	// EnterAExp_additiveExpression is called when entering the aExp_additiveExpression production.
	EnterAExp_additiveExpression(c *AExp_additiveExpressionContext)

	// EnterAExp_Identifier is called when entering the aExp_Identifier production.
	EnterAExp_Identifier(c *AExp_IdentifierContext)

	// EnterAExp_IntergerLiteral is called when entering the aExp_IntergerLiteral production.
	EnterAExp_IntergerLiteral(c *AExp_IntergerLiteralContext)

	// EnterAExp_listElementExpression is called when entering the aExp_listElementExpression production.
	EnterAExp_listElementExpression(c *AExp_listElementExpressionContext)

	// EnterAExprList is called when entering the aExprList production.
	EnterAExprList(c *AExprListContext)

	// EnterBExp is called when entering the bExp production.
	EnterBExp(c *BExpContext)

	// EnterIntegerExpression is called when entering the integerExpression production.
	EnterIntegerExpression(c *IntegerExpressionContext)

	// EnterListElementIndex is called when entering the listElementIndex production.
	EnterListElementIndex(c *ListElementIndexContext)

	// EnterListElementIndexList is called when entering the listElementIndexList production.
	EnterListElementIndexList(c *ListElementIndexListContext)

	// EnterListElementExpression is called when entering the listElementExpression production.
	EnterListElementExpression(c *ListElementExpressionContext)

	// EnterAssignInit is called when entering the assignInit production.
	EnterAssignInit(c *AssignInitContext)

	// EnterAssignInitList is called when entering the assignInitList production.
	EnterAssignInitList(c *AssignInitListContext)

	// EnterAssignStatement is called when entering the assignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterSelectionStatement is called when entering the selectionStatement production.
	EnterSelectionStatement(c *SelectionStatementContext)

	// EnterIterationStatement is called when entering the iterationStatement production.
	EnterIterationStatement(c *IterationStatementContext)

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterEntryList is called when entering the entryList production.
	EnterEntryList(c *EntryListContext)

	// EnterTypeSpecifier is called when entering the typeSpecifier production.
	EnterTypeSpecifier(c *TypeSpecifierContext)

	// EnterTypeSpecifierList is called when entering the typeSpecifierList production.
	EnterTypeSpecifierList(c *TypeSpecifierListContext)

	// EnterParaDeclaratorList is called when entering the paraDeclaratorList production.
	EnterParaDeclaratorList(c *ParaDeclaratorListContext)

	// EnterParaDeclaratorWithIdentity is called when entering the paraDeclaratorWithIdentity production.
	EnterParaDeclaratorWithIdentity(c *ParaDeclaratorWithIdentityContext)

	// EnterParaDeclaratorWithIdentityList is called when entering the paraDeclaratorWithIdentityList production.
	EnterParaDeclaratorWithIdentityList(c *ParaDeclaratorWithIdentityListContext)

	// EnterFuncTypeSpecifier is called when entering the funcTypeSpecifier production.
	EnterFuncTypeSpecifier(c *FuncTypeSpecifierContext)

	// EnterFuncTypeSpecifierWithName is called when entering the funcTypeSpecifierWithName production.
	EnterFuncTypeSpecifierWithName(c *FuncTypeSpecifierWithNameContext)

	// EnterFuncReturnPara is called when entering the funcReturnPara production.
	EnterFuncReturnPara(c *FuncReturnParaContext)

	// EnterFuncReturnParaList is called when entering the funcReturnParaList production.
	EnterFuncReturnParaList(c *FuncReturnParaListContext)

	// EnterFuncStatement is called when entering the funcStatement production.
	EnterFuncStatement(c *FuncStatementContext)

	// EnterFuncStatementList is called when entering the funcStatementList production.
	EnterFuncStatementList(c *FuncStatementListContext)

	// EnterFuncBody is called when entering the funcBody production.
	EnterFuncBody(c *FuncBodyContext)

	// EnterFuncInitExpression is called when entering the funcInitExpression production.
	EnterFuncInitExpression(c *FuncInitExpressionContext)

	// EnterFuncExpression is called when entering the funcExpression production.
	EnterFuncExpression(c *FuncExpressionContext)

	// ExitDeclarator is called when exiting the declarator production.
	ExitDeclarator(c *DeclaratorContext)

	// ExitDeclaratorList is called when exiting the declaratorList production.
	ExitDeclaratorList(c *DeclaratorListContext)

	// ExitSimpleTypeSpecifier is called when exiting the simpleTypeSpecifier production.
	ExitSimpleTypeSpecifier(c *SimpleTypeSpecifierContext)

	// ExitListElementTypeSpecifier is called when exiting the listElementTypeSpecifier production.
	ExitListElementTypeSpecifier(c *ListElementTypeSpecifierContext)

	// ExitListTypeSpecifier is called when exiting the listTypeSpecifier production.
	ExitListTypeSpecifier(c *ListTypeSpecifierContext)

	// ExitListInitExpression is called when exiting the listInitExpression production.
	ExitListInitExpression(c *ListInitExpressionContext)

	// ExitAExp_bracketExpression is called when exiting the aExp_bracketExpression production.
	ExitAExp_bracketExpression(c *AExp_bracketExpressionContext)

	// ExitAExp_FloatLiteral is called when exiting the aExp_FloatLiteral production.
	ExitAExp_FloatLiteral(c *AExp_FloatLiteralContext)

	// ExitAExp_multiplicativeExpression is called when exiting the aExp_multiplicativeExpression production.
	ExitAExp_multiplicativeExpression(c *AExp_multiplicativeExpressionContext)

	// ExitAExp_additiveExpression is called when exiting the aExp_additiveExpression production.
	ExitAExp_additiveExpression(c *AExp_additiveExpressionContext)

	// ExitAExp_Identifier is called when exiting the aExp_Identifier production.
	ExitAExp_Identifier(c *AExp_IdentifierContext)

	// ExitAExp_IntergerLiteral is called when exiting the aExp_IntergerLiteral production.
	ExitAExp_IntergerLiteral(c *AExp_IntergerLiteralContext)

	// ExitAExp_listElementExpression is called when exiting the aExp_listElementExpression production.
	ExitAExp_listElementExpression(c *AExp_listElementExpressionContext)

	// ExitAExprList is called when exiting the aExprList production.
	ExitAExprList(c *AExprListContext)

	// ExitBExp is called when exiting the bExp production.
	ExitBExp(c *BExpContext)

	// ExitIntegerExpression is called when exiting the integerExpression production.
	ExitIntegerExpression(c *IntegerExpressionContext)

	// ExitListElementIndex is called when exiting the listElementIndex production.
	ExitListElementIndex(c *ListElementIndexContext)

	// ExitListElementIndexList is called when exiting the listElementIndexList production.
	ExitListElementIndexList(c *ListElementIndexListContext)

	// ExitListElementExpression is called when exiting the listElementExpression production.
	ExitListElementExpression(c *ListElementExpressionContext)

	// ExitAssignInit is called when exiting the assignInit production.
	ExitAssignInit(c *AssignInitContext)

	// ExitAssignInitList is called when exiting the assignInitList production.
	ExitAssignInitList(c *AssignInitListContext)

	// ExitAssignStatement is called when exiting the assignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitSelectionStatement is called when exiting the selectionStatement production.
	ExitSelectionStatement(c *SelectionStatementContext)

	// ExitIterationStatement is called when exiting the iterationStatement production.
	ExitIterationStatement(c *IterationStatementContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitEntryList is called when exiting the entryList production.
	ExitEntryList(c *EntryListContext)

	// ExitTypeSpecifier is called when exiting the typeSpecifier production.
	ExitTypeSpecifier(c *TypeSpecifierContext)

	// ExitTypeSpecifierList is called when exiting the typeSpecifierList production.
	ExitTypeSpecifierList(c *TypeSpecifierListContext)

	// ExitParaDeclaratorList is called when exiting the paraDeclaratorList production.
	ExitParaDeclaratorList(c *ParaDeclaratorListContext)

	// ExitParaDeclaratorWithIdentity is called when exiting the paraDeclaratorWithIdentity production.
	ExitParaDeclaratorWithIdentity(c *ParaDeclaratorWithIdentityContext)

	// ExitParaDeclaratorWithIdentityList is called when exiting the paraDeclaratorWithIdentityList production.
	ExitParaDeclaratorWithIdentityList(c *ParaDeclaratorWithIdentityListContext)

	// ExitFuncTypeSpecifier is called when exiting the funcTypeSpecifier production.
	ExitFuncTypeSpecifier(c *FuncTypeSpecifierContext)

	// ExitFuncTypeSpecifierWithName is called when exiting the funcTypeSpecifierWithName production.
	ExitFuncTypeSpecifierWithName(c *FuncTypeSpecifierWithNameContext)

	// ExitFuncReturnPara is called when exiting the funcReturnPara production.
	ExitFuncReturnPara(c *FuncReturnParaContext)

	// ExitFuncReturnParaList is called when exiting the funcReturnParaList production.
	ExitFuncReturnParaList(c *FuncReturnParaListContext)

	// ExitFuncStatement is called when exiting the funcStatement production.
	ExitFuncStatement(c *FuncStatementContext)

	// ExitFuncStatementList is called when exiting the funcStatementList production.
	ExitFuncStatementList(c *FuncStatementListContext)

	// ExitFuncBody is called when exiting the funcBody production.
	ExitFuncBody(c *FuncBodyContext)

	// ExitFuncInitExpression is called when exiting the funcInitExpression production.
	ExitFuncInitExpression(c *FuncInitExpressionContext)

	// ExitFuncExpression is called when exiting the funcExpression production.
	ExitFuncExpression(c *FuncExpressionContext)
}
