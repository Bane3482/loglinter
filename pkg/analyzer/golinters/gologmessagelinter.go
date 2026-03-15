package golinters

import (
	"github.com/Bane3482/go-log-message-linter/pkg/analyzer"
	"github.com/golangci/golangci-lint/pkg/goanalysis"
	"golang.org/x/tools/go/analysis"
)

func NewGoPrintfFuncName() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"loglinter",
		"checks for bad patterns of loggers messages",
		[]*analysis.Analyzer{analyzer.Analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
