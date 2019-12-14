package t21

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// for nil ListNode
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var FirstNode, LastNode, NewNode *ListNode = nil, nil, nil
	for true {
		// for normal
		if l1.Val > l2.Val {
			NewNode = l2
			l2 = l2.Next
		} else {
			NewNode = l1
			l1 = l1.Next
		}

		if LastNode == nil {
			FirstNode = NewNode
			LastNode = NewNode
		} else {
			LastNode.Next = NewNode
			LastNode = NewNode
		}

		// if l1 or l2 be nil
		if l1 == nil {
			LastNode.Next = l2
			break
		} else if l2 == nil {
			LastNode.Next = l1
			break
		}
	}

	return FirstNode
}

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	r := mergeTwoLists(l1, l2)

	//l1 := &ListNode{Val:1, Next:nil}
	//r := mergeTwoLists(l1, nil)

	//l1 := &ListNode{Val:2, Next:nil}
	//l2 := &ListNode{Val:1, Next:nil}
	//r := mergeTwoLists(l1, l2)

	if r != nil {
		fmt.Print(r.Val)
	}
	for r != nil {
		r = r.Next
		if r != nil {
			fmt.Print(r.Val)
		}
	}
}
