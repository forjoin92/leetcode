package q7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr(t *testing.T) {
	haystack1 := "hello"
	needle1 := "ll"
	expected1 := 2
	res1 := strStr(haystack1, needle1)
	assert.Equal(t, expected1, res1)

	haystack2 := "aaaaa"
	needle2 := "bba"
	expected2 := -1
	res2 := strStr(haystack2, needle2)
	assert.Equal(t, expected2, res2)
}
