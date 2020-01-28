package parser

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/AleckDarcy/210A/zz/ast"
	"github.com/AleckDarcy/210A/zz/grammar"
)

type ParseTreeListener struct {
	antlr.BaseParseTreeListener

	stack Stack
}

func (p *ParseTreeListener) ExitIdentifier(c *grammar.IdentifierContext) {
	name := c.GetText()
	p.stack.Push(ast.IdentifierHelper.New(name))
}

func (p *ParseTreeListener) ExitDeclarator(c *grammar.DeclaratorContext) {}

func (p *ParseTreeListener) ExitDeclaratorList(c *grammar.DeclaratorListContext) {}

func (p *ParseTreeListener) ExitSimpleTypeSpecifier(c *grammar.SimpleTypeSpecifierContext) {
	p.stack.Push(ast.SimpleTypeSpecifierHelper.New(c.GetText()))
}

func (p *ParseTreeListener) ExitListElementTypeSpecifier(c *grammar.ListElementTypeSpecifierContext) {
	item, _ := p.stack.PopByType(ast.NoderListElementTypeSpecifierer) // todo: error
	switch s := item.(type) {
	case *ast.SimpleTypeSpecifier:
		p.stack.Push(ast.ListElementTypeSpecifierHelper.New(s, ast.ListElementTypeSpecifierSimple))
	case *ast.ListElementTypeSpecifier:
		p.stack.Push(ast.ListElementTypeSpecifierHelper.New(s, ast.ListElementTypeSpecifierNested))
	}
}

func (p *ParseTreeListener) ExitListTypeSpecifier(c *grammar.ListTypeSpecifierContext) {
	item, _ := p.stack.PopByType(ast.NoderListElementTypeSpecifier) // todo: error
	p.stack.Push(ast.ListTypeSpecifierHelper.New(item.(*ast.ListElementTypeSpecifier)))
}

func (p *ParseTreeListener) ExitListInitExpression(c *grammar.ListInitExpressionContext) {
	size, _ := p.stack.PopByType(ast.NoderAExpr)                      // todo: error
	typeSpecifier, _ := p.stack.PopByType(ast.NoderListTypeSpecifier) // todo: error
	p.stack.Push(ast.ListInitExprHelper.New(typeSpecifier.(*ast.ListTypeSpecifier), size.(ast.AExpr)))
}

func (p *ParseTreeListener) ExitAExp_bracketExpression(c *grammar.AExp_bracketExpressionContext) {}

