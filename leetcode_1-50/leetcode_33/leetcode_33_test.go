package leetcode_33

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search1(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	expected := 4
	actual := search(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_search2(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	expected := -1
	actual := search(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_search3(t *testing.T) {
	nums := []int{1}
	target := 0
	expected := -1
	actual := search(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_search4(t *testing.T) {
	nums := []int{5, 1, 3}
	target := 5
	expected := 0
	actual := search(nums, target)
	assert.Equal(t, expected, actual)
}

func Test_search5(t *testing.T) {
	nums := []int{3, 1}
	target := 1
	expected := 1
	actual := search(nums, target)
	assert.Equal(t, expected, actual)
}
