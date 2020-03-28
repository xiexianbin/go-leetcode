package t23

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	head := l1
	l1p := l1
	for {
		if l1.Val <= l2.Val {
			if l1.Next == nil {
				l1.Next = l2
				break
			} else {
				l1p = l1
				l1 = l1.Next
			}
		} else {
			if l2.Next == nil {
				l2.Next = l1p.Next
				l1p.Next = l2
				break
			} else {
				cur := &ListNode{Val: l2.Val, Next: nil}
				cur.Next = l1p.Next
				l1p.Next = cur
				l1p = l1p.Next
				l2 = l2.Next
			}
		}
	}
	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	head := lists[0]
	for i := 1; i < len(lists); i++ {
		head = mergeTwoLists(head, lists[i])
	}

	return head
}

func TestTest(t *testing.T) {
	lists := []*ListNode{{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		{Val: 2, Next: &ListNode{Val: 6, Next: nil}}}
	head := mergeKLists(lists)
	for {
		if head.Next == nil {
			fmt.Printf("%d\n", head.Val)
			break
		}
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}

	// Output:
}
