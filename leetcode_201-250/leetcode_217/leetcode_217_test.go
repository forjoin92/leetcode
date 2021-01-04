package leetcode_217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate(t *testing.T) {
	nums1 := []int{1, 2, 3, 1}
	expected1 := true
	contain1 := containsDuplicate(nums1)
	assert.Equal(t, expected1, contain1)

	nums2 := []int{1, 2, 3, 4}
	expected2 := false
	contain2 := containsDuplicate(nums2)
	assert.Equal(t, expected2, contain2)

	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	expected3 := true
	contain3 := containsDuplicate(nums3)
	assert.Equal(t, expected3, contain3)
}
