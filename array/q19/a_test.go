package q19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validMountainArray(t *testing.T) {
	nums1 := []int{2, 1}
	expected1 := false
	actual1 := validMountainArray(nums1)
	assert.Equal(t, expected1, actual1)

	nums2 := []int{3, 5, 5}
	expected2 := false
	actual2 := validMountainArray(nums2)
	assert.Equal(t, expected2, actual2)

	nums3 := []int{0, 3, 2, 1}
	expected3 := true
	actual3 := validMountainArray(nums3)
	assert.Equal(t, expected3, actual3)

	nums4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected4 := false
	actual4 := validMountainArray(nums4)
	assert.Equal(t, expected4, actual4)

	nums5 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	expected5 := false
	actual5 := validMountainArray(nums5)
	assert.Equal(t, expected5, actual5)
}
