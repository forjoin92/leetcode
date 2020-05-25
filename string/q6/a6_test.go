package q6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	str1 := "42"
	expected1 := 42
	res1 := myAtoi(str1)
	assert.Equal(t, expected1, res1)

	str2 := "   -42"
	expected2 := -42
	res2 := myAtoi(str2)
	assert.Equal(t, expected2, res2)

	str3 := "4193 with words"
	expected3 := 4193
	res3 := myAtoi(str3)
	assert.Equal(t, expected3, res3)

	str4 := "words and 987"
	expected4 := 0
	res4 := myAtoi(str4)
	assert.Equal(t, expected4, res4)

	str5 := "-91283472332"
	expected5 := -2147483648
	res5 := myAtoi(str5)
	assert.Equal(t, expected5, res5)

	str6 := "9223372036854775808"
	expected6 := 2147483647
	res6 := myAtoi(str6)
	assert.Equal(t, expected6, res6)
}
