package leetcode_350

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersect(t *testing.T) {
	nums11 := []int{1, 2, 2, 1}
	nums12 := []int{2, 2}
	expected1 := []int{2, 2}
	res1 := intersect(nums11, nums12)
	assert.Subset(t, expected1, res1)

	nums21 := []int{4, 9, 5}
	nums22 := []int{9, 4, 9, 8, 4}
	expected2 := []int{4, 9}
	res2 := intersect(nums21, nums22)
	assert.Subset(t, expected2, res2)
}
