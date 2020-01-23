package parser

import (
	"errors"
	"fmt"

	"github.com/AleckDarcy/210A/zz/ast"
)

type Stack struct {
	list []ast.BasicNoder
}

func (s *Stack) Depth() int {
	return len(s.list)
}

func (s *Stack) Push(item ast.BasicNoder) {
	s.list = append(s.list, item)
}

func (s *Stack) PopByType(typ ast.NoderType) (ast.BasicNoder, error) {
	length := len(s.list)
	if length == 0 {
		return nil, errors.New("stack is empty")
	}

	item := (s.list)[length-1]

	switch typ {
	case ast.NoderBasic:
		if _, ok := item.(ast.BasicNoder); !ok {
			return nil, errors.New("item is not a BasicNoder")
		}
	case ast.NoderAExpr:
		if _, ok := item.(ast.AExpr); !ok {
			return nil, errors.New("item is not a AExpr")
		}
	case ast.NoderBExpr:
		if _, ok := item.(ast.BExpr); !ok {
			return nil, errors.New("item is not a BExpr")
		}
	case ast.NoderTypeSpecifierer:
		if _, ok := item.(ast.TypeSpecifierer); !ok {
			return nil, errors.New("item is not a TypeSpecifierer")
		}
	case ast.NoderListElementTypeSpecifierer:
		if _, ok := item.(ast.ListElementTypeSpecifierer); !ok {
			return nil, errors.New("item is not a ListElementTypeSpecifierer")
		}
	case ast.NoderDeclarator:
		if _, ok := item.(ast.Declaratorer); !ok {
			return nil, errors.New("item is not a Declaratorer")
		}
	case ast.NoderAssignIniter:
		if _, ok := item.(ast.AssignIniter); !ok {
			return nil, errors.New("item is not a AssignIniter")
		}
	case ast.NoderFuncStatementer:
		if _, ok := item.(ast.FuncStatementer); !ok {
			return nil, errors.New("item is not a FuncStatementer")
		}
	case ast.NoderFuncReturnPara:
		if _, ok := item.(ast.FuncReturnParaer); !ok {
			return nil, errors.New("item is not a FuncReturnParaer")
		}
	case ast.NoderDefinitioner:
		if _, ok := item.(ast.Definitioner); !ok {
			return nil, errors.New("item is not a Definitioner")
		}
	case ast.NoderIdentifier:
		if _, ok := item.(*ast.Identifier); !ok {
			return nil, errors.New("item is not a Identifier")
		}
	case ast.NoderAssignStatment:
		if _, ok := item.(*ast.AssignStmt); !ok {
			return nil, errors.New("item is not a AssignStmt")
		}
	case ast.NoderListElementIndex:
		if _, ok := item.(*ast.ListElementIndex); !ok {
			return nil, errors.New("item is not a ListElementIndex")
		}
	case ast.NoderListElementTypeSpecifier:
		if _, ok := item.(*ast.ListElementTypeSpecifier); !ok {
			return nil, errors.New("item is not a ListElementTypeSpecifier")
		}
	case ast.NoderListTypeSpecifier:
		if _, ok := item.(*ast.ListTypeSpecifier); !ok {
			return nil, errors.New("item is not a ListElementTypeSpecifier")
		}
	case ast.NoderIfExpr:
		if _, ok := item.(*ast.IfExpr); !ok {
			return nil, errors.New("item is not a IfExpr")
		}
	case ast.NoderIterationAssignStatement:
		if _, ok := item.(*ast.IterationAssignStmt); !ok {
			return nil, errors.New("item is not a IterationAssignStmt")
		}
	case ast.NoderElseExpr:
		if _, ok := item.(*ast.ElseExpr); !ok {
			return nil, errors.New("item is not a ElseExpr")
		}
	case ast.NoderParaDeclaratorWithIdentity:
		if _, ok := item.(*ast.ParaDeclaratorWithIdentity); !ok {
			return nil, errors.New("item is not a ParaDeclaratorWithIdentity")
		}
	case ast.NoderFuncIdentifier:
		if _, ok := item.(*ast.FuncIdentifier); !ok {
			return nil, errors.New("item is not a FuncIdentifier")
		}
	case ast.NoderFuncTypeSpecifier:
		if _, ok := item.(*ast.FuncTypeSpecifier); !ok {
			return nil, errors.New("item is not a FuncTypeSpecifier")
		}
	case ast.NoderFuncTypeSpecifierWithName:
		if _, ok := item.(*ast.FuncTypeSpecifierWithName); !ok {
			return nil, errors.New("item is not a FuncTypeSpecifierWithName")
		}
	default:
		panic(fmt.Sprintf("undefined NodeType %v", typ))
	}

	s.list = (s.list)[:length-1]

	return item, nil
}
