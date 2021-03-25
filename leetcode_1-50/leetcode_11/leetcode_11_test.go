package leetcode_11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea1(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	actual := maxArea(height)
	assert.Equal(t, expected, actual)
}

func Test_maxArea2(t *testing.T) {
	height := []int{1, 1}
	expected := 1
	actual := maxArea(height)
	assert.Equal(t, expected, actual)
}

func Test_maxArea3(t *testing.T) {
	height := []int{4, 3, 2, 1, 4}
	expected := 16
	actual := maxArea(height)
	assert.Equal(t, expected, actual)
}

func Test_maxArea4(t *testing.T) {
	height := []int{1, 2, 1}
	expected := 2
	actual := maxArea(height)
	assert.Equal(t, expected, actual)
}
