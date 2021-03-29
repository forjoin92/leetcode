package leetcode_190

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBits1(t *testing.T) {
	var num uint32 = 0b00000010100101000001111010011100
	var expected uint32 = 0b00111001011110000010100101000000
	actual := reverseBits(num)
	assert.Equal(t, expected, actual)
}

func Test_reverseBits2(t *testing.T) {
	var num uint32 = 0b11111111111111111111111111111101
	var expected uint32 = 0b10111111111111111111111111111111
	actual := reverseBits(num)
	assert.Equal(t, expected, actual)
}
