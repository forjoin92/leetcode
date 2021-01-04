package leetcode_283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	nums1 := []int{1, 2, 3}
	expected1 := []int{1, 2, 3}
	moveZeroes(nums1)
	assert.Equal(t, expected1, nums1)

	nums2 := []int{0, 1, 0, 3, 12}
	expected2 := []int{1, 3, 12, 0, 0}
	moveZeroes(nums2)
	assert.Equal(t, expected2, nums2)
}
