package leetcode_125

import "strings"

func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	head, tail := 0, len(s)-1
	for head < tail {
		if !isValid(lower[head]) {
			head++
			continue
		}
		if !isValid(lower[tail]) {
			tail--
			continue
		}
		if lower[head] != lower[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}

func isValid(c uint8) bool {
	if c < '0' || (c > '9' && c < 'a') || c > 'z' {
		return false
	}
	return true
}
