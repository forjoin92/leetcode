package leetcode_12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intToRoman(t *testing.T) {
	num1 := 3
	expected1 := "III"
	res1 := intToRoman(num1)
	assert.Equal(t, expected1, res1)

	num2 := 4
	expected2 := "IV"
	res2 := intToRoman(num2)
	assert.Equal(t, expected2, res2)

	num3 := 9
	expected3 := "IX"
	res3 := intToRoman(num3)
	assert.Equal(t, expected3, res3)

	num4 := 58
	expected4 := "LVIII"
	res4 := intToRoman(num4)
	assert.Equal(t, expected4, res4)

	num5 := 1994
	expected5 := "MCMXCIV"
	res5 := intToRoman(num5)
	assert.Equal(t, expected5, res5)
}
