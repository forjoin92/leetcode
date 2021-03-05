package leetcode_225

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyStack(t *testing.T) {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	assert.Equal(t, 2, myStack.Top())
	assert.Equal(t, 2, myStack.Pop())
	assert.Equal(t, false, myStack.Empty())
}
