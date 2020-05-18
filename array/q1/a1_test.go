package q1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates1(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := []int{0, 1, 2, 3, 4}
	length := removeDuplicates(nums)
	assert.Equal(t, len(expected), length)
	assert.Equal(t, expected, nums[:length])
}

func Test_removeDuplicates2(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4}
	expected := []int{0, 1, 2, 3, 4}
	length := removeDuplicates(nums)
	assert.Equal(t, len(expected), length)
	assert.Equal(t, expected, nums[:length])
}

func Test_removeDuplicates3(t *testing.T) {
	nums := []int{0}
	expected := []int{0}
	length := removeDuplicates(nums)
	assert.Equal(t, len(expected), length)
	assert.Equal(t, expected, nums[:length])
}
