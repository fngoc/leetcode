package easy

import (
	"strconv"
)

// IsPalindrome
// Time O(N)
// Memory O(1)
func IsPalindrome(x int) bool {
	strInt := strconv.Itoa(x)

	start := 0
	end := len(strInt) - 1
	for start < end {
		if strInt[start] != strInt[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}
