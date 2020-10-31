package t155

import (
	"fmt"
)

type MinStack struct {
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	l := len(this.stack)
	if l > 0 {
		this.stack = this.stack[:l-1]
	}
}

func (this *MinStack) Top() int {
	l := len(this.stack)
	if l > 0 {
		return this.stack[l-1]
	}
	return -9999
}

func (this *MinStack) GetMin() int {
	if len(this.stack) > 0 {
		min := this.stack[0]
		for i := 1; i < len(this.stack); i++ {
			if this.stack[i] < min {
				min = this.stack[i]
			}
		}
		return min
	}
	return -9999
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
/* https://leetcode-cn.com/explore/learn/card/queue-stack/218/stack-last-in-first-out-data-structure/877/ */
func ExampleTest218() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())

	// Output:
	//-3
	//0
	//-2

}
