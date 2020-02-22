package transformer

import (
	"fmt"
	"testing"

	"github.com/AleckDarcy/210A/zz/ast"
)

func (t *Transformer) Walk(noder ast.BasicNoder) string {
	switch node := noder.(type) {
	case *ast.AExprArith:
		return t.WalkAExprArith(node)
	case ast.BExpr:
		return t.WalkBExpr(node)
	case *ast.ListInitExpr:
		return t.WalkListInitExpr(node)
	case *ast.AssignStmt:
		return t.WalkAssignStmt(node, true)
	case *ast.FuncDefinition:
		return t.WalkFuncDefinition(node)
	case *ast.File:
		return t.WalkFile(node)
	default:
		panic("undefined")
	}
}

func Helper(t *testing.T, noder ast.BasicNoder) {
	tr := NewTransformer()

	//fmt.Println(noder)
	fmt.Println(tr.Walk(noder))
}

func TestTransformer_WalkAExprArith(t *testing.T) {
	Helper(t, ast.AExprDiv2)
}

func TestTransformer_WalkBExprBinary(t *testing.T) {
	Helper(t, ast.BExprBinary2)
	Helper(t, ast.BExprCompare7)
}

func TestTransformer_WalkListInitExpr(t *testing.T) {
	Helper(t, ast.ListInitExpr1)
}

func TestTransformer_WalkAssignStmt(t *testing.T) {
	Helper(t, ast.AssignStmt4)
}

func TestTransformer_WalkFuncDefinition(t *testing.T) {
	Helper(t, ast.FuncDefinition2)
}

func TestTransformer_WalkFile(t *testing.T) {
	Helper(t, ast.FileHelper.New([]ast.Definitioner{
		ast.AssignStmt6,
		ast.FuncDefinition1,
		ast.FuncDefinition2,
	}))
}
