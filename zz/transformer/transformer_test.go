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
	case *ast.CollectionElementExpr:
		return t.WalkCollectionElementExpr(node, false)
	case *ast.AssignStmt:
		return t.WalkAssignStmt(node, true)
	case *ast.IterationStmt:
		return t.WalkIterationStmt(node)
	case *ast.FuncDefinition:
		return t.WalkFuncDefinition(node)
	case *ast.FuncExecuteExpression:
		return t.WalkFuncExecuteExpression(node)
	case *ast.FuncExecuteStatement:
		return t.WalkFuncStatementer(node)
	case *ast.PrintStatement:
		return t.WalkFuncStatementer(node)
	case *ast.SelectionStmt:
		return t.WalkFuncStatementer(node)
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
	Helper(t, ast.AssignStmt8)
	Helper(t, ast.AssignStmt9)
	Helper(t, ast.AssignStmt10)
	Helper(t, ast.AExprAdd3)
	Helper(t, ast.AssignStmt11)
	Helper(t, ast.AssignStmt12)

	Helper(t, ast.IterationStmtHelper.New(
		ast.AssignStmtHelper.New(
			ast.AssignStmtFlagInit,
			[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("i"))},
			[]ast.AssignIniter{ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(0))}),
		ast.BExprCompareHelper.New(
			ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("i")),
			ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(2)),
			ast.BExprCompareLT,
		),
		ast.AssignStmtHelper.New(
			ast.AssignStmtFlagNormal,
			[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("i"))},
			[]ast.AssignIniter{ast.AExprArithHelper.New(
				ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("i")),
				ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(1)),
				ast.AExprArithAdd,
			)},
		),

		[]ast.FuncStatementer{
			ast.IterationStmtHelper.New(
				ast.AssignStmtHelper.New(
					ast.AssignStmtFlagInit,
					[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("j"))},
					[]ast.AssignIniter{ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(0))}),
				ast.BExprCompareHelper.New(
					ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("j")),
					ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(2)),
					ast.BExprCompareLT,
				),
				ast.AssignStmtHelper.New(
					ast.AssignStmtFlagNormal,
					[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("j"))},
					[]ast.AssignIniter{ast.AExprArithHelper.New(
						ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("j")),
						ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(1)),
						ast.AExprArithAdd,
					)},
				),

				[]ast.FuncStatementer{
					ast.AssignStmtHelper.New(
						ast.AssignStmtFlagNormal,
						[]ast.Declaratorer{
							ast.CollectionElementExprHelper.New(
								ast.IdentifierHelper.New("e"),
								[]*ast.CollectionElementIndex{
									ast.CollectionElementIndexHelper.New(
										ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("i")),
									),
									ast.CollectionElementIndexHelper.New(
										ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("j")),
									),
								},
							),
						},
						[]ast.AssignIniter{
							ast.AExprArithHelper.New(
								ast.AExprArithHelper.New(
									ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("i")),
									ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(2)),
									ast.AExprArithMul,
								),
								ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("j")),
								ast.AExprArithAdd,
							),
						},
					),
				},
			),
		},
	))

	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagInit,
		[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("f"))},
		[]ast.AssignIniter{
			ast.AExprArithHelper.New(
				ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("e")),
				ast.AExprSimpleHelper.New(ast.ArithTransposeHelper.New(ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("e")))),
				ast.AExprArithMul,
			),
		},
	))

	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagInit,
		[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("f1"))},
		[]ast.AssignIniter{
			ast.AExprArithHelper.New(
				ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("e")),
				ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("h")),
				ast.AExprArithMul,
			),
		},
	))

	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagInit,
		[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("f2"))},
		[]ast.AssignIniter{
			ast.AExprArithHelper.New(
				ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("h")),
				ast.AExprSimpleHelper.New(ast.IdentifierHelper.New("e")),
				ast.AExprArithMul,
			),
		},
	))
}

func TestTransformer_WalkBExprBinary(t *testing.T) {
	Helper(t, ast.BExprBinary4)
	Helper(t, ast.BExprBinary5)
	Helper(t, ast.BExprCompare4)
	Helper(t, ast.BExprCompare5)
	Helper(t, ast.BExprCompare6)
	Helper(t, ast.BExprCompare7)
}

