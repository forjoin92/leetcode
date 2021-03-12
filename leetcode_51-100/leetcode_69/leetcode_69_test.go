package leetcode_69

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mySqrt(t *testing.T) {
	nums := 4
	expected := 2
	res := mySqrt(nums)
	assert.Equal(t, expected, res)
}

func Test_mySqrt1(t *testing.T) {
	nums := 8
	expected := 2
	res := mySqrt(nums)
	assert.Equal(t, expected, res)
}
