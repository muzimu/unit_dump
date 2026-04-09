package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

var (
	reTestFile = regexp.MustCompile(`_test\.go$`)
	reTestFunc = regexp.MustCompile(`^Test`)
)

// isTestFile reports whether the given filename is a Go test file.
func isTestFile(name string) bool {
	return reTestFile.MatchString(name)
}

// isTestFunc reports whether the given function name is a Go test function.
func isTestFunc(name string) bool {
	return reTestFunc.MatchString(name)
}

// isDir reports whether path is a directory.
func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// collectFiles walks path and returns all non-directory file paths.
func collectFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// dumpTestFuncs prints the names of all Test* functions found in path.
func dumpTestFuncs(path string) {
	if isDir(path) || !isTestFile(path) {
		return
	}

	fset := token.NewFileSet()
	af, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return
	}

	for _, decl := range af.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if isTestFunc(fd.Name.Name) {
			fmt.Fprintln(os.Stdout, fd.Name.Name)
		}
	}
}
