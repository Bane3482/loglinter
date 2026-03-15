package analyzer

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"
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

func isEnglishLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func isSmallLetter(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

func isSensitiveData(name string) bool {
	name = strings.ToLower(name)
	return strings.Contains(name, "password") ||
		strings.Contains(name, "token") ||
		strings.Contains(name, "api") ||
		strings.Contains(name, "ip") ||
		strings.Contains(name, "ssh") ||
		strings.Contains(name, "cvv") ||
		strings.Contains(name, "pin")
}

func isCorrectMessage(expr ast.Expr) bool {
	switch n := expr.(type) {
	case *ast.BinaryExpr:
		{
			if n.Op == token.ADD {
				return isCorrectMessage(n.X) && isCorrectMessage(n.Y)
			}
			return true
		}
	case *ast.BasicLit:
		{
			if n.Kind == token.STRING {
				runes := ([]rune)(n.Value)
				for i := 1; i+1 < len(runes); i++ {
					if !isEnglishLetter(runes[i]) && !unicode.IsSpace(runes[i]) {
						return false
					} else if i == 1 && !isSmallLetter(runes[i]) {
						return false
					}
				}
			}
			return true
		}
	case *ast.Ident:
		{
			if isSensitiveData(n.Name) {
				return false
			}
			return true
		}
	}
	return true
}
