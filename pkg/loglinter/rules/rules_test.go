package rules

import "testing"

func TestIsEnglishMessage(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty string", "", true},
		{"single letter a", "a", true},
		{"single letter A", "A", true},
		{"single digit", "1", true},
		{"single symbol", "!", true},
		{"two letters", "ab", true},
		{"three letters", "abc", true},
		{"mixed ascii but non-letter in middle", "a1c", true},
		{"middle non-english letter ß", "aßc", false},
		{"middle non-english letter ä", "aäc", false},
		{"middle non-english letter é", "aéc", false},
		{"middle non-english letter ø", "aøc", false},
		{"middle digit", "a1c", true},
		{"middle symbol", "a#c", true},
		{"single non-letter symbol", "@", true},
		{"two non-letter symbols", "@#", true},
		{"middle cyrillic letter", "aяc", false},
		{"middle chinese char", "a你c", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("panic for input %q: %v", tc.input, r)
				}
			}()

			if got := isEnglishMessage(tc.input); got != tc.want {
				t.Errorf("isEnglishLetter(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsSensitiveData(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", false},
		{"no_sensitive", "username", false},
		{"random_text", "some_field_name", false},
		{"password_lower", "password", true},
		{"password_mixed", "UserPasswordField", true},
		{"password_upper", "PASSWORD_HASH", true},
		{"token_lower", "token", true},
		{"token_in_name", "authTokenValue", true},
		{"token_upper", "ACCESS_TOKEN", true},
		{"api_lower", "api", true},
		{"api_in_name", "PublicAPIKey", true},
		{"api_upper", "API_KEY", true},
		{"ip_lower", "ip", true},
		{"ip_in_name", "clientIPAddr", true},
		{"ip_upper", "IP_ADDRESS", true},
		{"ssh_lower", "ssh", true},
		{"ssh_in_name", "sshPrivateKey", true},
		{"ssh_upper", "SSH_KEY", true},
		{"cvv_lower", "cvv", true},
		{"cvv_in_name", "cardCVVCode", true},
		{"cvv_upper", "CARD_CVV", true},
		{"pin_lower", "pin", true},
		{"pin_in_name", "userPINCode", true},
		{"pin_upper", "CARD_PIN", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("panic for input %q: %v", tc.input, r)
				}
			}()

			got := isSensitiveData(tc.input)
			if got != tc.want {
				t.Errorf("isSensitiveData(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsSmallLetter(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", false},
		{"one_rune_lower_a", "a", true},
		{"one_rune_lower_z", "z", true},
		{"one_rune_upper_A", "A", false},
		{"one_rune_digit", "1", false},
		{"one_rune_symbol", "_", false},
		{"multi_lower_start", "abc", true},
		{"multi_upper_start", "Abc", false},
		{"multi_digit_start", "1abc", false},
		{"multi_symbol_start", "_abc", false},
		{"multi_non_ascii_start", "яabc", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("panic for input %q: %v", tc.input, r)
				}
			}()

			got := isSmallLetter(tc.input)
			if got != tc.want {
				t.Errorf("isSmallLetter(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsSpecialSymbol(t *testing.T) {
	tcs := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"only_letters", "AbcXYZ", true},
		{"letters_spaces", "Hello world", true},
		{"letters_digits", "user123", true},
		{"digits_spaces", "123 456", true},
		{"letters_digits_spaces", "User 123 Name", true},
		{"with_punctuation", "hello, world", false},
		{"with_exclamation", "hello!", false},
		{"with_question", "what?", false},
		{"with_underscore", "user_name", false},
		{"with_dash", "user-name", false},
		{"with_plus", "a+b", false},
		{"with_brackets", "func()", false},
		{"with_slash", "path/to/file", false},
		{"with_backslash", "path\\\\to\\\\file\\", false},
		{"with_emoji", "hello🙂", false},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("panic for input %q: %v", tc.input, r)
				}
			}()

			got := isSpecialSymbol(tc.input)
			if got != tc.want {
				t.Errorf("isSpecialSymbol(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
