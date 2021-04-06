package leetcode_80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates1(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	expected := 5
	actual := removeDuplicates(nums)
	assert.Equal(t, expected, actual)
}

func Test_removeDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	expected := 7
	actual := removeDuplicates(nums)
	assert.Equal(t, expected, actual)
}

func Test_removeDuplicates3(t *testing.T) {
	nums := []int{1}
	expected := 1
	actual := removeDuplicates(nums)
	assert.Equal(t, expected, actual)
}

func Test_removeDuplicates4(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 3}
	expected := 6
	actual := removeDuplicates(nums)
	assert.Equal(t, expected, actual)
}
