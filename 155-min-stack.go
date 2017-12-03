package main

import (
	"fmt"
)

func main() {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.GetMin())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.GetMin())
}

type MinStack struct {
	root *node
}

type node struct {
	value, min int
	next *node
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{nil}
}


func (this *MinStack) Push(x int)  {
	n := &node{x, x, this.root}
	if this.root != nil && this.root.min < x {
		n.min = this.root.min
	}
	this.root = n
}

func (this *MinStack) Pop()  {
	if this.root != nil {
		this.root = this.root.next
	}
}

func (this *MinStack) Top() int {
    if this.root != nil {
		return this.root.value
	}
	return 0
}


func (this *MinStack) GetMin() int {
    return this.root.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */