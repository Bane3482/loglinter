package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"loglint/internal/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
