package leetcode_212

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findWords1(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	expected := []string{"oath", "eat"}
	actual := findWords(board, words)
	assert.Equal(t, expected, actual)
}

func Test_findWords2(t *testing.T) {
	board := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	words := []string{"abcb"}
	var expected []string
	actual := findWords(board, words)
	assert.Equal(t, expected, actual)
}

func Test_findWords3(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'b', 'n'},
		{'o', 't', 'a', 'e'},
		{'a', 'h', 'k', 'r'},
		{'a', 'f', 'l', 'v'},
	}
	words := []string{"oa", "oaa"}
	expected := []string{"oaa", "oa"}
	actual := findWords(board, words)
	assert.Equal(t, expected, actual)
}
