package parser

import (
	"fmt"

	"github.com/AleckDarcy/210A/zz/ast"
)

func (p *ParseTreeListener) printStack() {
	if len(p.stack.list) == 0 {
		fmt.Println("Empty")

		return
	}
	for i := len(p.stack.list) - 1; i >= 0; i-- {
		fmt.Println(i, p.stack.list[i])
	}
}

func (p *ParseTreeListener) readFuncStatementerList() []ast.FuncStatementer {
	var List []ast.FuncStatementer
	for {
		item, _ := p.stack.PopByType(ast.NoderFuncStatementer) // todo: error
		if item == nil {
			break
		}

		List = append([]ast.FuncStatementer{item.(ast.FuncStatementer)}, List...)
	}

	return List
}

func (p *ParseTreeListener) readFuncReturnParaList() []ast.FuncReturnParaer {
	var List []ast.FuncReturnParaer
	for {
		item, _ := p.stack.PopByType(ast.NoderFuncReturnPara) // todo: error
		if item == nil {
			break
		}

		List = append([]ast.FuncReturnParaer{item.(ast.FuncReturnParaer)}, List...)
	}

	return List
}

func (p *ParseTreeListener) readDeclaratorerList() []ast.Declaratorer {
	var List []ast.Declaratorer
	for {
		item, _ := p.stack.PopByType(ast.NoderDeclarator) // todo: error
		if item == nil {
			break
		}

		List = append([]ast.Declaratorer{item.(ast.Declaratorer)}, List...)
	}

	return List
}

func (p *ParseTreeListener) readIdentifierList() []*ast.Identifier {
	var List []*ast.Identifier
	for {
		item, _ := p.stack.PopByType(ast.NoderDeclarator) // todo: error
		if item == nil {
			break
		}

		List = append([]*ast.Identifier{item.(*ast.Identifier)}, List...)
	}

	return List
}

func (p *ParseTreeListener) readTypeSpecifierList() []ast.TypeSpecifierer {
	var List []ast.TypeSpecifierer
	for {
		item, _ := p.stack.PopByType(ast.NoderTypeSpecifierer) // todo: error
		if item == nil {
			break
		}

		List = append([]ast.TypeSpecifierer{item.(ast.TypeSpecifierer)}, List...)
	}

	return List
}

func (p *ParseTreeListener) readParaDeclaratorWithIdentityList() []*ast.ParaDeclaratorWithIdentity {
	var List []*ast.ParaDeclaratorWithIdentity
	for {
		item, _ := p.stack.PopByType(ast.NoderParaDeclaratorWithIdentity) // todo: error
		if item == nil {
			break
		}

		List = append([]*ast.ParaDeclaratorWithIdentity{item.(*ast.ParaDeclaratorWithIdentity)}, List...)
	}

	return List
}
