package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "tracks/segment.go", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Decls)
	for _, elem := range f.Decls {
		switch x := elem.(type) {
		case *ast.GenDecl:
			if x.Tok == token.VAR {
				fmt.Println("it is a var")
				fmt.Println(x.Tok)
				spec := x.Specs[0].(*ast.ValueSpec)
				if spec.Names[0].String() == "TS_MAP" {
					fmt.Println("found the ts map")
				}
				vls := spec.Values
			}
		}
	}
}
