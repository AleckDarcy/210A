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

	// EnterAExp is called when entering the aExp production.
	EnterAExp(c *AExpContext)

	// EnterAExpList is called when entering the aExpList production.
	EnterAExpList(c *AExpListContext)

	// EnterBExp is called when entering the bExp production.
	EnterBExp(c *BExpContext)

	// EnterIntegerExpression is called when entering the integerExpression production.
	EnterIntegerExpression(c *IntegerExpressionContext)

	// EnterListElement is called when entering the listElement production.
	EnterListElement(c *ListElementContext)

	// EnterListElements is called when entering the listElements production.
	EnterListElements(c *ListElementsContext)

	// EnterListElementExpression is called when entering the listElementExpression production.
	EnterListElementExpression(c *ListElementExpressionContext)

	// EnterTupleSizes is called when entering the tupleSizes production.
	EnterTupleSizes(c *TupleSizesContext)

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

	// EnterTypeSpecifierWithIdentity is called when entering the typeSpecifierWithIdentity production.
	EnterTypeSpecifierWithIdentity(c *TypeSpecifierWithIdentityContext)

	// EnterTypeSpecifierWithIdentityList is called when entering the typeSpecifierWithIdentityList production.
	EnterTypeSpecifierWithIdentityList(c *TypeSpecifierWithIdentityListContext)

	// EnterFuncSpecifier is called when entering the funcSpecifier production.
	EnterFuncSpecifier(c *FuncSpecifierContext)

	// EnterFuncSpecifierWithName is called when entering the funcSpecifierWithName production.
	EnterFuncSpecifierWithName(c *FuncSpecifierWithNameContext)

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

	// EnterFuncExpression is called when entering the funcExpression production.
	EnterFuncExpression(c *FuncExpressionContext)

	// EnterFuncExpressionWithName is called when entering the funcExpressionWithName production.
	EnterFuncExpressionWithName(c *FuncExpressionWithNameContext)

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

	// ExitAExp is called when exiting the aExp production.
	ExitAExp(c *AExpContext)

	// ExitAExpList is called when exiting the aExpList production.
	ExitAExpList(c *AExpListContext)

	// ExitBExp is called when exiting the bExp production.
	ExitBExp(c *BExpContext)

	// ExitIntegerExpression is called when exiting the integerExpression production.
	ExitIntegerExpression(c *IntegerExpressionContext)

	// ExitListElement is called when exiting the listElement production.
	ExitListElement(c *ListElementContext)

	// ExitListElements is called when exiting the listElements production.
	ExitListElements(c *ListElementsContext)

	// ExitListElementExpression is called when exiting the listElementExpression production.
	ExitListElementExpression(c *ListElementExpressionContext)

	// ExitTupleSizes is called when exiting the tupleSizes production.
	ExitTupleSizes(c *TupleSizesContext)

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

	// ExitTypeSpecifierWithIdentity is called when exiting the typeSpecifierWithIdentity production.
	ExitTypeSpecifierWithIdentity(c *TypeSpecifierWithIdentityContext)

	// ExitTypeSpecifierWithIdentityList is called when exiting the typeSpecifierWithIdentityList production.
	ExitTypeSpecifierWithIdentityList(c *TypeSpecifierWithIdentityListContext)

	// ExitFuncSpecifier is called when exiting the funcSpecifier production.
	ExitFuncSpecifier(c *FuncSpecifierContext)

	// ExitFuncSpecifierWithName is called when exiting the funcSpecifierWithName production.
	ExitFuncSpecifierWithName(c *FuncSpecifierWithNameContext)

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

	// ExitFuncExpression is called when exiting the funcExpression production.
	ExitFuncExpression(c *FuncExpressionContext)

	// ExitFuncExpressionWithName is called when exiting the funcExpressionWithName production.
	ExitFuncExpressionWithName(c *FuncExpressionWithNameContext)
}
