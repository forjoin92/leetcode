package leetcode_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	expected1 := []int{0,1}
	res1 := twoSum(nums1, target1)
	assert.Subset(t, expected1, res1)

	nums2 := []int{4, 4}
	target2 := 8
	expected2 := []int{0, 1}
	res2 := twoSum(nums2, target2)
	assert.Equal(t, expected2, res2)
}
