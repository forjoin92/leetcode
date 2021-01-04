package leetcode_29

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divide1(t *testing.T) {
	dividend := 10
	divisor := 3
	expected := 3
	res := divide(dividend, divisor)
	assert.Equal(t, expected, res)
}

func Test_divide2(t *testing.T) {
	dividend := 7
	divisor := -3
	expected := -2
	res := divide(dividend, divisor)
	assert.Equal(t, expected, res)
}

func Test_divide3(t *testing.T) {
	dividend := -2147483648
	divisor := -1
	expected := 2147483647
	res := divide(dividend, divisor)
	assert.Equal(t, expected, res)
}
