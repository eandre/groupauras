package main

import (
	"go/ast"
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

	target := filepath.Clean(filepath.Join("output", filepath.Clean(dir)))
	if err := os.MkdirAll(target, 0755); err != nil {
		log.Fatalln("Could not create directory:", err)
	}

	for _, pkg := range pkgs {
		var files []*ast.File
		for _, file := range pkg.Files {
			files = append(files, file)
		}

		parser, err := lunar.NewParser(pkg.Name, fset, files)
		if err != nil {
			log.Fatalln("Could not create parser:", err)
		}
		parser.MarkTransientPackage("github.com/eandre/sbm/wow")

		for name, file := range pkg.Files {
			name = filepath.Base(name)
			out, err := os.Create(filepath.Join(target, name+".lua"))
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
}
