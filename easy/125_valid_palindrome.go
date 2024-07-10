package easy

import (
	"strings"
	"unicode"
)

// IsPalindromeStr
// Time O(N)
// Memory O(1)
func IsPalindromeStr(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if isLetterOrDigit(s[start]) && isLetterOrDigit(s[end]) {
			if strings.ToLower(string(s[start])) == strings.ToLower(string(s[end])) {
				start += 1
				end -= 1
				continue
			}
		}

		if !isLetterOrDigit(s[start]) {
			start += 1
			continue
		}

		if !isLetterOrDigit(s[end]) {
			end -= 1
			continue
		}
		return false
	}
	return true
}

func isLetterOrDigit(b byte) bool {
	return unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b))
}
