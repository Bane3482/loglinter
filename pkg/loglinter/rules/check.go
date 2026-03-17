package rules

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func IsLoggerType(pass *analysis.Pass, expr ast.Expr) bool {

	if ident, ok := expr.(*ast.Ident); ok {
		obj := pass.TypesInfo.ObjectOf(ident)
		if obj == nil {
			return false
		}
		if pkg, ok := obj.(*types.PkgName); ok {
			return pkg.Imported().Path() == "log/slog"
		}
	}

	t := pass.TypesInfo.TypeOf(expr)

	if t == nil {
		return false
	}

	for {
		switch tt := t.(type) {
		case *types.Pointer:
			t = tt.Elem()
			continue
		default:
			return strings.Contains(tt.Underlying().String(), "go.uber.org/zap")
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
				if !isEnglishLetter(n.Value) {
					return n.Value, 1
				} else if !isSmallLetter(n.Value) {
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
			return "nil", 0
		}
	default:
		return "nil", 0
	}
}
