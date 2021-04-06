package leetcode_45

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump1(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	expected := 2
	actual := jump(nums)
	assert.Equal(t, expected, actual)
}

func Test_jump2(t *testing.T) {
	nums := []int{2, 3, 0, 1, 4}
	expected := 2
	actual := jump(nums)
	assert.Equal(t, expected, actual)
}

func Test_jump3(t *testing.T) {
	nums := []int{0}
	expected := 0
	actual := jump(nums)
	assert.Equal(t, expected, actual)
}

func Test_jump4(t *testing.T) {
	nums := []int{3, 2, 1}
	expected := 1
	actual := jump(nums)
	assert.Equal(t, expected, actual)
}

func Test_jump5(t *testing.T) {
	nums := []int{4, 1, 1, 3, 1, 1, 1}
	expected := 2
	actual := jump(nums)
	assert.Equal(t, expected, actual)
}
