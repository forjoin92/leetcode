package q10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	s1 := "abcabcbb"
	expected1 := 3
	res1 := lengthOfLongestSubstring(s1)
	assert.Equal(t, expected1, res1)

	s2 := "bbbbb"
	expected2 := 1
	res2 := lengthOfLongestSubstring(s2)
	assert.Equal(t, expected2, res2)

	s3 := "pwwkew"
	expected3 := 3
	res3 := lengthOfLongestSubstring(s3)
	assert.Equal(t, expected3, res3)

	s4 := "aab"
	expected4 := 2
	res4 := lengthOfLongestSubstring(s4)
	assert.Equal(t, expected4, res4)

	s5 := "abba"
	expected5 := 2
	res5 := lengthOfLongestSubstring(s5)
	assert.Equal(t, expected5, res5)
}
