package leetcode_300

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLIS1(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	expected := 4
	actual := lengthOfLIS(nums)
	assert.Equal(t, expected, actual)
}

func Test_lengthOfLIS2(t *testing.T) {
	nums := []int{0, 1, 0, 3, 2, 3}
	expected := 4
	actual := lengthOfLIS(nums)
	assert.Equal(t, expected, actual)
}

func Test_lengthOfLIS3(t *testing.T) {
	nums := []int{7, 7, 7, 7, 7, 7, 7}
	expected := 1
	actual := lengthOfLIS(nums)
	assert.Equal(t, expected, actual)
}
