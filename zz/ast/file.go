package ast

import "fmt"

type File struct {
	defList []Definitioner
}

func (f *File) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sFile {\n"+
		"%s..DefinistionList:\n"+
		"%s\n"+
		"%s}",
		ident, ident, IterableToString(ident+"....", IteratableDefinitioner(f.defList)),
		ident,
	)
}

func (f *File) String() string {
	return f.toString("")
}