func (p *ParseTreeListener) ExitAExp_FloatLiteral(c *grammar.AExp_FloatLiteralContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitAExp_multiplicativeExpression(c *grammar.AExp_multiplicativeExpressionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitAExp_additiveExpression(c *grammar.AExp_additiveExpressionContext) {
	e2, _ := p.stack.PopByType(ast.NoderAExpr)
	e1, _ := p.stack.PopByType(ast.NoderAExpr)

	switch c.GetChild(1).(antlr.TerminalNode).GetText() {
	case "+":
		p.stack.Push(ast.AExprArithHelper.New(e1.(ast.AExpr), e2.(ast.AExpr), ast.AExprArithAdd))
	case "-":
		fmt.Println("todo ExitAExp_additiveExpression: -")
	}
}

func (p *ParseTreeListener) ExitAExp_identifier(c *grammar.AExp_identifierContext) {
	id, _ := p.stack.PopByType(ast.NoderIdentifier) // todo: error
	p.stack.Push(ast.AExprSimpleHelper.New(id.(*ast.Identifier)))
}

func (p *ParseTreeListener) ExitAExp_IntergerLiteral(c *grammar.AExp_IntergerLiteralContext) {
	value, _ := strconv.ParseInt(c.GetText(), 10, 64) // todo: err
	p.stack.Push(ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(value)))
}

func (p *ParseTreeListener) ExitAExp_listElementExpression(c *grammar.AExp_listElementExpressionContext) {
}

func (p *ParseTreeListener) ExitBExpr_aExpr(c *grammar.BExpr_aExprContext) {
	e2, _ := p.stack.PopByType(ast.NoderAExpr)
	e1, _ := p.stack.PopByType(ast.NoderAExpr)

	switch c.GetChild(1).(antlr.TerminalNode).GetText() {
	case "==":
		p.stack.Push(ast.BExprCompareHelper.New(e1.(ast.AExpr), e2.(ast.AExpr), ast.BExprCompareEQ))
	case "<":
		p.stack.Push(ast.BExprCompareHelper.New(e1.(ast.AExpr), e2.(ast.AExpr), ast.BExprCompareLT))
	default:
		fmt.Println("todo ExitBExpr_aExpr: < > <= >= !=")
	}
}

func (p *ParseTreeListener) ExitBExpr_bExpr(c *grammar.BExpr_bExprContext) {
	e2, _ := p.stack.PopByType(ast.NoderBExpr)
	e1, _ := p.stack.PopByType(ast.NoderBExpr)

	switch c.GetChild(1).(antlr.TerminalNode).GetText() {
	case "==":
		p.stack.Push(ast.BExprBinaryHelper.New(e1.(ast.BExpr), e2.(ast.BExpr), ast.BExprBinaryEQ))
	default:
		fmt.Println("todo ExitBExpr_bExpr: && || !=")
	}
}

func (p *ParseTreeListener) ExitBExpr_bang(c *grammar.BExpr_bangContext) {
	e1, _ := p.stack.PopByType(ast.NoderBExpr)
	p.stack.Push(ast.BExprBinaryHelper.New(e1.(ast.BExpr), ast.BinaryLiteralHelper.New(false), ast.BExprBinaryEQ))
}

func (p *ParseTreeListener) ExitBExpr_bracketExpression(c *grammar.BExpr_bracketExpressionContext) {}

func (p *ParseTreeListener) ExitListElementIndex(c *grammar.ListElementIndexContext) {
	e, _ := p.stack.PopByType(ast.NoderAExpr) // todo: error
	p.stack.Push(ast.ListElementIndexHelper.New(e.(ast.AExpr)))
}

func (p *ParseTreeListener) ExitListElementIndexList(c *grammar.ListElementIndexListContext) {}

func (p *ParseTreeListener) ExitListElementExpression(c *grammar.ListElementExpressionContext) {
	list := p.readListElementIndexList()
	name, _ := p.stack.PopByType(ast.NoderIdentifier)
	p.stack.Push(ast.ListElementExprHelper.New(name.(*ast.Identifier), list))
}

func (p *ParseTreeListener) ExitAssignInit(c *grammar.AssignInitContext) {}

func (p *ParseTreeListener) ExitAssignInitList(c *grammar.AssignInitListContext) {}

func (p *ParseTreeListener) ExitAssignStatement(c *grammar.AssignStatementContext) {
	initList := p.readAssignIniterList()
	declList := p.readDeclaratorerList()
	p.stack.Push(ast.AssignStmtHelper.New(declList, initList))
}

func (p *ParseTreeListener) ExitIfExpr(c *grammar.IfExprContext) {
	stmtList := p.readFuncStatementerList()
	item, _ := p.stack.PopByType(ast.NoderBExpr) // todo: error
	p.stack.Push(ast.IfExprHelper.New(item.(ast.BExpr), stmtList))
}

func (p *ParseTreeListener) ExitElsifExpr(c *grammar.ElsifExprContext) {
	stmtList := p.readFuncStatementerList()
	binExpr, _ := p.stack.PopByType(ast.NoderBExpr) // todo: error
	p.stack.Push(ast.IfExprHelper.New(binExpr.(ast.BExpr), stmtList))
}

func (p *ParseTreeListener) ExitElseExpr(c *grammar.ElseExprContext) {
	p.stack.Push(ast.ElseExprHelper.New(p.readFuncStatementerList()))
}

func (p *ParseTreeListener) ExitTernaryIfExpr(c *grammar.TernaryIfExprContext) {
	stmt, _ := p.stack.PopByType(ast.NoderFuncStatementer) // todo: error
	binExpr, _ := p.stack.PopByType(ast.NoderBExpr)        // todo: error
	p.stack.Push(ast.IfExprHelper.New(binExpr.(ast.BExpr), []ast.FuncStatementer{stmt.(ast.FuncStatementer)}))
}

func (p *ParseTreeListener) ExitTernaryElseExpr(c *grammar.TernaryElseExprContext) {
	stmt, _ := p.stack.PopByType(ast.NoderFuncStatementer) // todo: error
	p.stack.Push(ast.ElseExprHelper.New([]ast.FuncStatementer{stmt.(ast.FuncStatementer)}))
}

func (p *ParseTreeListener) ExitSelectionStatement(c *grammar.SelectionStatementContext) {
	elseExpr, _ := p.stack.PopByType(ast.NoderElseExpr)
	ifExprList := p.readIfExprList()
	if elseExpr != nil {
		p.stack.Push(ast.SelectionStmtHelper.New(ifExprList, elseExpr.(*ast.ElseExpr)))
	} else {
		p.stack.Push(ast.SelectionStmtHelper.New(ifExprList, nil))
	}
}

func (p *ParseTreeListener) ExitIterationAssignStatement(c *grammar.IterationAssignStatementContext) {
	stmt, _ := p.stack.PopByType(ast.NoderAssignStatment)
	if stmt != nil {
		p.stack.Push(ast.IterationAssignStmtHelper.New(stmt.(*ast.AssignStmt)))
	} else {
		p.stack.Push(ast.IterationAssignStmtHelper.New(nil))
	}
}

func (p *ParseTreeListener) ExitIterationStatement(c *grammar.IterationStatementContext) {
	stmtList := p.readFuncStatementerList()
	increStmt, _ := p.stack.PopByType(ast.NoderIterationAssignStatement)
	binExprItf, _ := p.stack.PopByType(ast.NoderBExpr)
	initStmt, _ := p.stack.PopByType(ast.NoderIterationAssignStatement)

	if binExprItf != nil {
		p.stack.Push(
			ast.IterationStmtHelper.New(
				initStmt.(*ast.IterationAssignStmt).AssignStmt(),
				binExprItf.(ast.BExpr),
				increStmt.(*ast.IterationAssignStmt).AssignStmt(),
				stmtList,
			),
		)
	} else {
		p.stack.Push(
			ast.IterationStmtHelper.New(
				initStmt.(*ast.IterationAssignStmt).AssignStmt(),
				ast.BinaryLiteralHelper.New(true),
				increStmt.(*ast.IterationAssignStmt).AssignStmt(),
				stmtList,
			),
		)
	}
}

func (p *ParseTreeListener) ExitDefinition(c *grammar.DefinitionContext) {
}

func (p *ParseTreeListener) ExitDefinitionList(c *grammar.DefinitionListContext) {
}

func (p *ParseTreeListener) ExitFile(c *grammar.FileContext) {
	defList := p.readDefinitionerList()
	p.stack.Push(ast.FileHelper.New(defList))
}

func (p *ParseTreeListener) ExitTypeSpecifier(c *grammar.TypeSpecifierContext) {
}

func (p *ParseTreeListener) ExitTypeSpecifierList(c *grammar.TypeSpecifierListContext) {}

func (p *ParseTreeListener) ExitParaDeclarator(c *grammar.ParaDeclaratorContext) {
	id, _ := p.stack.PopByType(ast.NoderIdentifier)
	p.stack.Push(ast.ParaDeclaratorHelper.New(id.(*ast.Identifier)))
}

func (p *ParseTreeListener) ExitParaDeclaratorList(c *grammar.ParaDeclaratorListContext) {}

func (p *ParseTreeListener) ExitParaDeclaratorWithIdentity(c *grammar.ParaDeclaratorWithIdentityContext) {
	typeSpecifier, _ := p.stack.PopByType(ast.NoderTypeSpecifierer)
	declList := p.readParaDeclaratorListToIdentifierList()
	p.stack.Push(ast.ParaDeclaratorWithIdentityHelper.New(declList, typeSpecifier.(ast.TypeSpecifierer)))
}

func (p *ParseTreeListener) ExitParaDeclaratorWithIdentityList(c *grammar.ParaDeclaratorWithIdentityListContext) {
}

func (p *ParseTreeListener) ExitFuncTypeSpecifier(c *grammar.FuncTypeSpecifierContext) {
	returnList := p.readTypeSpecifierList()
	paraList := p.readParaDeclaratorWithIdentityList()

	p.stack.Push(ast.FuncTypeSpecifierHelper.New(paraList, returnList))
}

func (p *ParseTreeListener) ExitFuncIdentifier(c *grammar.FuncIdentifierContext) {
	name, _ := p.stack.PopByType(ast.NoderIdentifier) // todo: error
	p.stack.Push(ast.FuncIdentityHelper.New(name.(*ast.Identifier)))
}

func (p *ParseTreeListener) ExitFuncTypeSpecifierWithName(c *grammar.FuncTypeSpecifierWithNameContext) {
	returnList := p.readTypeSpecifierList()
	paraList := p.readParaDeclaratorWithIdentityList()
	name, _ := p.stack.PopByType(ast.NoderFuncIdentifier) // todo: error
	p.stack.Push(ast.FuncTypeSpecifierWithNameHelper.New(name.(*ast.FuncIdentifier).Name(), paraList, returnList))
}

func (p *ParseTreeListener) ExitFuncReturnPara(c *grammar.FuncReturnParaContext) {}

func (p *ParseTreeListener) ExitFuncReturnParaList(c *grammar.FuncReturnParaListContext) {}

func (p *ParseTreeListener) ExitFuncReturnStatement(c *grammar.FuncReturnStatementContext) {
	returnList := p.readFuncReturnParaList()
	p.stack.Push(ast.FuncReturnStatementHelper.New(returnList))
}

func (p *ParseTreeListener) ExitFuncStatement(c *grammar.FuncStatementContext) {}

func (p *ParseTreeListener) ExitFuncStatementList(c *grammar.FuncStatementListContext) {}

func (p *ParseTreeListener) ExitFuncBody(c *grammar.FuncBodyContext) {}

func (p *ParseTreeListener) ExitFuncInitExpression(c *grammar.FuncInitExpressionContext) {
	stmtList := p.readFuncStatementerList()
	typeSpecifier, _ := p.stack.PopByType(ast.NoderFuncTypeSpecifier)
	p.stack.Push(ast.FuncInitExprHelper.New(typeSpecifier.(*ast.FuncTypeSpecifier), stmtList))
}

func (p *ParseTreeListener) ExitFuncDefinition(c *grammar.FuncDefinitionContext) {
	stmt := p.readFuncStatementerList()
	typeSpecifier, _ := p.stack.PopByType(ast.NoderFuncTypeSpecifierWithName)
	p.stack.Push(ast.FuncDefinitionHelper.New(typeSpecifier.(*ast.FuncTypeSpecifierWithName), stmt))
}

func (p *ParseTreeListener) ExitFuncExecutePara(c *grammar.FuncExecuteParaContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitFuncExecuteParaList(c *grammar.FuncExecuteParaListContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitFuncExecuteExpression(c *grammar.FuncExecuteExpressionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitFuncExecuteStatement(c *grammar.FuncExecuteStatementContext) {
	panic("implement me")
}

/*
	Enter methods
	do nothing
*/
func (p *ParseTreeListener) EnterIdentifier(c *grammar.IdentifierContext) {}

func (p *ParseTreeListener) EnterDeclarator(c *grammar.DeclaratorContext) {}

func (p *ParseTreeListener) EnterDeclaratorList(c *grammar.DeclaratorListContext) {}

func (p *ParseTreeListener) EnterSimpleTypeSpecifier(c *grammar.SimpleTypeSpecifierContext) {}

func (p *ParseTreeListener) EnterListElementTypeSpecifier(c *grammar.ListElementTypeSpecifierContext) {
}

func (p *ParseTreeListener) EnterListTypeSpecifier(c *grammar.ListTypeSpecifierContext) {}

func (p *ParseTreeListener) EnterListInitExpression(c *grammar.ListInitExpressionContext) {}

func (p *ParseTreeListener) EnterAExp_bracketExpression(c *grammar.AExp_bracketExpressionContext) {}

func (p *ParseTreeListener) EnterAExp_FloatLiteral(c *grammar.AExp_FloatLiteralContext) {}

func (p *ParseTreeListener) EnterAExp_multiplicativeExpression(c *grammar.AExp_multiplicativeExpressionContext) {
}

func (p *ParseTreeListener) EnterAExp_additiveExpression(c *grammar.AExp_additiveExpressionContext) {}

func (p *ParseTreeListener) EnterAExp_identifier(c *grammar.AExp_identifierContext) {}

func (p *ParseTreeListener) EnterAExp_IntergerLiteral(c *grammar.AExp_IntergerLiteralContext) {}

func (p *ParseTreeListener) EnterAExp_listElementExpression(c *grammar.AExp_listElementExpressionContext) {
}

func (p *ParseTreeListener) EnterBExpr_aExpr(c *grammar.BExpr_aExprContext) {}

func (p *ParseTreeListener) EnterBExpr_bExpr(c *grammar.BExpr_bExprContext) {}

func (p *ParseTreeListener) EnterBExpr_bang(c *grammar.BExpr_bangContext) {}

func (p *ParseTreeListener) EnterBExpr_bracketExpression(c *grammar.BExpr_bracketExpressionContext) {}

func (p *ParseTreeListener) EnterListElementIndex(c *grammar.ListElementIndexContext) {}

func (p *ParseTreeListener) EnterListElementIndexList(c *grammar.ListElementIndexListContext) {}

func (p *ParseTreeListener) EnterListElementExpression(c *grammar.ListElementExpressionContext) {}

func (p *ParseTreeListener) EnterAssignInit(c *grammar.AssignInitContext) {}

func (p *ParseTreeListener) EnterAssignInitList(c *grammar.AssignInitListContext) {}

func (p *ParseTreeListener) EnterAssignStatement(c *grammar.AssignStatementContext) {}

func (p *ParseTreeListener) EnterIfExpr(c *grammar.IfExprContext) {}

func (p *ParseTreeListener) EnterElsifExpr(c *grammar.ElsifExprContext) {}

func (p *ParseTreeListener) EnterElseExpr(c *grammar.ElseExprContext) {}

func (p *ParseTreeListener) EnterTernaryIfExpr(c *grammar.TernaryIfExprContext) {}

func (p *ParseTreeListener) EnterTernaryElseExpr(c *grammar.TernaryElseExprContext) {}

func (p *ParseTreeListener) EnterSelectionStatement(c *grammar.SelectionStatementContext) {}

func (p *ParseTreeListener) EnterIterationAssignStatement(c *grammar.IterationAssignStatementContext) {
}

func (p *ParseTreeListener) EnterIterationStatement(c *grammar.IterationStatementContext) {}

func (p *ParseTreeListener) EnterDefinition(c *grammar.DefinitionContext) {}

func (p *ParseTreeListener) EnterDefinitionList(c *grammar.DefinitionListContext) {}

func (p *ParseTreeListener) EnterFile(c *grammar.FileContext) {}

func (p *ParseTreeListener) EnterTypeSpecifier(c *grammar.TypeSpecifierContext) {}

func (p *ParseTreeListener) EnterTypeSpecifierList(c *grammar.TypeSpecifierListContext) {}

func (p *ParseTreeListener) EnterParaDeclarator(c *grammar.ParaDeclaratorContext) {}

func (p *ParseTreeListener) EnterParaDeclaratorList(c *grammar.ParaDeclaratorListContext) {}

func (p *ParseTreeListener) EnterParaDeclaratorWithIdentity(c *grammar.ParaDeclaratorWithIdentityContext) {
}

func (p *ParseTreeListener) EnterParaDeclaratorWithIdentityList(c *grammar.ParaDeclaratorWithIdentityListContext) {
}

func (p *ParseTreeListener) EnterFuncTypeSpecifier(c *grammar.FuncTypeSpecifierContext) {}

func (p *ParseTreeListener) EnterFuncIdentifier(c *grammar.FuncIdentifierContext) {}

func (p *ParseTreeListener) EnterFuncTypeSpecifierWithName(c *grammar.FuncTypeSpecifierWithNameContext) {
}

func (p *ParseTreeListener) EnterFuncReturnPara(c *grammar.FuncReturnParaContext) {}

func (p *ParseTreeListener) EnterFuncReturnParaList(c *grammar.FuncReturnParaListContext) {}

func (p *ParseTreeListener) EnterFuncReturnStatement(c *grammar.FuncReturnStatementContext) {}

func (p *ParseTreeListener) EnterFuncStatement(c *grammar.FuncStatementContext) {}

func (p *ParseTreeListener) EnterFuncStatementList(c *grammar.FuncStatementListContext) {}

func (p *ParseTreeListener) EnterFuncBody(c *grammar.FuncBodyContext) {}

func (p *ParseTreeListener) EnterFuncInitExpression(c *grammar.FuncInitExpressionContext) {}

func (p *ParseTreeListener) EnterFuncDefinition(c *grammar.FuncDefinitionContext) {}

func (p *ParseTreeListener) EnterFuncExecutePara(c *grammar.FuncExecuteParaContext) {}

func (p *ParseTreeListener) EnterFuncExecuteParaList(c *grammar.FuncExecuteParaListContext) {}

func (p *ParseTreeListener) EnterFuncExecuteExpression(c *grammar.FuncExecuteExpressionContext) {}

func (p *ParseTreeListener) EnterFuncExecuteStatement(c *grammar.FuncExecuteStatementContext) {}
