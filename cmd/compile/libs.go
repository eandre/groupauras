package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type LibsMetadata struct {
	Libs []string `json:"libs"`
}

func CopyLibs(metadataPath, output string) (tocPaths []string, err error) {
	metadataBytes, err := ioutil.ReadFile(metadataPath)
	if err != nil {
		return nil, err
	}
	var metadata LibsMetadata
	if err := json.Unmarshal(metadataBytes, &metadata); err != nil {
		return nil, err
	}

	baseDir := filepath.Dir(metadataPath)
	for _, lib := range metadata.Libs {
		srcPath := filepath.Join(baseDir, lib)
		dstPath := filepath.Join(output, "lualibs", lib)
		paths, err := copyDir(srcPath, dstPath)
		if err != nil {
			return nil, err
		}
		for i, path := range paths {
			paths[i] = strings.TrimPrefix(path[len(output)+1:], "/")
		}
		tocPaths = append(tocPaths, paths...)
	}
	return tocPaths, nil
}

func copyDir(src, dst string) (paths []string, err error) {
	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".lua") {
			return nil
		}
		in, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dst, path[len(src):])
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}
		out, err := os.Create(targetPath)
		if err != nil {
			return err
		}

		if _, err := out.Write(in); err != nil {
			return err
		}
		paths = append(paths, targetPath)
		return nil
	})
	return paths, err
}