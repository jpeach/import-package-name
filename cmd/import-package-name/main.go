package main

import (
	"flag"

	"github.com/jpeach/import-package-name/pkg/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

var importSpec analyzer.Flag

func init() {
	flag.Var(importSpec, "imports", "comma-separated list of Name=Path import specs")
}

func main() {
	// Unused; just to not crash on -unsafeptr flag from go vet.
	flag.Bool("unsafeptr", false, "")

	singlechecker.Main(analyzer.Analyzer)
}
