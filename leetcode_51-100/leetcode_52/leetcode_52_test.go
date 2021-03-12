package leetcode_52

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_totalNQueens1(t *testing.T) {
	n := 4
	expected := 2
	actual := totalNQueens(n)
	assert.Equal(t, expected, actual)
}
func Test_totalNQueens2(t *testing.T) {
	n := 1
	expected := 1
	actual := totalNQueens(n)
	assert.Equal(t, expected, actual)
}
