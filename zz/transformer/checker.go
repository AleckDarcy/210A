package transformer

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/AleckDarcy/210A/zz/ast"
)

type DefinitionType int64

const (
	DefinitionFunction DefinitionType = 1 << iota >> 1
	DefinitionVariant
	DefinitionMatrix
)

func (t DefinitionType) IsFunction() bool {
	return t&DefinitionFunction != 0
}

func (t DefinitionType) IsVariant() bool {
	return t&DefinitionVariant != 0
}

func (t DefinitionType) IsMatrix() bool {
	return t&DefinitionMatrix != 0
}

type AssignStmtType int64

const (
	AssignStmtVariant AssignStmtType = iota
	AssignStmtMatrix
	AssignStmtFunctionReturn
)

type VariantType int64

const (
	VariantInvalid VariantType = 1 << iota >> 1
	VariantGlobal
	VariantInteger
	VariantFloat
	VariantBinary
	VariantList
	VariantMatrix
	VariantFunction
)

func (t VariantType) PriorLevel() int {
	if t.IsFunction() {
		return 10
	} else if t.IsMatrix() || t.IsList() {
		return 10
	} else if t.IsFloat() {
		return 5
	} else if t.IsInterger() {
		return 1
	} else if t.IsBinary() {
		return 0
	}

	panic("unreachable code")
}

func (t VariantType) PriorTo(a VariantType) bool {
	level1 := t.PriorLevel()
	level2 := a.PriorLevel()

	return level1 > level2
}

func (t VariantType) IsInvalid() bool {
	return t == VariantInvalid
}

func (t VariantType) IsGloabl() bool {
	return t&VariantGlobal != 0
}

func (t VariantType) IsInterger() bool {
	return t&VariantInteger != 0
}

func (t VariantType) IsFloat() bool {
	return t&VariantFloat != 0
}

func (t VariantType) IsBinary() bool {
	return t&VariantBinary != 0
}

func (t VariantType) IsList() bool {
	return t&VariantList != 0
}

func (t VariantType) IsMatrix() bool {
	return t&VariantMatrix != 0
}

func (t VariantType) IsFunction() bool {
	return t&VariantFunction != 0
}

type VariantInfo struct {
	name string
	typ  VariantType
	used bool

	listNode    *ast.ListTypeSpecifier // list
	matrixNode  *ast.MatrixInitExpr    // matrix
	paraTypes   []VariantType
	returnTypes []VariantType
	returnNodes []ast.BasicNoder
}

func NewVariantInfo(name string, typ VariantType) *VariantInfo {
	return &VariantInfo{name: name, typ: typ}
}

func (i *VariantInfo) Name() string {
	return i.name
}

func (i *VariantInfo) Type() VariantType {
	return i.typ
}

type VariantInfos struct {
	variants map[string]*VariantInfo
}

type checker struct {
	*StoreStack
}

var Checker = &checker{
	StoreStack: &StoreStack{variants: map[string]*VariantInfo{}},
}

func (c *checker) EnterBlock() {
	c.StoreStack = &StoreStack{
		variants: map[string]*VariantInfo{},
		Parent:   c.StoreStack,
	}
}

func (c *checker) ExitBlock() {
	c.StoreStack = c.StoreStack.Parent
}

func (c *checker) CheckIdentifierType(node *ast.Identifier) (VariantType, error) {
	name := node.Name()
	if info, ok := c.Get(name); ok {
		return info.typ, nil
	}

	return VariantInvalid, errors.New(fmt.Sprintf("check idenfier type err: identifier %s not found", node.Name()))
}

func (c *checker) CheckAExprType(noder ast.AExpr) (VariantType, error) {
	switch node := noder.(type) {
	case *ast.Identifier:
		return c.CheckIdentifierType(node)
	case *ast.IntegerLiteral:
		return VariantInteger, nil
	case *ast.FloatLiteral:
		return VariantFloat, nil
	case *ast.AExprSimple:
		return c.CheckAExprType(node.E())
	case *ast.AExprArith:
		return c.CheckAExprArithType(node)
	case *ast.CollectionElementExpr:
		t, err := c.CheckIdentifierType(node.Identifier())
		if err != nil {
			return VariantInvalid, err
		}

		if t.IsMatrix() || t.IsList() {
			return t, err
		}

		return VariantInvalid, errors.New("variant is not a collection")
	default:
		fmt.Println("ssssss", reflect.TypeOf(noder).Elem().Name())
		panic("unreachable code")
	}
}

