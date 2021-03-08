package leetcode_239

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSlidingWindow1(t *testing.T) {
	nums := []int{1}
	k := 1
	expected := []int{1}

	actual := maxSlidingWindow(nums, k)
	assert.Equal(t, expected, actual)
}

func Test_maxSlidingWindow2(t *testing.T) {
	nums := []int{1, -1}
	k := 1
	expected := []int{1, -1}
	actual := maxSlidingWindow(nums, k)
	assert.Equal(t, expected, actual)
}

func Test_maxSlidingWindow3(t *testing.T) {
	nums := []int{9, 11}
	k := 2
	expected := []int{11}
	actual := maxSlidingWindow(nums, k)
	assert.Equal(t, expected, actual)
}

func Test_maxSlidingWindow4(t *testing.T) {
	nums := []int{4, -2}
	k := 2
	expected := []int{4}
	actual := maxSlidingWindow(nums, k)
	assert.Equal(t, expected, actual)
}

func Test_maxSlidingWindow5(t *testing.T) {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3
	expected := []int{3, 3, 2, 5}
	actual := maxSlidingWindow(nums, k)
	assert.Equal(t, expected, actual)
}
