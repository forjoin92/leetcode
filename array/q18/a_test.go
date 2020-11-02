package q18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSumClosest(t *testing.T) {
	nums1 := []int{-1, 2, 1, -4}
	target1 := 1
	expected1 := 2
	actual1 := threeSumClosest(nums1, target1)
	assert.Equal(t, expected1, actual1)
}
