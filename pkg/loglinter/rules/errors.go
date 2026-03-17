package rules

var errorType = []string{
	"message contains not english text",
	"message contains potencial sensitive data",
	"message starts not from small letter",
	"message contains special symbols",
}

func ErrorType(code int) string {
	return errorType[code]
}
