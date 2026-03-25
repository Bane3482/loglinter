package rules

import (
	"unicode"
)

func isSpecialSymbol(s string) bool {
	runes := ([]rune)(s)

	for i := 0; i < len(runes); i++ {
		if !unicode.IsLetter(runes[i]) &&
			!unicode.IsSpace(runes[i]) &&
			!unicode.IsDigit(runes[i]) {
			return false
		}
	}
	return true
}
