package rules

func isSmallLetter(s string) bool {
	runes := ([]rune)(s)
	return (runes[1] >= 'a' && runes[1] <= 'z')
}
