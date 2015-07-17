package main

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/eandre/lunar"
)

func main() {
	dir := os.Args[1]
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		log.Fatalln("Could not parse packages:", err)
	}

	if len(pkgs) != 1 {
		log.Fatalf("Found multiple packages in directory: %v", pkgs)
	}

	pkgname := filepath.Clean(filepath.Base(dir))

	target := "output"
	if err := os.MkdirAll(target, 0755); err != nil {
		log.Fatalln("Could not create directory:", err)
	}

	for _, pkg := range pkgs {
		file, err := lunar.MergeFiles(fset, pkg.Files, pkgname)
		if err != nil {
			log.Fatalln("Could not merge files:", err)
		}

		format.Node(os.Stdout, fset, file)

		parser, err := lunar.NewParser(pkg.Name, fset, []*ast.File{file})
		if err != nil {
			log.Fatalln("Could not create parser:", err)
		}
		parser.MarkTransientPackage("github.com/eandre/sbm/wow")

		out, err := os.Create(filepath.Join(target, pkgname+".lua"))
		if err != nil {
			log.Fatalln("Could not create file:", err)
		}
		defer out.Close()

		pe := parser.ParseNode(out, file)
		if pe != nil {
			log.Fatalln("Could not parse node:", pe)
		}
	}
}
