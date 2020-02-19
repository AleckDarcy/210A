package transformer

import (
	"fmt"

	"github.com/AleckDarcy/210A/zz/ast"
)

func (t *Transformer) WalkAExprList(list []ast.AExpr) string {
	var listStr string
	for i, elem := range list {
		if i == 0 {
			listStr = t.WalkAExpr(elem)
		} else {
			listStr += ", " + t.WalkAExpr(elem)
		}
	}

	return listStr
}

func (t *Transformer) WalkDeclaratorerList(list []ast.Declaratorer) string {
	var listStr string
	for i, elem := range list {
		if i == 0 {
			listStr = t.WalkDeclaratorer(elem)
		} else {
			listStr += ", " + t.WalkDeclaratorer(elem)
		}
	}

	return listStr
}

func (t *Transformer) WalkAssignIniterList(list []ast.AssignIniter) string {
	var listStr string
	for i, elem := range list {
		if i == 0 {
			listStr = t.WalkAssignIniter(elem)
		} else {
			listStr += ", " + t.WalkAssignIniter(elem)
		}
	}

	return listStr
}

func (t *Transformer) WalkIfExprList(list []*ast.IfExpr) string {
	var listStr string
	for i, elem := range list {
		if i == 0 {
			listStr = t.WalkIfExpr(elem)
		} else {
			listStr += fmt.Sprintf(" else %s", t.WalkIfExpr(elem))
		}
	}

	return listStr
}

func (t *Transformer) WalkParaDeclaratorWithIdentityList(list []*ast.ParaDeclaratorWithIdentity) string {
	var listStr string
	for i, elem := range list {
		if err := Checker.SetParaDeclaratorWithIdentity(elem); err != nil {
			fmt.Println("error:", err)
		}

		if i == 0 {
			listStr = t.WalkParaDeclaratorWithIdentity(elem)
		} else {
			listStr += fmt.Sprintf(", %s", t.WalkParaDeclaratorWithIdentity(elem))
		}
	}

	return listStr
}

func (t *Transformer) WalkTypeSpecifiererList(list []ast.TypeSpecifierer) string {
	var listStr string
	for i, elem := range list {
		if i == 0 {
			listStr = t.WalkTypeSpecifierer(elem)
		} else {
			listStr += fmt.Sprintf(", %s", t.WalkTypeSpecifierer(elem))
		}
	}

	return listStr
}

func (t *Transformer) WalkFuncExecuteParaerList(list []ast.FuncExecuteParaer) string {
	var listStr string
	for i, elem := range list {
		if i == 0 {
			listStr += fmt.Sprintf("%s", t.WalkFuncExecuteParaer(elem))
		} else {
			listStr += fmt.Sprintf(", %s", t.WalkFuncExecuteParaer(elem))
		}
	}

	return listStr
}

func (t *Transformer) WalkFuncStatementerList(list []ast.FuncStatementer) string {
	var listStr string
	for i, elem := range list {
		if i != len(list)-1 {
			listStr += fmt.Sprintf("%s\n", t.WalkFuncStatementer(elem))
		} else {
			listStr += fmt.Sprintf("%s", t.WalkFuncStatementer(elem))
		}
	}

	return listStr
}

func (t *Transformer) WalkFuncReturnParaerList(list []ast.FuncReturnParaer) string {
	var listStr string
	for i, elem := range list {
		if i == 0 {
			listStr = t.WalkFuncReturnParaer(elem)
		} else {
			listStr += fmt.Sprintf(", %s", t.WalkFuncReturnParaer(elem))
		}
	}

	return listStr
}

func (t *Transformer) WalkDefinitionerList(list []ast.Definitioner) string {
	var listStr string
	for i, elem := range list {
		if i == 0 {
			listStr = t.WalkDefinitioner(elem)
		} else {
			listStr += fmt.Sprintf("\n%s", t.WalkDefinitioner(elem))
		}
	}

	return listStr
}
