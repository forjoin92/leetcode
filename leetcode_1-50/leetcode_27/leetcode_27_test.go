package leetcode_27

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expectedNums := []int{2, 2}
	expected := 2
	actual := removeElement(nums, val)
	assert.Equal(t, expectedNums, nums[:expected])
	assert.Equal(t, expected, actual)
}

func Test_removeElement2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	expectedNums := []int{0, 1, 3, 0, 4}
	expected := 5
	actual := removeElement(nums, val)
	assert.Equal(t, expectedNums, nums[:expected])
	assert.Equal(t, expected, actual)
}
