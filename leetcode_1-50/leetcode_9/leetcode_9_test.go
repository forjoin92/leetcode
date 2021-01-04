package leetcode_9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	x1 := 121
	expected1 := true
	res1 := isPalindrome(x1)
	assert.Equal(t, expected1, res1)

	x2 := -121
	expected2 := false
	res2 := isPalindrome(x2)
	assert.Equal(t, expected2, res2)

	x3 := 10
	expected3 := false
	res3 := isPalindrome(x3)
	assert.Equal(t, expected3, res3)
}
