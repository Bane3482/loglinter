package rules

import "unicode"

func isEnglishLetter(s string) bool {
	check := func(r rune) bool {
		return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
	}

	runes := ([]rune)(s)

	for i := 1; i+1 < len(runes); i++ {
		if !check(runes[i]) && !unicode.IsSpace(runes[i]) {
			return false
		}
	}
	return true
}
