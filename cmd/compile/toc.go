package main

import (
	"bufio"
	"fmt"
	"go/types"
	"io"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/loader"
)

type tocWriter struct {
	keys  []string
	files []string
}

func (tw *tocWriter) AddKey(key, value string) {
	tw.keys = append(tw.keys, fmt.Sprintf("## %s: %s", key, value))
}

func (tw *tocWriter) AddFile(fn string) {
	tw.files = append(tw.files, fn)
}

func (tw *tocWriter) Output(w io.Writer) error {
	writer := bufio.NewWriter(w)
	for _, key := range tw.keys {
		writer.WriteString(key + "\r\n")
	}
	writer.WriteString("\r\n")
	sep := string([]rune{filepath.Separator})
	for _, fn := range tw.files {
		writer.WriteString(strings.Replace(fn, sep, "/", -1) + "\r\n")
	}
	return writer.Flush()
}

func sortPkgs(infos []*loader.PackageInfo) []*loader.PackageInfo {
	infoToPkg := make(map[*loader.PackageInfo]*types.Package)
	pkgToInfo := make(map[*types.Package]*loader.PackageInfo)
	remaining := make(map[*types.Package]bool)
	for _, info := range infos {
		infoToPkg[info] = info.Pkg
		pkgToInfo[info.Pkg] = info
		remaining[info.Pkg] = true
	}

	sorted := make([]*loader.PackageInfo, 0, len(infos))

	for len(remaining) > 0 {
		// Find a package with no remaining dependencies
	InfoLoop:
		for _, info := range infos {
			pkg := info.Pkg
			if !remaining[pkg] {
				continue
			}

			// Go through the candidate's dependencies to see if any of them
			// are still remaining.
			for _, dep := range pkg.Imports() {
				if remaining[dep] {
					continue InfoLoop
				}
			}

			// No dependencies remaining; add it to the output
			sorted = append(sorted, info)
			delete(remaining, pkg)
		}
	}

	return sorted
}
