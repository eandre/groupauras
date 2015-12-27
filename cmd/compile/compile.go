package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/eandre/lunar"
	"golang.org/x/tools/go/loader"
	"strings"
)

func main() {
	var conf loader.Config
	_, err := conf.FromArgs(os.Args[1:], false)
	if err != nil {
		log.Fatalln("Could not load packages:", err)
	}

	prog, err := conf.Load()
	if err != nil {
		log.Fatalln("Could not parse packages:", err)
	}

	target := "output"
	if err := os.MkdirAll(target, 0755); err != nil {
		log.Fatalln("Could not create directory:", err)
	}

	parser := lunar.NewParser(prog)
	parser.MarkTransientPackage("github.com/eandre/sbm/wow")
	for _, pkg := range prog.InitialPackages() {
		if parser.IsTransientPkg(pkg.Pkg) {
			continue
		}

		pkgPath := filepath.Join(target, pkg.Pkg.Path())
		if err := os.MkdirAll(pkgPath, 0755); err != nil {
			log.Fatalln("Could not create directory:", err)
		}
		for _, f := range pkg.Files {
			fname := filepath.Base(prog.Fset.File(f.Pos()).Name())
			fname = strings.TrimSuffix(fname, ".go")
			out, err := os.Create(filepath.Join(pkgPath, fname + ".lua"))
			if err != nil {
				log.Fatalln("Could not create file:", err)
			}
			defer out.Close()

			if err := parser.ParseNode(out, f); err != nil {
				log.Fatalf("Could not parse file %s: %v", fname, err)
			}
		}
	}
}
