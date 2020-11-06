package q23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortByBits1(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	expected := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}
	actual := sortByBits(arr)
	assert.Equal(t, expected, actual)
}

func Test_sortByBits2(t *testing.T) {
	arr := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	expected := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	actual := sortByBits(arr)
	assert.Equal(t, expected, actual)
}

func Test_sortByBits3(t *testing.T) {
	arr := []int{10000, 10000}
	expected := []int{10000, 10000}
	actual := sortByBits(arr)
	assert.Equal(t, expected, actual)
}

func Test_sortByBits4(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13, 17, 19}
	expected := []int{2, 3, 5, 17, 7, 11, 13, 19}
	actual := sortByBits(arr)
	assert.Equal(t, expected, actual)
}

func Test_sortByBits5(t *testing.T) {
	arr := []int{10, 100, 1000, 10000}
	expected := []int{10, 100, 10000, 1000}
	actual := sortByBits(arr)
	assert.Equal(t, expected, actual)
}
