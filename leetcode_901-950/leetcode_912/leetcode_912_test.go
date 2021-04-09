package leetcode_912

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortArray1(t *testing.T) {
	nums := []int{5, 2, 3, 1}
	expected := []int{1, 2, 3, 5}
	actual := sortArray(nums)
	assert.Equal(t, expected, actual)
}

func Test_sortArray2(t *testing.T) {
	nums := []int{5, 1, 1, 2, 0, 0}
	expected := []int{0, 0, 1, 1, 2, 5}
	actual := sortArray(nums)
	assert.Equal(t, expected, actual)
}
