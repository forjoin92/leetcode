package leetcode_38

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay(t *testing.T) {
	n1 := 1
	expected1 := "1"
	res1 := countAndSay(n1)
	assert.Equal(t, expected1, res1)

	n2 := 4
	expected2 := "1211"
	res2 := countAndSay(n2)
	assert.Equal(t, expected2, res2)

	n3 := 5
	expected3 := "111221"
	res3 := countAndSay(n3)
	assert.Equal(t, expected3, res3)
}
