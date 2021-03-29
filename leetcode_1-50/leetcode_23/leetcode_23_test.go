package leetcode_23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}
	lists := []*ListNode{
		l1, l2, l3,
	}
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		},
	}
	actual := mergeKLists(lists)
	assert.Equal(t, expected, actual)
}
