package leetcode_40

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum21(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	expected := [][]int{
		{1, 7},
		{1, 2, 5},
		{2, 6},
		{1, 1, 6},
	}
	actual := combinationSum2(candidates, target)
	assert.Equal(t, expected, actual)
}

func Test_combinationSum22(t *testing.T) {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	expected := [][]int{
		{1, 2, 2},
		{5},
	}
	actual := combinationSum2(candidates, target)
	assert.Equal(t, expected, actual)
}
