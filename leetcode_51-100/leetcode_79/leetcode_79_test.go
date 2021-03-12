package leetcode_79

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_exist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	assert.Equal(t, true, exist(board, word))
	word = "SEE"
	assert.Equal(t, true, exist(board, word))
	word = "ABCB"
	assert.Equal(t, false, exist(board, word))
}
