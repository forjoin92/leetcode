package leetcode_169

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement1(t *testing.T) {
	nums := []int{3, 2, 3}
	expected := 3
	res := majorityElement(nums)
	assert.Equal(t, expected, res)
}

func Test_majorityElement2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	expected := 2
	res := majorityElement(nums)
	assert.Equal(t, expected, res)
}
