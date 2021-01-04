package leetcode_28

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr1(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	expected := 2
	res1 := strStr(haystack, needle)
	assert.Equal(t, expected, res1)
}

func Test_strStr2(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"
	expected := -1
	res1 := strStr(haystack, needle)
	assert.Equal(t, expected, res1)
}
