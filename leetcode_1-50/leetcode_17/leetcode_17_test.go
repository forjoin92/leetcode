package leetcode_17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	digits1 := "23"
	expected1 := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	res1 := letterCombinations(digits1)
	assert.Equal(t, expected1, res1)
}
