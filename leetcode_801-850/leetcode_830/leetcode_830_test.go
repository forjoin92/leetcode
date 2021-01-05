package leetcode_830

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largeGroupPositions(t *testing.T) {
	s1 := "abbxxxxzzy"
	expected1 := [][]int{{3, 6}}
	res1 := largeGroupPositions(s1)
	assert.Equal(t, expected1, res1)

	s2 := "abc"
	var expected2 [][]int
	res2 := largeGroupPositions(s2)
	assert.Equal(t, expected2, res2)

	s3 := "abcdddeeeeaabbbcd"
	expected3 := [][]int{{3, 5}, {6, 9}, {12, 14}}
	res3 := largeGroupPositions(s3)
	assert.Equal(t, expected3, res3)

	s4 := "aba"
	var expected4 [][]int
	res4 := largeGroupPositions(s4)
	assert.Equal(t, expected4, res4)
}
