package leetcode_242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram(t *testing.T) {
	s1 := "anagram"
	t1 := "nagaram"
	expected1 := true
	res1 := isAnagram(s1, t1)
	assert.Equal(t, expected1, res1)

	s2 := "rat"
	t2 := "car"
	expected2 := false
	res2 := isAnagram(s2, t2)
	assert.Equal(t, expected2, res2)
}
