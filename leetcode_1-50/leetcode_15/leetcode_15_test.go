package leetcode_15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	expected1 := [][]int{
		{-1, 0, 1},
		{-1, -1, 2},
	}
	actual1 := threeSum(nums1)
	assert.Equal(t, expected1, actual1)
}
