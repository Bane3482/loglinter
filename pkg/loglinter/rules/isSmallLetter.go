package rules

func isSmallLetter(s string) bool {
	runes := ([]rune)(s)
	if len(runes) == 0 {
		return false
	}
	return (runes[0] >= 'a' && runes[0] <= 'z')
}