func (c *checker) GetDepthOfCollectionElementExpr(elem *ast.ListElementTypeSpecifier) int {
	tmpr := elem.Elem()
	depth := 1
	end := false
	for {
		switch tmp := tmpr.(type) {
		case *ast.ListElementTypeSpecifier:
			tmpr = tmp.Elem()
		case *ast.SimpleTypeSpecifier:
			end = true
		default:
			panic("unreachable code")
		}

		if end {
			break
		}
		depth++
	}

	return depth
}

func (c *checker) CheckAExprArithType(node *ast.AExprArith) (VariantType, error) {
	t1, err1 := c.CheckAExprType(node.E1())
	t2, err2 := c.CheckAExprType(node.E2())
	if err1 != nil {
		return VariantInvalid, err1
	} else if err2 != nil {
		return VariantInvalid, err2
	}

	getEleType := func(t VariantType, node *ast.CollectionElementExpr) (VariantType, error) {
		name := node.Name().Name()
		info, ok := c.Get(name)
		if !ok {
			return VariantInvalid, errors.New(fmt.Sprintf("variant %s undefined", name))
		}

		depth := c.GetDepthOfCollectionElementExpr(info.listNode.Elem())

		if depth != len(node.List()) {
			return VariantInvalid, errors.New("list index expression is not a simple type")
		}

		if t.IsInterger() {
			return VariantInteger, nil
		}

		return VariantFloat, nil
	}

	if t1.IsList() {
		if t1, err1 = getEleType(t1, node.E1().(*ast.CollectionElementExpr)); err1 != nil {
			return VariantInvalid, err1
		}
	}
	if t2.IsList() {
		if t2, err2 = getEleType(t2, node.E2().(*ast.CollectionElementExpr)); err2 != nil {
			return VariantInvalid, err2
		}
	} else if t1.IsBinary() || t2.IsBinary() {
		return VariantInvalid, errors.New("arithmetical operations do not support binary")
	}

	if t1 == t2 {
		if t1.IsMatrix() && node.Op() == ast.AExprArithDiv {
			return VariantInvalid, errors.New("matrix cannot be the divisor")
		}

		return t1, nil
	}

	if t1.PriorTo(t2) {
		return t1, nil
	}

	return t2, nil
}

func (c *checker) CheckArithTransposeType(node *ast.ArithTranspose) (VariantType, error) {
	typ, err := c.CheckAExprType(node.E())
	if err != nil {
		return typ, err
	}

	if typ.IsMatrix() {
		return typ, nil
	}

	return VariantInvalid, errors.New("cannot transpose non-matrix")
}

func (c *checker) CheckSimpleTypeSpecifierType(node *ast.SimpleTypeSpecifier) (VariantType, error) {
	switch node.Name() {
	case "int":
		return VariantInteger, nil
	case "float":
		return VariantFloat, nil
	case "matrix":
		return VariantMatrix, nil
	default:
		panic("unreachable code")
	}
}

func (c *checker) CheckTypeSpecifiererType(noder ast.TypeSpecifierer) (VariantType, error) {
	switch node := noder.(type) {
	case *ast.SimpleTypeSpecifier:
		return c.CheckSimpleTypeSpecifierType(node)
	case *ast.ListTypeSpecifier:
		return c.CheckListTypeSpecifierType(node)
	//case *ast.FuncTypeSpecifier:
	//	return t.WalkFuncTypeSpecifier(node)
	default:
		panic("unreachable code")
	}
}

func (c *checker) CheckListInitExprType(node *ast.ListInitExpr) (VariantType, error) {
	return c.CheckListTypeSpecifierType(node.TypeSpecifier())
}

