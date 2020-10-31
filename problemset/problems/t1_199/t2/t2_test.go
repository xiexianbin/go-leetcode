package t2

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a := l1
	b := l2
	r := &ListNode{Val: 0, Next: nil}
	h := r
	var tmp int
	for a != nil || b != nil {
		c := &ListNode{Val: 0, Next: nil}
		var av, bv int
		if a != nil {
			av = a.Val
		}
		if b != nil {
			bv = b.Val
		}
		c.Val = (av + bv + tmp) % 10
		tmp = (av + bv + tmp) / 10
		if a == nil || a.Next == nil {
			a = nil
		} else {
			a = a.Next
		}
		if b == nil || b.Next == nil {
			b = nil
		} else {
			b = b.Next
		}
		r.Next = c
		r = c
	}

	if tmp > 0 {
		r.Next = &ListNode{Val: tmp, Next: nil}
	}

	return h.Next
}

func ExampleAddTwoNumbers() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	r := addTwoNumbers(l1, l2)

	fmt.Println(r.Val)
	for r.Next != nil {
		fmt.Println(r.Val)
		r = r.Next
	}

	// Output:
	// 7
	// 7
	// 0
}
