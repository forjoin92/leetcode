package leetcode_191

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingWeight1(t *testing.T) {
	var num uint32 = 00000000000000000000000000001011
	expected := 3
	actual := hammingWeight(num)
	assert.Equal(t, expected, actual)
}

func Test_hammingWeight2(t *testing.T) {
	var num uint32 = 00000000000000000000000010000000
	expected := 1
	actual := hammingWeight(num)
	assert.Equal(t, expected, actual)
}

// func Test_hammingWeight3(t *testing.T) {
// 	var num uint32 = 11111111111111111111111111111101
// 	expected := 31
// 	actual := hammingWeight(num)
// 	assert.Equal(t, expected, actual)
// }