func (c *checker) CheckListTypeSpecifierType(node *ast.ListTypeSpecifier) (VariantType, error) {
	return c.CheckListElementTypeSpecifierType(node.Elem())
}

func (c *checker) CheckListElementTypeSpecifierType(node *ast.ListElementTypeSpecifier) (VariantType, error) {
	return c.CheckListElementTypeSpecifiererType(node.Elem())
}

func (c *checker) CheckListElementTypeSpecifiererType(noder ast.ListElementTypeSpecifierer) (VariantType, error) {
	switch node := noder.(type) {
	case *ast.SimpleTypeSpecifier:
		if node.Name() == "int" {
			return VariantList | VariantInteger, nil
		}

		return VariantList | VariantFloat, nil
	case *ast.ListElementTypeSpecifier:
		return c.CheckListElementTypeSpecifierType(node)
	default:
		panic("unreachable code")
	}
}

func (c *checker) CheckAssignIniterType(noder ast.AssignIniter) ([]VariantType, []ast.BasicNoder, error) {
	var t VariantType
	var err error

	switch node := noder.(type) {
	case *ast.IntegerLiteral:
		t = VariantInteger
	case *ast.FloatLiteral:
		t = VariantFloat
	case *ast.AExprSimple:
		t, err = c.CheckAExprType(node.E())
	case *ast.AExprArith:
		t, err = c.CheckAExprArithType(node)
	case *ast.ArithTranspose:
		t, err = c.CheckArithTransposeType(node)
	case *ast.ListInitExpr:
		t, err = c.CheckListInitExprType(node)
	case *ast.MatrixInitExpr:
		t = VariantFloat | VariantMatrix
	case *ast.FuncExecuteExpression:
		return c.CheckFuncExecuteExpressionReturnType(node)
	default:
		panic("unreachable code")
	}

	if err != nil {
		return nil, nil, err
	}

	return []VariantType{t}, []ast.BasicNoder{noder}, err
}

