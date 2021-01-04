package leetcode_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	expected := 2.0000
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expected, actual)
}

func Test_findMedianSortedArrays2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	expected := 2.5
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expected, actual)
}

func Test_findMedianSortedArrays3(t *testing.T) {
	nums1 := []int{0, 0}
	nums2 := []int{0, 0}
	expected := 0.00000
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expected, actual)
}

func Test_findMedianSortedArrays4(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{1}
	expected := 1.00000
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expected, actual)
}

func Test_findMedianSortedArrays5(t *testing.T) {
	nums1 := []int{2}
	nums2 := []int{}
	expected := 2.00000
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expected, actual)
}
