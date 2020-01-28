package parser

import (
	"fmt"
	"testing"

	"github.com/AleckDarcy/210A/zz/ast"
)

func (p *ParseTreeListener) PrintStack() {
	length := p.stack.Depth()
	if length == 0 {
		fmt.Println("Empty")

		return
	}
	for i := length - 1; i >= 0; i-- {
		fmt.Println(i, p.stack.list[i])
	}
}

func (p *ParseTreeListener) PopAll(t *testing.T) {
	for {
		item, _ := p.stack.PopByType(ast.NoderBasic)
		if item == nil {
			return
		}

		t.Log(item)
	}
}

func (p *ParseTreeListener) readListElementIndexList() []*ast.ListElementIndex {
	var list []*ast.ListElementIndex
	for {
		item, _ := p.stack.PopByType(ast.NoderListElementIndex) // todo: error
		if item == nil {
			break
		}

		list = append([]*ast.ListElementIndex{item.(*ast.ListElementIndex)}, list...)
	}

	return list
}

func (p *ParseTreeListener) readAssignIniterList() []ast.AssignIniter {
	var list []ast.AssignIniter
	for {
		item, _ := p.stack.PopByType(ast.NoderAssignIniter) // todo: error
		if item == nil {
			break
		}

		list = append([]ast.AssignIniter{item.(ast.AssignIniter)}, list...)
	}

	return list
}

func (p *ParseTreeListener) readIfExprList() []*ast.IfExpr {
	var list []*ast.IfExpr
	for {
		item, _ := p.stack.PopByType(ast.NoderIfExpr) // todo: error
		if item == nil {
			break
		}

		list = append([]*ast.IfExpr{item.(*ast.IfExpr)}, list...)
	}

	return list
}

func (p *ParseTreeListener) readFuncStatementerList() []ast.FuncStatementer {
	var list []ast.FuncStatementer
	for {
		item, _ := p.stack.PopByType(ast.NoderFuncStatementer) // todo: error
		if item == nil {
			break
		}

		list = append([]ast.FuncStatementer{item.(ast.FuncStatementer)}, list...)
	}

	return list
}

func (p *ParseTreeListener) readFuncReturnParaList() []ast.FuncReturnParaer {
	var list []ast.FuncReturnParaer
	for {
		item, _ := p.stack.PopByType(ast.NoderFuncReturnPara) // todo: error
		if item == nil {
			break
		}

		list = append([]ast.FuncReturnParaer{item.(ast.FuncReturnParaer)}, list...)
	}

	return list
}

func (p *ParseTreeListener) readDeclaratorerList() []ast.Declaratorer {
	var list []ast.Declaratorer
	for {
		item, _ := p.stack.PopByType(ast.NoderDeclarator) // todo: error
		if item == nil {
			break
		}

		list = append([]ast.Declaratorer{item.(ast.Declaratorer)}, list...)
	}

	return list
}

func (p *ParseTreeListener) readTypeSpecifierList() []ast.TypeSpecifierer {
	var list []ast.TypeSpecifierer
	for {
		item, _ := p.stack.PopByType(ast.NoderTypeSpecifierer) // todo: error
		if item == nil {
			break
		}

		list = append([]ast.TypeSpecifierer{item.(ast.TypeSpecifierer)}, list...)
	}

	return list
}

func (p *ParseTreeListener) readParaDeclaratorListToIdentifierList() []*ast.Identifier {
	var list []*ast.Identifier
	for {
		item, _ := p.stack.PopByType(ast.NoderParaDeclarator) // todo: error
		if item == nil {
			break
		}

		list = append([]*ast.Identifier{item.(*ast.ParaDeclarator).Name()}, list...)
	}

	return list
}

func (p *ParseTreeListener) readParaDeclaratorWithIdentityList() []*ast.ParaDeclaratorWithIdentity {
	var list []*ast.ParaDeclaratorWithIdentity
	for {
		item, _ := p.stack.PopByType(ast.NoderParaDeclaratorWithIdentity) // todo: error
		if item == nil {
			break
		}

		list = append([]*ast.ParaDeclaratorWithIdentity{item.(*ast.ParaDeclaratorWithIdentity)}, list...)
	}

	return list
}

func (p *ParseTreeListener) readDefinitionerList() []ast.Definitioner {
	var list []ast.Definitioner
	for {
		item, _ := p.stack.PopByType(ast.NoderDefinitioner) // todo: error
		if item == nil {
			break
		}

		list = append([]ast.Definitioner{item.(ast.Definitioner)}, list...)
	}

	return list
}

func (p *ParseTreeListener) readFuncExecuteParaListToFuncExecuteParaerList() []ast.FuncExecuteParaer {
	var list []ast.FuncExecuteParaer
	for {
		item, _ := p.stack.PopByType(ast.NoderFuncExecutePara) // todo: error
		if item == nil {
			break
		}

		list = append([]ast.FuncExecuteParaer{item.(*ast.FuncExecutePara).Para()}, list...)
	}

	return list
}
