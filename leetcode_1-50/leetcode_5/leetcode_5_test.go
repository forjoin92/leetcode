package leetcode_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	s1 := "babad"
	expected1 := "bab"
	res1 := longestPalindrome(s1)
	assert.Equal(t, expected1, res1)

	s2 := "cbbd"
	expected2 := "bb"
	res2 := longestPalindrome(s2)
	assert.Equal(t, expected2, res2)

	s3 := "a"
	expected3 := "a"
	res3 := longestPalindrome(s3)
	assert.Equal(t, expected3, res3)

	s4 := "bb"
	expected4 := "bb"
	res4 := longestPalindrome(s4)
	assert.Equal(t, expected4, res4)
}
