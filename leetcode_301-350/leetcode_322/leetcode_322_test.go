package leetcode_322

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coinChange1(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	expected := 3
	actual := coinChange(coins, amount)
	assert.Equal(t, expected, actual)
}

func Test_coinChange2(t *testing.T) {
	coins := []int{2}
	amount := 3
	expected := -1
	actual := coinChange(coins, amount)
	assert.Equal(t, expected, actual)
}

func Test_coinChange3(t *testing.T) {
	coins := []int{1}
	amount := 0
	expected := 0
	actual := coinChange(coins, amount)
	assert.Equal(t, expected, actual)
}

func Test_coinChange4(t *testing.T) {
	coins := []int{1}
	amount := 1
	expected := 1
	actual := coinChange(coins, amount)
	assert.Equal(t, expected, actual)
}

func Test_coinChange5(t *testing.T) {
	coins := []int{1}
	amount := 2
	expected := 2
	actual := coinChange(coins, amount)
	assert.Equal(t, expected, actual)
}
