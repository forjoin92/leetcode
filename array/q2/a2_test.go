package q2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	expected1 := 7
	profit1 := maxProfit(prices1)
	assert.Equal(t, expected1, profit1)

	prices2 := []int{1, 2, 3, 4, 5}
	expected2 := 4
	profit2 := maxProfit(prices2)
	assert.Equal(t, expected2, profit2)

	prices3 := []int{7, 6, 4, 3, 1}
	expected3 := 0
	profit3 := maxProfit(prices3)
	assert.Equal(t, expected3, profit3)
}
