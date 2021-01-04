package leetcode_22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis1(t *testing.T) {
	n := 0
	expected := []string{}
	actual := generateParenthesis(n)
	assert.Equal(t, expected, actual)
}

func Test_generateParenthesis2(t *testing.T) {
	n := 1
	expected := []string{"()"}
	actual := generateParenthesis(n)
	assert.Equal(t, expected, actual)
}

func Test_generateParenthesis3(t *testing.T) {
	n := 2
	expected := []string{
		"()()",
		"(())",
	}
	actual := generateParenthesis(n)
	assert.Equal(t, expected, actual)
}

func Test_generateParenthesis4(t *testing.T) {
	n := 3
	expected := []string{
		"()()()",
		"(())()",
		"(()())",
		"()(())",
		"((()))",
	}
	actual := generateParenthesis(n)
	assert.Equal(t, expected, actual)
}
