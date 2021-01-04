package leetcode_136

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	nums1 := []int{2, 2, 1}
	expected1 := 1
	contain1 := singleNumber(nums1)
	assert.Equal(t, expected1, contain1)

	nums2 := []int{4, 1, 2, 1, 2}
	expected2 := 4
	contain2 := singleNumber(nums2)
	assert.Equal(t, expected2, contain2)

	nums3 := []int{1}
	expected3 := 1
	contain3 := singleNumber(nums3)
	assert.Equal(t, expected3, contain3)
}
