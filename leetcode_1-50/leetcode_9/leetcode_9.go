package leetcode_9

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	head, tail := 0, len(s)-1
	for ; head < tail; head, tail = head+1, tail-1 {
		if s[head] != s[tail] {
			return false
		}
	}
	return true
}
