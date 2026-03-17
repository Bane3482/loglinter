package rules

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"slices"

	"golang.org/x/tools/go/analysis"
)

var loggerNames = []string{"log/slog", "go.uber.org/zap"}

func IsLoggerType(pass *analysis.Pass, tv types.TypeAndValue) bool {
	t := tv.Type
	for {
		switch tt := t.Underlying().(type) {
		case *types.Named:
			pkg := tt.Obj().Pkg()
			if pkg != nil {
				fmt.Println(pkg.Path())
				if slices.Contains(loggerNames, pkg.Path()) {
					return true
				}
			} else {
				fmt.Println("nulll")
			}
			t = tt.Underlying()
		case *types.Pointer:
			t = tt.Elem()
		default:
			return false
		}
	}
}

func IsLogMethod(name string) bool {
	switch name {
	case "Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal":
		return true
	default:
		return false
	}
}

func IsCorrectMessage(expr ast.Expr) (string, int) {
	switch n := expr.(type) {
	case *ast.BinaryExpr:
		{
			if n.Op == token.ADD {
				first, ok1 := IsCorrectMessage(n.X)
				second, ok2 := IsCorrectMessage(n.Y)
				if ok1 != 0 {
					return first, ok1
				} else if ok2 != 0 {
					return second, ok2
				}
			}
			return "nil", 0
		}
	case *ast.BasicLit:
		{
			if n.Kind == token.STRING {
				if isEnglishLetter(n.Value) {
					return n.Value, 1
				} else if isSmallLetter(n.Value) {
					return n.Value, 3
				}
			}
			return "nil", 0
		}
	case *ast.Ident:
		{
			if isSensitiveData(n.Name) {
				return n.Name, 2
			}
		}
	}
	return "nil", 0
}
