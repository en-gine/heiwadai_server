package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory to lint")
		os.Exit(1)
	}

	dir := os.Args[1]
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		return lintFile(path)
	})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func lintFile(filePath string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Inspect(node, func(n ast.Node) bool {
		compLit, ok := n.(*ast.CompositeLit)
		if !ok {
			return true
		}

		if ident, ok := compLit.Type.(*ast.Ident); ok {
			if ident.Obj != nil && ident.Obj.Kind == ast.Typ {
				fmt.Printf("%s: line %d: direct struct initialization of %s is discouraged\n",
					filePath, fset.Position(compLit.Pos()).Line, ident.Name)
			}
		}

		return true
	})

	return nil
}
