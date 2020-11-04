package q17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid1(t *testing.T) {
	s := "()"
	expected := true
	res1 := isValid(s)
	assert.Equal(t, expected, res1)
}

func Test_isValid2(t *testing.T) {
	s := "()[]{}"
	expected := true
	res1 := isValid(s)
	assert.Equal(t, expected, res1)
}

func Test_isValid3(t *testing.T) {
	s := "(]"
	expected := false
	res1 := isValid(s)
	assert.Equal(t, expected, res1)
}

func Test_isValid4(t *testing.T) {
	s := "([)]"
	expected := false
	res1 := isValid(s)
	assert.Equal(t, expected, res1)
}

func Test_isValid5(t *testing.T) {
	s := "{[]}"
	expected := true
	res1 := isValid(s)
	assert.Equal(t, expected, res1)
}
