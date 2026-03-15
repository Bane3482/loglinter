package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks for bad patterns of loggers messages",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		xExpr, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			return true
		}
		if xExpr.Name == "log" && isLogMethod(selectorExpr.Sel.Name) ||
			xExpr.Name == "slog" && isSlogMethod(selectorExpr.Sel.Name) ||
			xExpr.Name == "zap" && isZapMethod(selectorExpr.Sel.Name) {
			for _, arg := range callExpr.Args {
				if !isCorrectMessage(arg) {
					pass.Reportf(node.Pos(), "Incorrect message format: %s", arg)
				}

			}
		}
		return true
	}
	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
