package transformer

import (
	"fmt"
	"testing"

	"github.com/AleckDarcy/210A/zz/ast"
)

func Helper(t *testing.T, noder ast.BasicNoder) {
	tr := NewTransformer()

	//fmt.Println(noder)
	fmt.Println(tr.Walk(noder))
}

func TestTransformer_WalkAExprArith(t *testing.T) {
	Helper(t, ast.AExprAdd2)
}

func TestTransformer_WalkListInitExpr(t *testing.T) {
	Helper(t, ast.ListInitExpr1)
}

func TestTransformer_WalkAssignStmt(t *testing.T) {
	Helper(t, ast.AssignStmt1)
}

func TestTransformer_WalkFuncDefinition(t *testing.T) {
	Helper(t, ast.FuncDefinition2)
}

func TestTransformer_WalkFile(t *testing.T) {
	Helper(t, ast.FileHelper.New([]ast.Definitioner{
		ast.AssignStmt1,
		ast.AssignStmt3,
		ast.FuncDefinition1,
		ast.FuncDefinition2,
	}))
}
