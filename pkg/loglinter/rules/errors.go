package rules

var errorType = []string{
	"message contains not english text or special symbols",
	"message contains potencial sensitive data",
	"message starts not from small letter",
}

func ErrorType(code int) string {
	return errorType[code]
}
