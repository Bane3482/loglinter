package loglinter

import "strings"

func isSensitiveData(name string) bool {
	check := func(name string) bool {
		return strings.Contains(name, "password") ||
			strings.Contains(name, "token") ||
			strings.Contains(name, "api") ||
			strings.Contains(name, "ip") ||
			strings.Contains(name, "ssh") ||
			strings.Contains(name, "cvv") ||
			strings.Contains(name, "pin")
	}

	if check(strings.ToLower(name)) {
		return false
	}
	return true
}
