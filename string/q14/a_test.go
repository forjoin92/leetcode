package q14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	s1 := "III"
	expected1 := 3
	res1 := romanToInt(s1)
	assert.Equal(t, expected1, res1)

	s2 := "IV"
	expected2 := 4
	res2 := romanToInt(s2)
	assert.Equal(t, expected2, res2)

	s3 := "IX"
	expected3 := 9
	res3 := romanToInt(s3)
	assert.Equal(t, expected3, res3)

	s4 := "LVIII"
	expected4 := 58
	res4 := romanToInt(s4)
	assert.Equal(t, expected4, res4)

	s5 := "MCMXCIV"
	expected5 := 1994
	res5 := romanToInt(s5)
	assert.Equal(t, expected5, res5)
}
