// Code generated from zz/grammar/ZZ.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar // ZZ
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ZZListener is a complete listener for a parse tree produced by ZZParser.
type ZZListener interface {
	antlr.ParseTreeListener

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

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

	// EnterMatrixInitExpression is called when entering the matrixInitExpression production.
	EnterMatrixInitExpression(c *MatrixInitExpressionContext)

	// EnterAExp_bracketExpression is called when entering the aExp_bracketExpression production.
	EnterAExp_bracketExpression(c *AExp_bracketExpressionContext)

	// EnterAExp_FloatLiteral is called when entering the aExp_FloatLiteral production.
	EnterAExp_FloatLiteral(c *AExp_FloatLiteralContext)

	// EnterAExp_multiplicativeExpression is called when entering the aExp_multiplicativeExpression production.
	EnterAExp_multiplicativeExpression(c *AExp_multiplicativeExpressionContext)

	// EnterAExp_additiveExpression is called when entering the aExp_additiveExpression production.
	EnterAExp_additiveExpression(c *AExp_additiveExpressionContext)

	// EnterAExp_transpose is called when entering the aExp_transpose production.
	EnterAExp_transpose(c *AExp_transposeContext)

	// EnterAExp_identifier is called when entering the aExp_identifier production.
	EnterAExp_identifier(c *AExp_identifierContext)

	// EnterAExp_IntergerLiteral is called when entering the aExp_IntergerLiteral production.
	EnterAExp_IntergerLiteral(c *AExp_IntergerLiteralContext)

	// EnterAExp_listElementExpression is called when entering the aExp_listElementExpression production.
	EnterAExp_listElementExpression(c *AExp_listElementExpressionContext)

	// EnterAExprList is called when entering the aExprList production.
	EnterAExprList(c *AExprListContext)

	// EnterBExpr_aExpr is called when entering the bExpr_aExpr production.
	EnterBExpr_aExpr(c *BExpr_aExprContext)

	// EnterBExpr_bang is called when entering the bExpr_bang production.
	EnterBExpr_bang(c *BExpr_bangContext)

	// EnterBExpr_bExpr is called when entering the bExpr_bExpr production.
	EnterBExpr_bExpr(c *BExpr_bExprContext)

	// EnterBExpr_bracketExpression is called when entering the bExpr_bracketExpression production.
	EnterBExpr_bracketExpression(c *BExpr_bracketExpressionContext)

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

	// EnterAssignInitStatement is called when entering the assignInitStatement production.
	EnterAssignInitStatement(c *AssignInitStatementContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterElsifExpr is called when entering the elsifExpr production.
	EnterElsifExpr(c *ElsifExprContext)

	// EnterElseExpr is called when entering the elseExpr production.
	EnterElseExpr(c *ElseExprContext)

	// EnterTernaryIfExpr is called when entering the ternaryIfExpr production.
	EnterTernaryIfExpr(c *TernaryIfExprContext)

	// EnterTernaryElseExpr is called when entering the ternaryElseExpr production.
	EnterTernaryElseExpr(c *TernaryElseExprContext)

	// EnterSelectionStatement is called when entering the selectionStatement production.
	EnterSelectionStatement(c *SelectionStatementContext)

	// EnterIterationAssignInitStatement is called when entering the iterationAssignInitStatement production.
	EnterIterationAssignInitStatement(c *IterationAssignInitStatementContext)

	// EnterIterationAssignStatement is called when entering the iterationAssignStatement production.
	EnterIterationAssignStatement(c *IterationAssignStatementContext)

	// EnterIterationStatement is called when entering the iterationStatement production.
	EnterIterationStatement(c *IterationStatementContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterDefinitionList is called when entering the definitionList production.
	EnterDefinitionList(c *DefinitionListContext)

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterTypeSpecifier is called when entering the typeSpecifier production.
	EnterTypeSpecifier(c *TypeSpecifierContext)

	// EnterTypeSpecifierList is called when entering the typeSpecifierList production.
	EnterTypeSpecifierList(c *TypeSpecifierListContext)

	// EnterParaDeclarator is called when entering the paraDeclarator production.
	EnterParaDeclarator(c *ParaDeclaratorContext)

	// EnterParaDeclaratorList is called when entering the paraDeclaratorList production.
	EnterParaDeclaratorList(c *ParaDeclaratorListContext)

	// EnterParaDeclaratorWithIdentity is called when entering the paraDeclaratorWithIdentity production.
	EnterParaDeclaratorWithIdentity(c *ParaDeclaratorWithIdentityContext)

	// EnterParaDeclaratorWithIdentityList is called when entering the paraDeclaratorWithIdentityList production.
	EnterParaDeclaratorWithIdentityList(c *ParaDeclaratorWithIdentityListContext)

	// EnterFuncIdentifier is called when entering the funcIdentifier production.
	EnterFuncIdentifier(c *FuncIdentifierContext)

	// EnterFuncTypeSpecifierWithName is called when entering the funcTypeSpecifierWithName production.
	EnterFuncTypeSpecifierWithName(c *FuncTypeSpecifierWithNameContext)

	// EnterFuncReturnPara is called when entering the funcReturnPara production.
	EnterFuncReturnPara(c *FuncReturnParaContext)

	// EnterFuncReturnParaList is called when entering the funcReturnParaList production.
	EnterFuncReturnParaList(c *FuncReturnParaListContext)

	// EnterFuncReturnStatement is called when entering the funcReturnStatement production.
	EnterFuncReturnStatement(c *FuncReturnStatementContext)

	// EnterFuncStatement is called when entering the funcStatement production.
	EnterFuncStatement(c *FuncStatementContext)

	// EnterFuncStatementList is called when entering the funcStatementList production.
	EnterFuncStatementList(c *FuncStatementListContext)

	// EnterFuncBody is called when entering the funcBody production.
	EnterFuncBody(c *FuncBodyContext)

	// EnterFuncDefinition is called when entering the funcDefinition production.
	EnterFuncDefinition(c *FuncDefinitionContext)

	// EnterFuncExecutePara is called when entering the funcExecutePara production.
	EnterFuncExecutePara(c *FuncExecuteParaContext)

	// EnterFuncExecuteParaList is called when entering the funcExecuteParaList production.
	EnterFuncExecuteParaList(c *FuncExecuteParaListContext)

	// EnterFuncExecuteExpression is called when entering the funcExecuteExpression production.
	EnterFuncExecuteExpression(c *FuncExecuteExpressionContext)

	// EnterFuncExecuteStatement is called when entering the funcExecuteStatement production.
	EnterFuncExecuteStatement(c *FuncExecuteStatementContext)

	// EnterPrintList is called when entering the printList production.
	EnterPrintList(c *PrintListContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

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

	// ExitMatrixInitExpression is called when exiting the matrixInitExpression production.
	ExitMatrixInitExpression(c *MatrixInitExpressionContext)

	// ExitAExp_bracketExpression is called when exiting the aExp_bracketExpression production.
	ExitAExp_bracketExpression(c *AExp_bracketExpressionContext)

	// ExitAExp_FloatLiteral is called when exiting the aExp_FloatLiteral production.
	ExitAExp_FloatLiteral(c *AExp_FloatLiteralContext)

	// ExitAExp_multiplicativeExpression is called when exiting the aExp_multiplicativeExpression production.
	ExitAExp_multiplicativeExpression(c *AExp_multiplicativeExpressionContext)

	// ExitAExp_additiveExpression is called when exiting the aExp_additiveExpression production.
	ExitAExp_additiveExpression(c *AExp_additiveExpressionContext)

	// ExitAExp_transpose is called when exiting the aExp_transpose production.
	ExitAExp_transpose(c *AExp_transposeContext)

	// ExitAExp_identifier is called when exiting the aExp_identifier production.
	ExitAExp_identifier(c *AExp_identifierContext)

	// ExitAExp_IntergerLiteral is called when exiting the aExp_IntergerLiteral production.
	ExitAExp_IntergerLiteral(c *AExp_IntergerLiteralContext)

	// ExitAExp_listElementExpression is called when exiting the aExp_listElementExpression production.
	ExitAExp_listElementExpression(c *AExp_listElementExpressionContext)

	// ExitAExprList is called when exiting the aExprList production.
	ExitAExprList(c *AExprListContext)

	// ExitBExpr_aExpr is called when exiting the bExpr_aExpr production.
	ExitBExpr_aExpr(c *BExpr_aExprContext)

	// ExitBExpr_bang is called when exiting the bExpr_bang production.
	ExitBExpr_bang(c *BExpr_bangContext)

	// ExitBExpr_bExpr is called when exiting the bExpr_bExpr production.
	ExitBExpr_bExpr(c *BExpr_bExprContext)

	// ExitBExpr_bracketExpression is called when exiting the bExpr_bracketExpression production.
	ExitBExpr_bracketExpression(c *BExpr_bracketExpressionContext)

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

	// ExitAssignInitStatement is called when exiting the assignInitStatement production.
	ExitAssignInitStatement(c *AssignInitStatementContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitElsifExpr is called when exiting the elsifExpr production.
	ExitElsifExpr(c *ElsifExprContext)

	// ExitElseExpr is called when exiting the elseExpr production.
	ExitElseExpr(c *ElseExprContext)

	// ExitTernaryIfExpr is called when exiting the ternaryIfExpr production.
	ExitTernaryIfExpr(c *TernaryIfExprContext)

	// ExitTernaryElseExpr is called when exiting the ternaryElseExpr production.
	ExitTernaryElseExpr(c *TernaryElseExprContext)

	// ExitSelectionStatement is called when exiting the selectionStatement production.
	ExitSelectionStatement(c *SelectionStatementContext)

	// ExitIterationAssignInitStatement is called when exiting the iterationAssignInitStatement production.
	ExitIterationAssignInitStatement(c *IterationAssignInitStatementContext)

	// ExitIterationAssignStatement is called when exiting the iterationAssignStatement production.
	ExitIterationAssignStatement(c *IterationAssignStatementContext)

	// ExitIterationStatement is called when exiting the iterationStatement production.
	ExitIterationStatement(c *IterationStatementContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitDefinitionList is called when exiting the definitionList production.
	ExitDefinitionList(c *DefinitionListContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitTypeSpecifier is called when exiting the typeSpecifier production.
	ExitTypeSpecifier(c *TypeSpecifierContext)

	// ExitTypeSpecifierList is called when exiting the typeSpecifierList production.
	ExitTypeSpecifierList(c *TypeSpecifierListContext)

	// ExitParaDeclarator is called when exiting the paraDeclarator production.
	ExitParaDeclarator(c *ParaDeclaratorContext)

	// ExitParaDeclaratorList is called when exiting the paraDeclaratorList production.
	ExitParaDeclaratorList(c *ParaDeclaratorListContext)

	// ExitParaDeclaratorWithIdentity is called when exiting the paraDeclaratorWithIdentity production.
	ExitParaDeclaratorWithIdentity(c *ParaDeclaratorWithIdentityContext)

	// ExitParaDeclaratorWithIdentityList is called when exiting the paraDeclaratorWithIdentityList production.
	ExitParaDeclaratorWithIdentityList(c *ParaDeclaratorWithIdentityListContext)

	// ExitFuncIdentifier is called when exiting the funcIdentifier production.
	ExitFuncIdentifier(c *FuncIdentifierContext)

	// ExitFuncTypeSpecifierWithName is called when exiting the funcTypeSpecifierWithName production.
	ExitFuncTypeSpecifierWithName(c *FuncTypeSpecifierWithNameContext)

	// ExitFuncReturnPara is called when exiting the funcReturnPara production.
	ExitFuncReturnPara(c *FuncReturnParaContext)

	// ExitFuncReturnParaList is called when exiting the funcReturnParaList production.
	ExitFuncReturnParaList(c *FuncReturnParaListContext)

	// ExitFuncReturnStatement is called when exiting the funcReturnStatement production.
	ExitFuncReturnStatement(c *FuncReturnStatementContext)

	// ExitFuncStatement is called when exiting the funcStatement production.
	ExitFuncStatement(c *FuncStatementContext)

	// ExitFuncStatementList is called when exiting the funcStatementList production.
	ExitFuncStatementList(c *FuncStatementListContext)

	// ExitFuncBody is called when exiting the funcBody production.
	ExitFuncBody(c *FuncBodyContext)

	// ExitFuncDefinition is called when exiting the funcDefinition production.
	ExitFuncDefinition(c *FuncDefinitionContext)

	// ExitFuncExecutePara is called when exiting the funcExecutePara production.
	ExitFuncExecutePara(c *FuncExecuteParaContext)

	// ExitFuncExecuteParaList is called when exiting the funcExecuteParaList production.
	ExitFuncExecuteParaList(c *FuncExecuteParaListContext)

	// ExitFuncExecuteExpression is called when exiting the funcExecuteExpression production.
	ExitFuncExecuteExpression(c *FuncExecuteExpressionContext)

	// ExitFuncExecuteStatement is called when exiting the funcExecuteStatement production.
	ExitFuncExecuteStatement(c *FuncExecuteStatementContext)

	// ExitPrintList is called when exiting the printList production.
	ExitPrintList(c *PrintListContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)
}
