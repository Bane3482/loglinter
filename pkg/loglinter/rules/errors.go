package rules

var errorType = map[int]string{
	1: "message contains not english text or special symbols",
	2: "message contains potencial sensitive data",
	3: "message starts not from small letter",
}
