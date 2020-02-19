package ast

import "fmt"

type PrintStatement struct {
	list []BasicNoder
}

var PrintStatementHelper *PrintStatement

func (s *PrintStatement) New(list []BasicNoder) *PrintStatement {
	return &PrintStatement{list: list}
}

func (s *PrintStatement) funcStatementer() {}

func (s *PrintStatement) toString(indent string) string {
	return fmt.Sprintf(""+
		"%sPrintStatement {\n"+
		"%s..InitList:\n"+
		"%s\n"+
		"%s}",
		indent, indent, IterableToString(indent+"....", IteratableBasicNoderList(s.list)),
		indent,
	)
}

func (s *PrintStatement) String() string {
	return s.toString("")
}

func (s *PrintStatement) List() []BasicNoder {
	return s.list
}