func TestTransformer_WalkListInitExpr(t *testing.T) {
	Helper(t, ast.ListInitExpr1)
}

func TestTransformer_WalkAssignStmt(t *testing.T) {
	Helper(t, ast.AssignStmt4)
	Helper(t, ast.AssignStmt13)
	Helper(t, ast.IterationStmtHelper.New(
		ast.AssignStmtHelper.New(
			ast.AssignStmtFlagInit,
			[]ast.Declaratorer{
				ast.DeclaratorHelper.New(ast.IdentifierHelper.New("i")),
				ast.DeclaratorHelper.New(ast.IdentifierHelper.New("j")),
			},
			[]ast.AssignIniter{
				ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(0)),
				ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(0)),
			}),
		nil,
		nil,
		[]ast.FuncStatementer{},
	))
}

func TestTransformer_WalkFuncDefinition(t *testing.T) {
	Helper(t, ast.FuncDefinition1)
	Helper(t, ast.FuncDefinition2)
	Helper(t, ast.FuncDefinition3)
	Helper(t, ast.FuncDefinition4)
	Helper(t, ast.FuncExecuteStatement1)
	Helper(t, ast.FuncExecuteExpression2)
	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagInit,
		[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("out_1"))},
		[]ast.AssignIniter{ast.FuncExecuteExpression1}))
	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagInit,
		[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("out_2"))},
		[]ast.AssignIniter{ast.AExprSimpleHelper.New(ast.FloatLiteralHelper.New(2.1))}))
	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagInit,
		[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("out_3"))},
		[]ast.AssignIniter{ast.AExprSimpleHelper.New(ast.FloatLiteralHelper.New(2.1))}))
	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagNormal,
		[]ast.Declaratorer{
			ast.DeclaratorHelper.New(ast.IdentifierHelper.New("out_2")),
			ast.DeclaratorHelper.New(ast.IdentifierHelper.New("out_3")),
		},
		//[]ast.AssignIniter{ast.FuncExecuteExpression1}))
		[]ast.AssignIniter{
			ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(2)),
			ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(2)),
		}))
	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagInit,
		[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("out_4"))},
		[]ast.AssignIniter{ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(2))}))
	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagInit,
		[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("out_5"))},
		[]ast.AssignIniter{ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(2))}))
	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagNormal,
		[]ast.Declaratorer{
			ast.DeclaratorHelper.New(ast.IdentifierHelper.New("out_4")),
			ast.DeclaratorHelper.New(ast.IdentifierHelper.New("out_5")),
		},
		//[]ast.AssignIniter{ast.FuncExecuteExpression1}))
		[]ast.AssignIniter{
			ast.AExprSimpleHelper.New(ast.FloatLiteralHelper.New(2)),
			ast.AExprSimpleHelper.New(ast.FloatLiteralHelper.New(2)),
		}))
}

func TestTransformer_WalkSelectionStmt(t *testing.T) {
	Helper(t, ast.AssignStmtHelper.New(
		ast.AssignStmtFlagInit,
		[]ast.Declaratorer{ast.DeclaratorHelper.New(ast.IdentifierHelper.New("b"))},
		[]ast.AssignIniter{ast.AExprSimpleHelper.New(ast.IntegerLiteralHelper.New(0))}),
	)
	Helper(t, ast.AssignStmt2)
	Helper(t, ast.SelectionStmt1)
	Helper(t, ast.SelectionStmt3)
}

func TestTransformer_WalkPrintStatement(t *testing.T) {
	Helper(t, ast.PrintStatement1)
}
func TestTransformer_WalkFile(t *testing.T) {
	Helper(t, ast.FileHelper.New([]ast.Definitioner{
		ast.AssignStmt6,
		ast.FuncDefinition1,
		ast.FuncDefinition2,
	}))
}

//
//func dd() {
//	var c = make([]int, 4)
//	c.Set(1, 10)
//	c.Get(1) + c.Get(1)
//}
