package main

import (
	"github.com/eandre/lunar"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	dir := os.Args[1]
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		log.Fatalln("Could not parse packages:", err)
	}

	for name, pkg := range pkgs {
		file := ast.MergePackageFiles(pkg, ast.FilterUnassociatedComments|ast.FilterImportDuplicates)
		out, err := os.Create(name + ".lua")
		if err != nil {
			log.Fatalln("Could not create file:", err)
		}
		defer out.Close()

		parser, err := lunar.NewParser("brf", fset, []*ast.File{file})
		if err != nil {
			log.Fatalln("Could not create parser:", err)
		}
		parser.MarkTransientPackage("github.com/eandre/sbm/wow")

		pe := parser.ParseNode(out, file)
		logFile, err := os.Create("output.log")
		if err != nil {
			log.Fatalln("Could not create log file:", err)
		}
		defer logFile.Close()
		ast.Fprint(logFile, fset, file, nil)
		if pe != nil {
			log.Fatalln("Could not parse node:", pe)
		}
	}
}
