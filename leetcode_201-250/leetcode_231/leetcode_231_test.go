package leetcode_231

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPowerOfTwo(t *testing.T) {
	n := 1
	assert.Equal(t, true, isPowerOfTwo(n))
	n = 16
	assert.Equal(t, true, isPowerOfTwo(n))
	n = 218
	assert.Equal(t, false, isPowerOfTwo(n))
}
