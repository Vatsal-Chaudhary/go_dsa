package string

import (
	"unicode"
)

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if !isAlphaNumeric(s[l]) {
			l++
			continue
		}
		if !isAlphaNumeric(s[r]) {
			r--
			continue
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}

		l++
		r--
	}

	return true
}

func isAlphaNumeric(ch byte) bool {
	return unicode.IsLetter(rune(ch)) ||
		unicode.IsDigit(rune(ch))
}
