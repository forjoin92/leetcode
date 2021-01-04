package leetcode_349

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected1 := []int{2}
	actual1 := intersection(nums1, nums2)
	assert.Equal(t, expected1, actual1)

	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	expected2 := []int{9, 4}
	actual2 := intersection(nums3, nums4)
	assert.Equal(t, expected2, actual2)
}
