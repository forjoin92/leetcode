package leetcode_38

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay1(t *testing.T) {
	n := 1
	expected := "1"
	actual := countAndSay(n)
	assert.Equal(t, expected, actual)
}

func Test_countAndSay2(t *testing.T) {
	n := 4
	expected := "1211"
	actual := countAndSay(n)
	assert.Equal(t, expected, actual)
}
func Test_countAndSay3(t *testing.T) {
	n := 5
	expected := "111221"
	actual := countAndSay(n)
	assert.Equal(t, expected, actual)
}
