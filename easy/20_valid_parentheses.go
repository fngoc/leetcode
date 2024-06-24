package easy

var m = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

// IsValid
// Time O(N)
// Memory O(N)
func IsValid(s string) bool {
	var stack []rune

	if len(s) == 0 || len(s) == 1 {
		return false
	}

	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == m[char] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
