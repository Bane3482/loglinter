package rules

import (
	"go/ast"
	"go/token"
)

func IsSlogMethod(name string) bool {
	switch name {
	case "Debug", "Info", "Warn", "Error":
		return true
	default:
		return false
	}
}

func IsZapMethod(name string) bool {
	switch name {
	case "Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal":
		return true
	default:
		return false
	}
}

func IsLoggerType(expr ast.Expr) string {
	return "nil"
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
