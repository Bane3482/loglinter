package rules

import "unicode"

func isEnglishLetter(s string) bool {
	runes := ([]rune)(s)

	for i := 1; i+1 < len(runes); i++ {
		if unicode.IsLetter(runes[i]) {
			if !unicode.Is(unicode.Latin, runes[i]) {
				return false
			}
		}
	}
	return true
}
