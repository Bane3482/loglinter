package rules

import "unicode"

func isEnglishMessage(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
				return false
			}
		}
	}
	return true
}
