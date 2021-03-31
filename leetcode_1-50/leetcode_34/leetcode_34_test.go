package leetcode_34

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}
	actual := searchRange(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_searchRange2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	expected := []int{-1, -1}
	actual := searchRange(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_searchRange3(t *testing.T) {
	nums := []int{}
	target := 0
	expected := []int{-1, -1}
	actual := searchRange(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_searchRange4(t *testing.T) {
	nums := []int{2, 2}
	target := 2
	expected := []int{0, 1}
	actual := searchRange(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_searchRange5(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5, 9}
	target := 3
	expected := []int{2, 5}
	actual := searchRange(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_searchRange6(t *testing.T) {
	nums := []int{0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 4, 4, 5, 5, 5, 5, 6, 7}
	target := 0
	expected := []int{0, 4}
	actual := searchRange(nums, target)
	assert.Equal(t, expected, actual)
}
