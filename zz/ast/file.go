package ast

import "fmt"

type File struct {
	defList []Definitioner
}

var FileHelper *File

func (f *File) New(defList []Definitioner) *File {
	return &File{defList: defList}
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

func (f *File) DefList() []Definitioner {
	return f.defList
}
