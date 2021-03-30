package leetcode_74

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrix1(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	expected := true
	actual := searchMatrix(matrix, target)
	assert.Equal(t, expected, actual)
}

func Test_searchMatrix2(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 13
	expected := false
	actual := searchMatrix(matrix, target)
	assert.Equal(t, expected, actual)
}

func Test_searchMatrix3(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 23
	expected := true
	actual := searchMatrix(matrix, target)
	assert.Equal(t, expected, actual)
}
