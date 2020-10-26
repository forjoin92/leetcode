package q13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallerNumbersThanCurrent1(t *testing.T) {
	nums := []int{8, 1, 2, 2, 3}
	expected := []int{4, 0, 1, 1, 3}
	actual := smallerNumbersThanCurrent(nums)
	assert.Equal(t, expected, actual)
}
