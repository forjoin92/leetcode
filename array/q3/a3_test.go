package q3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	expected1 := []int{5, 6, 7, 1, 2, 3, 4}
	k1 := 3
	rotate(nums1, k1)
	assert.Equal(t, expected1, nums1)

	nums2 := []int{-1, -100, 3, 99}
	expected2 := []int{3, 99, -1, -100}
	k2 := 2
	rotate(nums2, k2)
	assert.Equal(t, expected2, nums2)

	nums3 := []int{1, 2}
	expected3 := []int{1, 2}
	k3 := 2
	rotate(nums3, k3)
	assert.Equal(t, expected3, nums3)

	nums4 := []int{1, 2}
	expected4 := []int{2, 1}
	k4 := 3
	rotate(nums4, k4)
	assert.Equal(t, expected4, nums4)
}
