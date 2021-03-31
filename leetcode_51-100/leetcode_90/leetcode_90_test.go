package leetcode_90

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetsWithDup1(t *testing.T) {
	nums := []int{1, 2, 2}
	expected := [][]int{
		{1},
		{1, 2},
		{1, 2, 2},
		{2},
		{2, 2},
		{},
	}
	actual := subsetsWithDup(nums)
	assert.Equal(t, expected, actual)
}

func Test_subsetsWithDup2(t *testing.T) {
	nums := []int{0}
	expected := [][]int{
		{0},
		{},
	}
	actual := subsetsWithDup(nums)
	assert.Equal(t, expected, actual)
}

func Test_subsetsWithDup3(t *testing.T) {
	nums := []int{1, 1}
	expected := [][]int{
		{1},
		{1, 1},
		{},
	}
	actual := subsetsWithDup(nums)
	assert.Equal(t, expected, actual)
}

func Test_subsetsWithDup4(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 3},
		{2},
		{2, 3},
		{3},
		{},
	}
	actual := subsetsWithDup(nums)
	assert.Equal(t, expected, actual)
}

func Test_subsetsWithDup5(t *testing.T) {
	nums := []int{5, 5, 5, 5, 5}
	expected := [][]int{
		{5},
		{5, 5},
		{5, 5, 5},
		{5, 5, 5, 5},
		{5, 5, 5, 5, 5},
		{},
	}
	actual := subsetsWithDup(nums)
	assert.Equal(t, expected, actual)
}

func Test_subsetsWithDup6(t *testing.T) {
	nums := []int{4, 4, 4, 1, 4}
	expected := [][]int{
		{1},
		{1, 4},
		{1, 4, 4},
		{1, 4, 4, 4},
		{1, 4, 4, 4, 4},
		{4},
		{4, 4},
		{4, 4, 4},
		{4, 4, 4, 4},
		{},
	}
	actual := subsetsWithDup(nums)
	assert.Equal(t, expected, actual)
}
