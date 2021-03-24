package leetcode_456

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_find132pattern1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := false
	actual := find132pattern(nums)
	assert.Equal(t, expected, actual)
}

func Test_find132pattern2(t *testing.T) {
	nums := []int{3, 1, 4, 2}
	expected := true
	actual := find132pattern(nums)
	assert.Equal(t, expected, actual)
}

func Test_find132pattern3(t *testing.T) {
	nums := []int{-1, 3, 2, 0}
	expected := true
	actual := find132pattern(nums)
	assert.Equal(t, expected, actual)
}

func Test_find132pattern4(t *testing.T) {
	nums := []int{-2, 1, -1}
	expected := true
	actual := find132pattern(nums)
	assert.Equal(t, expected, actual)
}

func Test_find132pattern5(t *testing.T) {
	nums := []int{3, 5, 0, 3, 4}
	expected := true
	actual := find132pattern(nums)
	assert.Equal(t, expected, actual)
}
