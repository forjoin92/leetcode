package leetcode_1006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_clumsy1(t *testing.T) {
	N := 4
	expected := 7
	actual := clumsy(N)
	assert.Equal(t, expected, actual)
}

func Test_clumsy2(t *testing.T) {
	N := 10
	expected := 12
	actual := clumsy(N)
	assert.Equal(t, expected, actual)
}

func Test_clumsy3(t *testing.T) {
	N := 3
	expected := 6
	actual := clumsy(N)
	assert.Equal(t, expected, actual)
}
