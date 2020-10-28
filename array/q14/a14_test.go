package q14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniqueOccurrences1(t *testing.T) {
	nums := []int{1, 2, 2, 1, 1, 3}
	expected := true
	actual := uniqueOccurrences(nums)
	assert.Equal(t, expected, actual)
}

func Test_uniqueOccurrences2(t *testing.T) {
	nums := []int{1, 2}
	expected := false
	actual := uniqueOccurrences(nums)
	assert.Equal(t, expected, actual)
}

func Test_uniqueOccurrences3(t *testing.T) {
	nums := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	expected := true
	actual := uniqueOccurrences(nums)
	assert.Equal(t, expected, actual)
}
