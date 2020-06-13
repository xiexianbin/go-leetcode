package t92

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	first, next, pNext, last := new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	f := false
	i := head
	c := 1
	for {
		if m == c+1 {
			first = i
			last = i.Next
			i = i.Next
			last.Next = nil
			f = true
			c++
			continue
		}
		if n == c {
			break
		}

		if f {
			next, pNext = i.Next, i.Next
			pNext.Next = i
		}

		i = next
		c++
	}

	first.Next = i
	last.Next = next

	return head
}

func TestReverseBetween(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	m, n := 2, 4
	nl := reverseBetween(l, m, n)
	for {
		if nl.Next != nil {
			fmt.Println(nl.Val)
			nl = nl.Next
		}
	}

}
