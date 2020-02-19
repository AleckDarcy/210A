package transformer

import (
	"fmt"

	"github.com/AleckDarcy/210A/zz/ast"
)

type StoreStack struct {
	variants map[string]*VariantInfo

	Parent *StoreStack
}

type StoreType int64

const (
	StoreInteger StoreType = iota
	StoreFloat
	StoreFuncDefinition
)

type StoreMeta struct {
	typ StoreType
}

func (s *StoreStack) Set(name string, info *VariantInfo) {
	if _, ok := s.variants[name]; ok {
		panic(name)
	}
	s.variants[name] = info
}

func (s *StoreStack) Update(name string) {

}

func (s *StoreStack) Get(name string) (*VariantInfo, bool) {
	info, ok := s.variants[name]
	if ok {
		return info, true
	}

	if s.Parent != nil {
		return s.Parent.Get(name)
	}

	return nil, false
}

func (s *StoreStack) GetCurrentStack() *StoreStack {
	return s
}

func (s *StoreStack) GetVariantInfos() map[string]*VariantInfo {
	return s.variants
}

func (s *StoreStack) GetFromCurrentStack(name string) (*VariantInfo, bool) {
	info, ok := s.variants[name]

	return info, ok
}

func (s *StoreStack) SetParaDeclaratorWithIdentity(node *ast.ParaDeclaratorWithIdentity) error {
	typ, err := Checker.CheckTypeSpecifiererType(node.TypeSpecifier())
	if err != nil {
		return err
	}

	for _, decl := range node.DeclList() {
		name := decl.Name()
		fmt.Println("para name:", name, ", type:", typ)
		if typ.IsList() {
			info := &VariantInfo{
				name:     name,
				typ:      typ,
				used:     true,
				listNode: node.TypeSpecifier().(*ast.ListTypeSpecifier),
			}

			s.variants[name] = info

			fmt.Printf("set list: %p %+v", info, info)
		} else if typ.IsMatrix() {
			info := &VariantInfo{
				name:       name,
				typ:        typ,
				used:       true,
				matrixNode: ast.MatrixInitExprHelper.New(nil),
			}

			s.variants[name] = info
		} else {
			s.variants[name] = &VariantInfo{
				name: name,
				typ:  typ,
				used: true,
			}
		}
	}

	return nil
}
