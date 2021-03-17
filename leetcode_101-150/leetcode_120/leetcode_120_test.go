package leetcode_120

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumTotal1(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	expected := 11
	res := minimumTotal(triangle)
	assert.Equal(t, expected, res)
}

func Test_minimumTotal2(t *testing.T) {
	triangle := [][]int{
		{-10},
	}
	expected := -10
	res := minimumTotal(triangle)
	assert.Equal(t, expected, res)
}
