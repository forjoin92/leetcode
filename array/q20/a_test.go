package q20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert1(t *testing.T) {
	intervals := [][]int{
		{1, 3},
		{6, 9},
	}
	newInterval := []int{2, 5}
	expected := [][]int{
		{1, 5},
		{6, 9},
	}
	actual := insert(intervals, newInterval)
	assert.Equal(t, expected, actual)
}

func Test_insert2(t *testing.T) {
	intervals := [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	newInterval := []int{4, 8}
	expected := [][]int{
		{1, 2},
		{3, 10},
		{12, 16},
	}
	actual := insert(intervals, newInterval)
	assert.Equal(t, expected, actual)
}

func Test_insert3(t *testing.T) {
	intervals := [][]int{
		{1, 5},
	}
	newInterval := []int{0, 0}
	expected := [][]int{
		{0, 0},
		{1, 5},
	}
	actual := insert(intervals, newInterval)
	assert.Equal(t, expected, actual)
}
