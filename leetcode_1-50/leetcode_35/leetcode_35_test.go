package leetcode_35

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert1(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	expected := 2
	actual := searchInsert(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_searchInsert2(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	expected := 1
	actual := searchInsert(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_searchInsert3(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	expected := 4
	actual := searchInsert(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_searchInsert4(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 0
	expected := 0
	actual := searchInsert(nums, target)
	assert.Equal(t, expected, actual)
}
