package q1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	expected1 := []byte{'o', 'l', 'l', 'e', 'h'}
	reverseString(s1)
	assert.Equal(t, expected1, s1)

	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	expected2 := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	reverseString(s2)
	assert.Equal(t, expected2, s2)
}
