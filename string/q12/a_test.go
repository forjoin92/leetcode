package q12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	s1 := "LEETCODEISHIRING"
	numRows1 := 3
	expected1 := "LCIRETOESIIGEDHN"
	res1 := convert(s1, numRows1)
	assert.Equal(t, expected1, res1)

	s2 := "LEETCODEISHIRING"
	numRows2 := 4
	expected2 := "LDREOEIIECIHNTSG"
	res2 := convert(s2, numRows2)
	assert.Equal(t, expected2, res2)

	s3 := "AB"
	numRows3 := 1
	expected3 := "AB"
	res3 := convert(s3, numRows3)
	assert.Equal(t, expected3, res3)
}
