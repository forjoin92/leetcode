package leetcode_39

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum1(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	expected := [][]int{
		{2, 2, 3},
		{7},
	}
	actual := combinationSum(candidates, target)
	assert.Equal(t, expected, actual)
}

func Test_combinationSum2(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8
	expected := [][]int{
		{2, 2, 2, 2},
		{2, 3, 3},
		{3, 5},
	}
	actual := combinationSum(candidates, target)
	assert.Equal(t, expected, actual)
}

func Test_combinationSum3(t *testing.T) {
	candidates := []int{2, 7, 6, 3, 5, 1}
	target := 9
	expected := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 2},
		{1, 1, 1, 1, 1, 1, 3},
		{1, 1, 1, 1, 1, 2, 2},
		{1, 1, 1, 1, 2, 3},
		{1, 1, 1, 1, 5},
		{1, 1, 1, 2, 2, 2},
		{1, 1, 1, 3, 3},
		{1, 1, 1, 6},
		{1, 1, 2, 2, 3},
		{1, 1, 2, 5},
		{1, 1, 7},
		{1, 2, 2, 2, 2},
		{1, 2, 3, 3},
		{1, 2, 6},
		{1, 3, 5},
		{2, 2, 2, 3},
		{2, 2, 5},
		{2, 7},
		{3, 3, 3},
		{3, 6},
	}
	actual := combinationSum(candidates, target)
	assert.Equal(t, expected, actual)
}
