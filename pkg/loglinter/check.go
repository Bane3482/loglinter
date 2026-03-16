package loglinter

import (
	"go/ast"
	"go/token"
)

func isSlogMethod(name string) bool {
	switch name {
	case "Debug", "Info", "Warn", "Error":
		return true
	default:
		return false
	}
}

func isZapMethod(name string) bool {
	switch name {
	case "Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal":
		return true
	default:
		return false
	}
}

func isLogMethod(name string) bool {
	switch name {
	case "Print", "Printf", "Println", "Fatal", "Fatalf", "Fatalln", "Panic", "Panicf", "Panicln":
		return true
	default:
		return false
	}
}

func isCorrectMessage(expr ast.Expr) (string, bool) {
	switch n := expr.(type) {
	case *ast.BinaryExpr:
		{
			if n.Op == token.ADD {
				first, ok1 := isCorrectMessage(n.X)
				second, ok2 := isCorrectMessage(n.Y)
				if !ok1 {
					return first, ok1
				} else if !ok2 {
					return second, ok2
				}
			}
			return "nil", true
		}
	case *ast.BasicLit:
		{
			if n.Kind == token.STRING {
				if !isEnglishLetter(n.Value) || !isSmallLetter(n.Value) {
					return n.Value, false
				}
			}
			return "nil", true
		}
	case *ast.Ident:
		{
			if isSensitiveData(n.Name) {
				return n.Name, false
			}
		}
	}
	return "nil", true
}
