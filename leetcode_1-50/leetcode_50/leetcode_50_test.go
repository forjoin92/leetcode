package leetcode_50

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myPow1(t *testing.T) {
	x, n := 2.0, 10
	expected := 1024.0
	assert.Equal(t, expected, myPow(x, n))
}

func Test_myPow2(t *testing.T) {
	x, n := 2.10000, 3
	expected := 9.261
	assert.Equal(t, expected, myPow(x, n))
}

func Test_myPow3(t *testing.T) {
	x, n := 2.0, -2
	expected := 0.25
	assert.Equal(t, expected, myPow(x, n))
}
