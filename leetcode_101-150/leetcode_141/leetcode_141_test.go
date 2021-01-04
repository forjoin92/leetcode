package leetcode_141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	l1 := &ListNode{
		Val: 3,
	}
	l2 := &ListNode{
		Val: 2,
	}
	l3 := &ListNode{
		Val: 0,
	}
	l4 := &ListNode{
		Val: -4,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l2

	head := l1
	expected := true
	res := hasCycle(head)
	assert.Equal(t, expected, res)
}

func Test_hasCycle1(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l1,
	}
	l1.Next = l2
	head := l1
	expected := true
	res := hasCycle(head)
	assert.Equal(t, expected, res)
}

func Test_hasCycle2(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	expected := false
	res := hasCycle(head)
	assert.Equal(t, expected, res)
}

func Test_hasCycle3(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
	}
	l2 := &ListNode{
		Val: 2,
	}
	l1.Next = l2
	head := l1
	expected := false
	res := hasCycle(head)
	assert.Equal(t, expected, res)
}
