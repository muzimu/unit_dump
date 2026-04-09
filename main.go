// unit_dump scans Go test files under the given path (default: current directory)
// and prints the name of every Test* function, one per line.
//
// Intended to be used with fzf:
//
//	alias gout='go test -v -run $(unit_dump | fzf)'
package main

import (
	"fmt"
	"os"
)

func main() {
	root := "./"
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	files, err := collectFiles(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unit_dump: %v\n", err)
		os.Exit(1)
	}

	for _, f := range files {
		dumpTestFuncs(f)
	}
}
