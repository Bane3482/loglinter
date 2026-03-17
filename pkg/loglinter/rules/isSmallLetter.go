package rules

func isSmallLetter(s string) bool {
	runes := ([]rune)(s)
	return (runes[0] >= 'a' && runes[0] <= 'z')
}
