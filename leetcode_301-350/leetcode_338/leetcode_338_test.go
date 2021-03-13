package leetcode_338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countBits1(t *testing.T) {
	num := 2
	expected := []int{0, 1, 1}
	actual := countBits(num)
	assert.Equal(t, expected, actual)
}

func Test_countBits2(t *testing.T) {
	num := 5
	expected := []int{0, 1, 1, 2, 1, 2}
	actual := countBits(num)
	assert.Equal(t, expected, actual)
}
