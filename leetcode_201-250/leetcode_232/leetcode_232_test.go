package leetcode_232

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyStack(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1)
	myQueue.Push(2)
	assert.Equal(t, 1, myQueue.Peek())
	assert.Equal(t, 1, myQueue.Pop())
	assert.Equal(t, false, myQueue.Empty())
}
