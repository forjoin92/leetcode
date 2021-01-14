package leetcode_1018

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_prefixesDivBy51(t *testing.T) {
	nums := []int{0, 1, 1}
	expected := []bool{true, false, false}
	actual := prefixesDivBy5(nums)
	assert.Equal(t, expected, actual)
}

func Test_prefixesDivBy52(t *testing.T) {
	nums := []int{1, 1, 1}
	expected := []bool{false, false, false}
	actual := prefixesDivBy5(nums)
	assert.Equal(t, expected, actual)
}

func Test_prefixesDivBy53(t *testing.T) {
	nums := []int{0, 1, 1, 1, 1, 1}
	expected := []bool{true, false, false, false, true, false}
	actual := prefixesDivBy5(nums)
	assert.Equal(t, expected, actual)
}

func Test_prefixesDivBy54(t *testing.T) {
	nums := []int{1, 1, 1, 0, 1}
	expected := []bool{false, false, false, false, false}
	actual := prefixesDivBy5(nums)
	assert.Equal(t, expected, actual)
}

func Test_prefixesDivBy55(t *testing.T) {
	nums := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}
	expected := []bool{false, false, true, false, false, false, false, false, false, false, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, true, false, false, true, true, true, true, true, true, true, false, false, true, false, false, false, false, true, true}
	actual := prefixesDivBy5(nums)
	assert.Equal(t, expected, actual)
}
