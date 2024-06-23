package easy

// LongestCommonPrefix
// Time O(N^2)
// Memory O(1)
func LongestCommonPrefix(strs []string) string {
	prefixChars := ""
	mainCharIndex := 0

	if len(strs) == 0 || len(strs[0]) == 0 {
		return prefixChars
	}

	char := strs[mainCharIndex][mainCharIndex]
	for {
		for _, str := range strs {
			if mainCharIndex >= len(str) {
				return prefixChars
			}

			if char == str[mainCharIndex] {
				continue
			} else {
				return prefixChars
			}
		}
		mainCharIndex += 1
		prefixChars += string(char)
		if mainCharIndex >= len(strs[0]) {
			return prefixChars
		}
		char = strs[0][mainCharIndex]
	}
}
