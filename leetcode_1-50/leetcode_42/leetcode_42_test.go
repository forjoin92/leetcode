package leetcode_42

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trap1(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6
	actual := trap(height)
	assert.Equal(t, expected, actual)
}

func Test_trap2(t *testing.T) {
	height := []int{4, 2, 0, 3, 2, 5}
	expected := 9
	actual := trap(height)
	assert.Equal(t, expected, actual)
}

func Test_trap3(t *testing.T) {
	height := []int{4, 2, 3}
	expected := 1
	actual := trap(height)
	assert.Equal(t, expected, actual)
}

func Test_trap4(t *testing.T) {
	height := []int{9, 6, 8, 8, 5, 6, 3}
	expected := 3
	actual := trap(height)
	assert.Equal(t, expected, actual)
}
