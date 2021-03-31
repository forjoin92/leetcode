package leetcode_977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedSquares1(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	expected := []int{0, 1, 9, 16, 100}
	actual := sortedSquares(nums)
	assert.Equal(t, expected, actual)
}

func Test_sortedSquares2(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}
	expected := []int{4, 9, 9, 49, 121}
	actual := sortedSquares(nums)
	assert.Equal(t, expected, actual)
}
