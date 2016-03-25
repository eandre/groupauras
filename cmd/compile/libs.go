package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type LibsMetadata struct {
	Libs []struct {
		Path   string   `json:"path"`
		TOC    []string `json:"toc"`
		Embeds []string `json:"embeds"`
	} `json:"libs"`
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
	var embeds []string
	for _, lib := range metadata.Libs {
		srcPath := filepath.Join(baseDir, lib.Path)
		dstPath := filepath.Join(output, "lualibs", lib.Path)
		paths, err := CopyDir(srcPath, dstPath, nil)
		if err != nil {
			return nil, err
		}

		var tocs []string
		for _, path := range paths {
			relPath := strings.TrimPrefix(path[len(dstPath)+1:], "/")

			// See if it matches TOC entries
			matches := false
			for _, pattern := range lib.TOC {
				matched, err := filepath.Match(pattern, relPath)
				if err != nil {
					return nil, fmt.Errorf("bad pattern %s/%s: %v", lib.Path, pattern, err)
				}
				if matched {
					matches = true
					break
				}
			}
			if matches {
				tocs = append(tocs, strings.TrimPrefix(path[len(output)+1:], "/"))
				continue
			}

			// Otherwise see if it matches Embeds entries
			for _, pattern := range lib.Embeds {
				matched, err := filepath.Match(pattern, relPath)
				if err != nil {
					return nil, fmt.Errorf("bad pattern %s/%s: %v", lib.Path, pattern, err)
				}
				if matched {
					matches = true
					break
				}
			}
			if matches {
				embeds = append(embeds, strings.TrimPrefix(path[len(output)+1:], "/"))
				continue
			}
		}
		tocPaths = append(tocPaths, tocs...)
	}

	// If we have any embeds, write an embed.xml file
	if len(embeds) != 0 {
		embedPath := filepath.Join(output, "embeds.xml")
		if err := writeEmbeds(embedPath, embeds); err != nil {
			return nil, fmt.Errorf("could not write embeds: %v", err)
		}
		tocPaths = append(tocPaths, "embeds.xml")
	}
	return tocPaths, nil
}

func CopyDir(src, dst string, filter func(path string, info os.FileInfo) bool) (paths []string, err error) {
	if filter == nil {
		filter = func(string, os.FileInfo) bool {
			return true
		}
	}

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !filter(path, info) {
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
		defer out.Close()

		if _, err := out.Write(in); err != nil {
			return err
		}
		paths = append(paths, targetPath)
		return nil
	})
	return paths, err
}

func writeEmbeds(path string, entries []string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	w := bufio.NewWriter(out)
	w.WriteString(xmlHeader + "\n")
	for _, entry := range entries {
		if strings.HasSuffix(entry, ".lua") {
			w.WriteString("\t <Script file=\"" + entry + "\"/>\n")
		} else {
			w.WriteString("\t <Include file=\"" + entry + "\"/>\n")
		}
	}
	w.WriteString(xmlFooter)
	return w.Flush()
}

const xmlHeader = `<Ui xmlns="http://www.blizzard.com/wow/ui/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.blizzard.com/wow/ui/
	..\FrameXML\UI.xsd">`
const xmlFooter = `</Ui>`
