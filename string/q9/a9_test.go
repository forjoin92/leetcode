package q9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	strs1 := []string{"flower", "flow", "flight"}
	expected1 := "fl"
	res1 := longestCommonPrefix(strs1)
	assert.Equal(t, expected1, res1)

	strs2 := []string{"dog", "racecar", "car"}
	expected2 := ""
	res2 := longestCommonPrefix(strs2)
	assert.Equal(t, expected2, res2)

	strs3 := []string{"ab", "a"}
	expected3 := "a"
	res3 := longestCommonPrefix(strs3)
	assert.Equal(t, expected3, res3)
}
