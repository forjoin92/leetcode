package leetcode_31

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermutation1(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := []int{1, 3, 2}
	nextPermutation(nums)
	assert.Equal(t, expected, nums)
}

func Test_nextPermutation2(t *testing.T) {
	nums := []int{3, 2, 1}
	expected := []int{1, 2, 3}
	nextPermutation(nums)
	assert.Equal(t, expected, nums)
}

func Test_nextPermutation3(t *testing.T) {
	nums := []int{1, 1, 5}
	expected := []int{1, 5, 1}
	nextPermutation(nums)
	assert.Equal(t, expected, nums)
}

func Test_nextPermutation4(t *testing.T) {
	nums := []int{1}
	expected := []int{1}
	nextPermutation(nums)
	assert.Equal(t, expected, nums)
}
