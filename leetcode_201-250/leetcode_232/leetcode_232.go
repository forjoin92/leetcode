package leetcode_232

type MyQueue struct {
	a []int
	b []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.a = append(this.a, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.b) == 0 {
		for i := len(this.a) - 1; i >= 0; i-- {
			this.b = append(this.b, this.a[i])
			this.a = this.a[:len(this.a)-1]
		}
	}
	a := this.b[len(this.b)-1]
	this.b = this.b[:len(this.b)-1]
	return a
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.b) == 0 {
		for i := len(this.a) - 1; i >= 0; i-- {
			this.b = append(this.b, this.a[i])
			this.a = this.a[:len(this.a)-1]
		}
	}
	return this.b[len(this.b)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.a) == 0 && len(this.b) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
