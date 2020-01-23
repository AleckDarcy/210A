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

func (p *ParseTreeListener) EnterIdentifier(c *grammar.IdentifierContext) {}

func (p *ParseTreeListener) EnterDeclarator_identifier(c *grammar.Declarator_identifierContext) {
	// todo: remove tag
}

func (p *ParseTreeListener) EnterDeclarator_listElementExpression(c *grammar.Declarator_listElementExpressionContext) {
	// todo: remove tag
}

func (p *ParseTreeListener) EnterDeclaratorList(c *grammar.DeclaratorListContext) {}

func (p *ParseTreeListener) EnterSimpleTypeSpecifier(c *grammar.SimpleTypeSpecifierContext) {}

func (p *ParseTreeListener) EnterListElementTypeSpecifier(c *grammar.ListElementTypeSpecifierContext) {
}

func (p *ParseTreeListener) EnterListTypeSpecifier(c *grammar.ListTypeSpecifierContext) {}

func (p *ParseTreeListener) EnterListInitExpression(c *grammar.ListInitExpressionContext) {}

func (p *ParseTreeListener) EnterAExp_bracketExpression(c *grammar.AExp_bracketExpressionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterAExp_FloatLiteral(c *grammar.AExp_FloatLiteralContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterAExp_multiplicativeExpression(c *grammar.AExp_multiplicativeExpressionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterAExp_additiveExpression(c *grammar.AExp_additiveExpressionContext) {}

func (p *ParseTreeListener) EnterAExp_identifier(c *grammar.AExp_identifierContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterAExp_IntergerLiteral(c *grammar.AExp_IntergerLiteralContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterAExp_listElementExpression(c *grammar.AExp_listElementExpressionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterAExprList(c *grammar.AExprListContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterBExpr_aExpr(c *grammar.BExpr_aExprContext) {}

func (p *ParseTreeListener) EnterBExpr_bExpr(c *grammar.BExpr_bExprContext) {}

func (p *ParseTreeListener) EnterBExpr_bang(c *grammar.BExpr_bangContext) {}

func (p *ParseTreeListener) EnterBExpr_bracketExpression(c *grammar.BExpr_bracketExpressionContext) {}

func (p *ParseTreeListener) EnterIntegerExpression(c *grammar.IntegerExpressionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterListElementIndex(c *grammar.ListElementIndexContext) {}

func (p *ParseTreeListener) EnterListElementIndexList(c *grammar.ListElementIndexListContext) {}

func (p *ParseTreeListener) EnterListElementExpression(c *grammar.ListElementExpressionContext) {}

func (p *ParseTreeListener) EnterAssignInit_aExpr(c *grammar.AssignInit_aExprContext) {}

func (p *ParseTreeListener) EnterAssignInit_listInitExpression(c *grammar.AssignInit_listInitExpressionContext) {
}

func (p *ParseTreeListener) EnterAssignInit_funcInitExpression(c *grammar.AssignInit_funcInitExpressionContext) {
}

func (p *ParseTreeListener) EnterAssignInitList(c *grammar.AssignInitListContext) {}

func (p *ParseTreeListener) EnterAssignStatement(c *grammar.AssignStatementContext) {}

func (p *ParseTreeListener) EnterIfExpr(c *grammar.IfExprContext) {}

func (p *ParseTreeListener) EnterElsifExpr(c *grammar.ElsifExprContext) {}

func (p *ParseTreeListener) EnterElseExpr(c *grammar.ElseExprContext) {}

func (p *ParseTreeListener) EnterTernaryIfExpr(c *grammar.TernaryIfExprContext) {}

func (p *ParseTreeListener) EnterTernaryElseExpr(c *grammar.TernaryElseExprContext) {}

func (p *ParseTreeListener) EnterSelectionStatement(c *grammar.SelectionStatementContext) {}

func (p *ParseTreeListener) EnterIterationStatement(c *grammar.IterationStatementContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterDefinition(c *grammar.DefinitionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterDefinitionList(c *grammar.DefinitionListContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterFile(c *grammar.FileContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterTypeSpecifier(c *grammar.TypeSpecifierContext) {}

func (p *ParseTreeListener) EnterTypeSpecifierList(c *grammar.TypeSpecifierListContext) {}

func (p *ParseTreeListener) EnterParaDeclaratorList(c *grammar.ParaDeclaratorListContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterParaDeclaratorWithIdentity(c *grammar.ParaDeclaratorWithIdentityContext) {
}

func (p *ParseTreeListener) EnterParaDeclaratorWithIdentityList(c *grammar.ParaDeclaratorWithIdentityListContext) {
}

func (p *ParseTreeListener) EnterFuncTypeSpecifier(c *grammar.FuncTypeSpecifierContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterFuncIdentifier(c *grammar.FuncIdentifierContext) {}

func (p *ParseTreeListener) EnterFuncTypeSpecifierWithName(c *grammar.FuncTypeSpecifierWithNameContext) {
}

func (p *ParseTreeListener) EnterFuncReturnPara(c *grammar.FuncReturnParaContext) {}

func (p *ParseTreeListener) EnterFuncReturnParaList(c *grammar.FuncReturnParaListContext) {}

func (p *ParseTreeListener) EnterFuncReturnStatement(c *grammar.FuncReturnStatementContext) {}

func (p *ParseTreeListener) EnterFuncStatement(c *grammar.FuncStatementContext) {}

func (p *ParseTreeListener) EnterFuncStatementList(c *grammar.FuncStatementListContext) {}

func (p *ParseTreeListener) EnterFuncBody(c *grammar.FuncBodyContext) {}

func (p *ParseTreeListener) EnterFuncInitExpression(c *grammar.FuncInitExpressionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) EnterFuncDefinition(c *grammar.FuncDefinitionContext) {}

func (p *ParseTreeListener) ExitIdentifier(c *grammar.IdentifierContext) {
	name := c.GetText()
	p.stack.Push(ast.IdentifierHelper.New(name))
}

func (p *ParseTreeListener) ExitDeclarator_identifier(c *grammar.Declarator_identifierContext) {
	// todo: remove tag
}

func (p *ParseTreeListener) ExitDeclarator_listElementExpression(c *grammar.Declarator_listElementExpressionContext) {
	// todo: remove tag
}

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

func (p *ParseTreeListener) ExitAExp_bracketExpression(c *grammar.AExp_bracketExpressionContext) {
	panic("implement me")
}

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

func (p *ParseTreeListener) ExitAExprList(c *grammar.AExprListContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitBExpr_aExpr(c *grammar.BExpr_aExprContext) {
	fmt.Println("c:", c.GetText())
	e2, _ := p.stack.PopByType(ast.NoderAExpr)
	e1, _ := p.stack.PopByType(ast.NoderAExpr)

	switch c.GetChild(1).(antlr.TerminalNode).GetText() {
	case "==":
		p.stack.Push(ast.BExprCompareHelper.New(e1.(ast.AExpr), e2.(ast.AExpr), ast.BExprCompareEQ))
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
	fmt.Println("todo ExitBExpr_bang")
}

func (p *ParseTreeListener) ExitBExpr_bracketExpression(c *grammar.BExpr_bracketExpressionContext) {
	fmt.Println("todo ExitBExpr_bracketExpression")
}

func (p *ParseTreeListener) ExitIntegerExpression(c *grammar.IntegerExpressionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitListElementIndex(c *grammar.ListElementIndexContext) {
	e, _ := p.stack.PopByType(ast.NoderAExpr) // todo: error
	p.stack.Push(ast.ListElementIndexHelper.New(e.(ast.AExpr)))
}

func (p *ParseTreeListener) ExitListElementIndexList(c *grammar.ListElementIndexListContext) {}

func (p *ParseTreeListener) ExitListElementExpression(c *grammar.ListElementExpressionContext) {
	list := []*ast.ListElementIndex{}
	for {
		item, _ := p.stack.PopByType(ast.NoderListElementIndex)
		if item == nil {
			break
		}

		list = append([]*ast.ListElementIndex{item.(*ast.ListElementIndex)}, list...)
	}

	name, _ := p.stack.PopByType(ast.NoderIdentifier)

	p.stack.Push(ast.ListElementExprHelper.New(name.(*ast.Identifier), list))
}

func (p *ParseTreeListener) ExitAssignInit_aExpr(c *grammar.AssignInit_aExprContext) {}

func (p *ParseTreeListener) ExitAssignInit_listInitExpression(c *grammar.AssignInit_listInitExpressionContext) {
}

func (p *ParseTreeListener) ExitAssignInit_funcInitExpression(c *grammar.AssignInit_funcInitExpressionContext) {
	fmt.Println("todo ExitAssignInit_funcInitExpression")
}

func (p *ParseTreeListener) ExitAssignInitList(c *grammar.AssignInitListContext) {}

func (p *ParseTreeListener) ExitAssignStatement(c *grammar.AssignStatementContext) {
	var initList []ast.AssignIniter
	for {
		item, _ := p.stack.PopByType(ast.NoderAssignIniter) // todo: error
		if item == nil {
			break
		}

		initList = append([]ast.AssignIniter{item.(ast.AssignIniter)}, initList...)
	}

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

	var ifExprList []*ast.IfExpr
	for {
		item, _ := p.stack.PopByType(ast.NoderIfExpr)
		if item == nil {
			break
		}

		ifExprList = append([]*ast.IfExpr{item.(*ast.IfExpr)}, ifExprList...)
	}

	if elseExpr != nil {
		p.stack.Push(ast.SelectionStmtHelper.New(ifExprList, elseExpr.(*ast.ElseExpr)))
	} else {
		p.stack.Push(ast.SelectionStmtHelper.New(ifExprList, nil))
	}
}

func (p *ParseTreeListener) ExitIterationStatement(c *grammar.IterationStatementContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitDefinition(c *grammar.DefinitionContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitDefinitionList(c *grammar.DefinitionListContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitFile(c *grammar.FileContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitTypeSpecifier(c *grammar.TypeSpecifierContext) {
}

func (p *ParseTreeListener) ExitTypeSpecifierList(c *grammar.TypeSpecifierListContext) {}

func (p *ParseTreeListener) ExitParaDeclaratorList(c *grammar.ParaDeclaratorListContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitParaDeclaratorWithIdentity(c *grammar.ParaDeclaratorWithIdentityContext) {
	typeSpecifier, _ := p.stack.PopByType(ast.NoderTypeSpecifierer)
	declList := p.readIdentifierList()
	p.stack.Push(ast.ParaDeclaratorWithIdentityHelper.New(declList, typeSpecifier.(ast.TypeSpecifierer)))
}

func (p *ParseTreeListener) ExitParaDeclaratorWithIdentityList(c *grammar.ParaDeclaratorWithIdentityListContext) {
}

func (p *ParseTreeListener) ExitFuncTypeSpecifier(c *grammar.FuncTypeSpecifierContext) {
	panic("implement me")
}

func (p *ParseTreeListener) ExitFuncIdentifier(c *grammar.FuncIdentifierContext) {
	name, _ := p.stack.PopByType(ast.NoderIdentifier) // todo: error
	p.stack.Push(ast.FuncIdentityHelper.New(name.(*ast.Identifier)))
}

func (p *ParseTreeListener) ExitFuncTypeSpecifierWithName(c *grammar.FuncTypeSpecifierWithNameContext) {
	returnList := p.readTypeSpecifierList()
	paraList := p.readParaDeclaratorWithIdentityList()
	name, _ := p.stack.PopByType(ast.NoderFuncIdentifier) // todo: error
	p.stack.Push(ast.FuncTypeSpecifierWithNameHelper.New(name.(*ast.FuncIdentifier), paraList, returnList))
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
	panic("implement me")
}

func (p *ParseTreeListener) ExitFuncDefinition(c *grammar.FuncDefinitionContext) {
	stmt := p.readFuncStatementerList()
	typeSpecifier, _ := p.stack.PopByType(ast.NoderFuncTypeSpecifierWithName)
	p.stack.Push(ast.FuncDefinitionHelper.New(typeSpecifier.(*ast.FuncTypeSpecifierWithName), stmt))
}
