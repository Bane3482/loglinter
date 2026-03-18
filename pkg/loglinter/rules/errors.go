package rules

var errorType = []string{
	"Message contains not english text",
	"Message contains potencial sensitive data",
	"Message starts not from small letter",
	"Message contains special symbols",
}

func ErrorType(code int) string {
	return errorType[code]
}
