package leetcode_70

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	n := 2
	expected := 2
	res := climbStairs(n)
	assert.Equal(t, expected, res)
}

func Test_climbStairs1(t *testing.T) {
	n := 3
	expected := 3
	res := climbStairs(n)
	assert.Equal(t, expected, res)
}
