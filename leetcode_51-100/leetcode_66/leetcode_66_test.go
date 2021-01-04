package leetcode_66

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	nums1 := []int{1, 2, 3}
	expected1 := []int{1, 2, 4}
	res1 := plusOne(nums1)
	assert.Equal(t, expected1, res1)

	nums2 := []int{4, 3, 2, 1}
	expected2 := []int{4, 3, 2, 2}
	res2 := plusOne(nums2)
	assert.Equal(t, expected2, res2)

	nums3 := []int{9}
	expected3 := []int{1, 0}
	res3 := plusOne(nums3)
	assert.Equal(t, expected3, res3)
}
