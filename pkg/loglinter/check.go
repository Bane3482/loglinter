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

func isCorrectMessage(expr ast.Expr) (string, int) {
	switch n := expr.(type) {
	case *ast.BinaryExpr:
		{
			if n.Op == token.ADD {
				first, ok1 := isCorrectMessage(n.X)
				second, ok2 := isCorrectMessage(n.Y)
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
		}
	}
	return "nil", 0
}
