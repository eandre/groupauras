package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/eandre/lunar"
	"golang.org/x/tools/go/loader"
)

func main() {
	var (
		output = flag.String("output", "./output", "Output directory")
		strip  = flag.String("strip", "", "Path prefix to strip")
	)
	flag.Parse()
	*strip = path.Clean(*strip)

	var conf loader.Config
	_, err := conf.FromArgs(flag.Args(), false)
	if err != nil {
		log.Fatalln("Could not load packages:", err)
	}

	prog, err := conf.Load()
	if err != nil {
		log.Fatalln("Could not parse packages:", err)
	}

	if err := os.MkdirAll(*output, 0755); err != nil {
		log.Fatalln("Could not create directory:", err)
	}

	parser := lunar.NewParser(prog)
	pkgName := filepath.Base(*output)
	var filenames []string

	var infos []*loader.PackageInfo
	for _, i := range prog.AllPackages {
		infos = append(infos, i)
	}
	pkgs := sortPkgs(infos)
	for _, pkg := range pkgs {
		if parser.IsTransientPkg(pkg.Pkg) {
			continue
		}

		path := pkg.Pkg.Path()
		if *strip != "" && strings.HasPrefix(path, *strip) {
			path = path[len(*strip):]
			// If the path starts with a slash, strip it
			if strings.HasPrefix(path, "/") {
				path = path[1:]
			}
		}

		pkgPath := filepath.Join(*output, path)
		if err := os.MkdirAll(pkgPath, 0755); err != nil {
			log.Fatalln("Could not create directory:", err)
		}

		for _, f := range pkg.Files {
			basePath := strings.TrimSuffix(prog.Fset.File(f.Pos()).Name(), ".go")
			fname := filepath.Base(basePath) + ".lua"
			filenames = append(filenames, filepath.Join(path, fname))

			out, err := os.Create(filepath.Join(pkgPath, fname))
			if err != nil {
				log.Fatalln("Could not create file:", err)
			}
			defer out.Close()

			// Check if a lua file already exists
			fi, err := os.Stat(basePath + ".lua")
			if err == nil && !fi.IsDir() {
				bytes, err := ioutil.ReadFile(basePath + ".lua")
				if err != nil {
					log.Fatalf("Could not read file %s: %v", fname, err)
				}
				if _, err := out.Write(bytes); err != nil {
					log.Fatalf("Could not write file %s: %v", fname, err)
				}
				continue
			}

			if err := parser.ParseNode(out, f); err != nil {
				log.Fatalf("Could not parse file %s: %v", fname, err)
			}
		}
	}

	libsMetaPath := filepath.Join(filepath.Dir(os.Args[0]), "..", "..", "lualibs", "libs.json")
	tocPaths, err := CopyLibs(libsMetaPath, *output)
	if err != nil {
		log.Fatalf("Could not copy lualibs: %v", err)
	}
	filenames = append(tocPaths, filenames...)

	if err := writePackage(prog, *output, pkgName, filenames); err != nil {
		log.Fatalf("Could not write package: %v", err)
	}
}

func writePackage(prog *loader.Program, root, pkgName string, filenames []string) error {
	prelude, err := os.Create(filepath.Join(root, "_prelude.lua"))
	if err != nil {
		return err
	}
	defer prelude.Close()
	if _, err := lunar.WriteBuiltins(prelude); err != nil {
		return err
	}

	postlude, err := os.Create(filepath.Join(root, "_postlude.lua"))
	if err != nil {
		return err
	}
	if _, err := postlude.Write([]byte("lunar_go_builtins.run_inits()")); err != nil {
		return err
	}

	toc, err := os.Create(filepath.Join(root, pkgName+".toc"))
	if err != nil {
		return err
	}
	defer toc.Close()
	var tw tocWriter
	tw.AddKey("Interface", "60200")
	tw.AddKey("Title", "Group Auras")
	tw.AddFile("_prelude.lua")
	for _, fn := range filenames {
		tw.AddFile(fn)
	}
	tw.AddFile("_postlude.lua")
	return tw.Output(toc)
}
