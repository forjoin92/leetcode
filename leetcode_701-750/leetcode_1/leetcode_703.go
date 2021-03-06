package leetcode_703

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	k int
	n int
	sort.IntSlice
}

func (r *KthLargest) Push(x interface{}) {
	if r.Len() < r.k {
		r.IntSlice = append(r.IntSlice, x.(int))
		if r.n > x.(int) {
			r.n = x.(int)
		}
	} else {
		if r.n < x.(int) {
			r.IntSlice = append(r.IntSlice, x.(int))
			_ = heap.Pop(r)
			r.n = heap.Pop(r).(int)
			r.IntSlice = append(r.IntSlice, r.n)
		}
	}
}

func (r *KthLargest) Pop() interface{} {
	tmp := r.IntSlice[len(r.IntSlice)-1]
	r.IntSlice = r.IntSlice[:len(r.IntSlice)-1]
	return tmp
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		k: k,
		n: 10001,
	}
	for _, num := range nums {
		heap.Push(&kth, num)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	return this.n
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