func (c *checker) CheckAssignStmtType(node *ast.AssignStmt) ([]VariantType, []VariantType, error) {
	declList, initList := node.DeclList(), node.InitList()

	lenDeclList, lenInitList := len(declList), len(initList)

	var declTypeList, initTypeList []VariantType
	var initNodeList []ast.BasicNoder
	var err error

	if node.Flag().IsInit() {
		for _, decl := range declList {
			name := decl.Identifier().Name()
			if _, ok := c.GetFromCurrentStack(name); ok {
				return nil, nil, errors.New(fmt.Sprintf("duplicated definition of variant %s", name))
			}
		}
	} else {
		for _, decl := range declList {
			t, err := c.CheckDeclaratorType(decl)
			if err != nil {
				return nil, nil, err
			}

			declTypeList = append(declTypeList, t)
		}
	}

	if lenInitList == 1 {
		init := initList[0]
		if funcExecExpr, ok := init.(*ast.FuncExecuteExpression); ok { // todo: check paras
			initTypeList, initNodeList, err = c.CheckFuncExecuteExpressionReturnType(funcExecExpr)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	if lenDeclList == lenInitList {
		for _, init := range initList {
			if _, ok := init.(*ast.FuncExecuteExpression); ok {
				return nil, nil, errors.New("function execution should be the only assigner")
			}

			typeList, noders, err := c.CheckAssignIniterType(init)
			if err != nil {
				return nil, nil, err
			}

			initTypeList = append(initTypeList, typeList...)
			initNodeList = append(initNodeList, noders...)
		}
	} else {
		fmt.Println("=============================")
		fmt.Println(node)
		fmt.Println("=============================")

		return nil, nil, errors.New(fmt.Sprintf("values not match %d != %d", lenDeclList, lenInitList))
	}

	if node.Flag().IsInit() {
		for i, decl := range declList {
			name, typ := decl.Identifier().Name(), initTypeList[i]
			if typ.IsList() {
				c.variants[name] = &VariantInfo{
					name:     name,
					typ:      typ,
					listNode: initNodeList[i].(*ast.ListInitExpr).TypeSpecifier(),
				}
			} else if typ.IsMatrix() {
				var matrixNode *ast.MatrixInitExpr

				switch tmp := initNodeList[i].(type) {
				case *ast.MatrixInitExpr:
					matrixNode = tmp
				case *ast.AExprArith:
					info1, _ := c.Get(tmp.E1().(*ast.AExprSimple).E().(*ast.Identifier).Name())
					info2, _ := c.Get(tmp.E1().(*ast.AExprSimple).E().(*ast.Identifier).Name())

					sizes1, sizes2 := info1.matrixNode.Sizes(), info2.matrixNode.Sizes()
					if len(sizes1) != 2 || len(sizes2) != 2 {
						if sizes1 == nil && sizes2 == nil { // matrix as parameter
							matrixNode = ast.MatrixInitExprHelper.New(nil)
						} else {
							return nil, nil, errors.New("todo")
						}
					} else {
						matrixNode = ast.MatrixInitExprHelper.New([]ast.AExpr{sizes1[0], sizes2[1]})
					}
				}

				c.variants[name] = &VariantInfo{
					name:       name,
					typ:        typ,
					matrixNode: matrixNode,
				}
			} else {
				c.Set(name, &VariantInfo{name: name, typ: typ})
			}
		}

		return initTypeList, initTypeList, nil
	}

	return declTypeList, initTypeList, nil
}

func (c *checker) CheckCollectionElementExprType(node *ast.CollectionElementExpr) (VariantType, VariantType, error) {
	info, ok := c.Get(node.Identifier().Name())
	if !ok {
		return VariantInvalid, VariantInvalid, errors.New(fmt.Sprintf("variant %s not defined", node.Identifier().Name()))
	}

	t := info.typ

	if t.IsList() {
		depth := c.GetDepthOfCollectionElementExpr(info.listNode.Elem())
		if depth > len(node.List()) {
			return VariantList, t, nil
		} else if depth == len(node.List()) {
			if t.IsInterger() {
				return VariantList, VariantInteger, nil
			}

			return VariantList, VariantFloat, nil
		}

		return VariantInvalid, VariantInvalid, errors.New("invalid list expression")
	}

	if len(info.matrixNode.Sizes()) != len(node.List()) {
		return VariantInvalid, VariantInvalid, errors.New("invalid matrix expression")
	}

	return VariantMatrix, VariantFloat, nil
}

func (c *checker) CheckDeclaratorType(noder ast.Declaratorer) (VariantType, error) {
	switch node := noder.(type) {
	case *ast.Identifier:
		return c.CheckIdentifierType(node)
	case *ast.CollectionElementExpr:
		_, typ, err := c.CheckCollectionElementExprType(node)
		return typ, err
	case *ast.Declarator:
		return c.CheckDeclaratorType(node.Declaratorer)
	default:
		panic("unreachable code")
	}
}

func (c *checker) CheckFuncExecuteParaerType(noder ast.FuncExecuteParaer) (VariantType, error) {
	switch node := noder.(type) {
	case *ast.AExprSimple:
		return c.CheckAExprType(node.E())
	case *ast.AExprArith:
		return c.CheckAExprArithType(node)
	//case *ast.BExprCompare: // todo
	//	return t.WalkBExprCompare(node)
	//case *ast.BExprBinary:
	//	return t.WalkBExprBinary(node)
	default:
		panic("unreachable code")
	}
}

func (c *checker) CheckFuncExecuteParaerListType(list []ast.FuncExecuteParaer) ([]VariantType, error) {
	listType := make([]VariantType, len(list))
	var err error
	for i, elem := range list {
		if listType[i], err = c.CheckFuncExecuteParaerType(elem); err != nil {
			return nil, err
		}
	}

	return listType, nil
}

func (c *checker) CheckFuncExecuteExpressionReturnType(node *ast.FuncExecuteExpression) ([]VariantType, []ast.BasicNoder, error) {
	name := node.Name().Name()
	info, ok := c.Get(name)
	if !ok {
		return nil, nil, errors.New("function not defined")
	}

	if info.typ.IsFunction() {
		return info.returnTypes, info.returnNodes, nil
	}

	return nil, nil, errors.New(fmt.Sprintf("%s is not a function", name))
}
