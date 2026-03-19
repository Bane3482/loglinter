package loglinter

import (
	"go/ast"

	"github.com/Bane3482/loglinter/pkg/loglinter/rules"
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
		if rules.IsLoggerType(pass, selectorExpr.X) && rules.IsLogMethod(selectorExpr.Sel.Name) {
			for _, arg := range callExpr.Args {
				if _, ok := rules.IsCorrectMessage(arg); ok != 0 {
					pass.Reportf(node.Pos(), "message starts not from small letter")
				}
			}
		}
	})
	return nil, nil
}
