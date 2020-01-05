package t622

import (
	"fmt"
	"testing"
)

type MyCircularQueue struct {
	queue []int
	front int
	rear  int
	size  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		front: -1,
		rear:  -1,
		size:  k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.rear = (this.rear + 1) % this.size
	if this.front == -1 {
		this.front = this.rear
	}
	this.queue[this.rear] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.rear == this.front {
		this.rear = -1
		this.front = -1
	} else {
		this.front = (this.front + 1) % this.size
	}

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.rear]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.front == -1 && this.rear == -1 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.front == (this.rear+1)%this.size {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

/* https://leetcode-cn.com/explore/learn/card/queue-stack/216/queue-first-in-first-out-data-structure/865/ */
func Test216(t *testing.T) {
	k := 3
	obj := Constructor(k)
	fmt.Println(obj.EnQueue(1))
	fmt.Println(obj.EnQueue(2))
	fmt.Println(obj.EnQueue(3))
	fmt.Println(obj.EnQueue(4))

	fmt.Println(obj.Rear())
	fmt.Println(obj.IsFull())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.EnQueue(4))
	fmt.Println(obj.Rear())

	fmt.Println(obj.DeQueue())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.DeQueue())

	fmt.Println(obj.Front())
	fmt.Println(obj.IsEmpty())
}
