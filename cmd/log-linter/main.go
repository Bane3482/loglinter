package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/go-log-message-linter/pkg/log-linter/internal/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
