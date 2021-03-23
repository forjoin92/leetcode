package leetcode_72

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDistance1(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	expected := 3
	actual := minDistance(word1, word2)
	assert.Equal(t, expected, actual)
}

func Test_minDistance2(t *testing.T) {
	word1 := "intention"
	word2 := "execution"
	expected := 5
	actual := minDistance(word1, word2)
	assert.Equal(t, expected, actual)
}
