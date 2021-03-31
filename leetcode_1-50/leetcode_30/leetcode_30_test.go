package leetcode_30

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findSubstring1(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	expected := []int{0, 9}
	res := findSubstring(s, words)
	assert.Equal(t, expected, res)
}

func Test_findSubstring2(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "word"}
	var expected []int
	res := findSubstring(s, words)
	assert.Equal(t, expected, res)
}

func Test_findSubstring3(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	expected := []int{8}
	res := findSubstring(s, words)
	assert.Equal(t, expected, res)
}

func Test_findSubstring4(t *testing.T) {
	s := "mississippi"
	words := []string{"mississippis"}
	expected := []int{}
	res := findSubstring(s, words)
	assert.Equal(t, expected, res)
}
