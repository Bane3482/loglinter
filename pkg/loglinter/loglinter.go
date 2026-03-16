package loglinter

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "loglinter",
	Doc:      "checks for bad patterns of loggers messages",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		callExpr := node.(*ast.CallExpr)
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		if isLogMethod(selectorExpr.Sel.Name) ||
			isSlogMethod(selectorExpr.Sel.Name) ||
			isZapMethod(selectorExpr.Sel.Name) {
			for _, arg := range callExpr.Args {
				if msg, ok := isCorrectMessage(arg); ok != 0 {
					pass.Reportf(node.Pos(), "Incorrect message format, %s: %s", errorType[ok], msg)
				}
			}
		}
	})
	return nil, nil
}
