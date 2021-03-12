package leetcode_51

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveNQueens1(t *testing.T) {
	n := 4
	expected := [][]string{
		{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."},
	}
	actual := solveNQueens(n)
	assert.Equal(t, expected, actual)
}
func Test_solveNQueens2(t *testing.T) {
	n := 1
	expected := [][]string{
		{"Q"},
	}
	actual := solveNQueens(n)
	assert.Equal(t, expected, actual)
}
