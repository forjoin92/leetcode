package leetcode_7

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	x1 := 123
	expected1 := 321
	res1 := reverse(x1)
	assert.Equal(t, expected1, res1)

	x2 := math.MaxInt32
	expected2 := 0
	res2 := reverse(x2)
	assert.Equal(t, expected2, res2)
}
