package leetcode_152

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProduct1(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	expected := 6
	res := maxProduct(nums)
	assert.Equal(t, expected, res)
}

func Test_majorityElement2(t *testing.T) {
	nums := []int{-2, 0, -1}
	expected := 0
	res := maxProduct(nums)
	assert.Equal(t, expected, res)
}
