package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/bane3482/go-log-message-linter/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
