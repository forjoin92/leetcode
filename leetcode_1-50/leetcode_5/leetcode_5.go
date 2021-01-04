package leetcode_5

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var res string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) && len(res) < j-i {
				res = s[i:j]
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for ; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
