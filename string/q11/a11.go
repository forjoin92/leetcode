package q11

func longestPalindrome(s string) string {
	var longest string
	if len(s) == 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) && len(longest) < j-i {
				longest = s[i:j]
			}
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	head, tail := 0, len(s)-1
	for head < tail {
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}
