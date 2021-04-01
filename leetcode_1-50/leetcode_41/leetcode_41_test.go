package leetcode_41

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstMissingPositive1(t *testing.T) {
	nums := []int{1, 2, 0}
	expected := 3
	actual := firstMissingPositive(nums)
	assert.Equal(t, expected, actual)
}

func Test_firstMissingPositive2(t *testing.T) {
	nums := []int{3, 4, -1, 1}
	expected := 2
	actual := firstMissingPositive(nums)
	assert.Equal(t, expected, actual)
}

func Test_firstMissingPositive3(t *testing.T) {
	nums := []int{7, 8, 9, 11, 12}
	expected := 1
	actual := firstMissingPositive(nums)
	assert.Equal(t, expected, actual)
}

func Test_firstMissingPositive4(t *testing.T) {
	nums := []int{1, 1}
	expected := 2
	actual := firstMissingPositive(nums)
	assert.Equal(t, expected, actual)
}
