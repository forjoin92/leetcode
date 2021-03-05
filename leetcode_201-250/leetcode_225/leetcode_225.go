package leetcode_225

import "container/list"

type MyStack struct {
	a *list.List
	b *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		a: list.New(),
		b: list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	_ = this.b.PushBack(x)
	for this.a.Len() != 0 {
		e := this.a.Front()
		this.b.PushBack(e.Value)
		this.a.Remove(e)
	}
	this.a, this.b = this.b, this.a
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	e := this.a.Front()
	tmp := e.Value
	this.a.Remove(e)
	return tmp.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.a.Front().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.a.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
