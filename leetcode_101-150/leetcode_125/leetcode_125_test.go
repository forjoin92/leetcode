package leetcode_125

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	s1 := "A man, a plan, a canal: Panama"
	expected1 := true
	res1 := isPalindrome(s1)
	assert.Equal(t, expected1, res1)

	s2 := "race a car"
	expected2 := false
	res2 := isPalindrome(s2)
	assert.Equal(t, expected2, res2)
}
