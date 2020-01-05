package t61

import (
	"fmt"
	"testing"
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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := head
	l := 1
	for {
		if head.Next == nil {
			head.Next = h
			head = head.Next
			break
		}
		head = head.Next
		l++
	}

	k = l - k%l
	c := head
	for i := 0; i < k; i++ {
		c = head
		head = head.Next
		if i == k-1 {
			c.Next = nil
		}
	}

	return head
}

func TestTest(t *testing.T) {
	head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}
	head = rotateRight(head, 4)
	for {
		if head == nil {
			break
		}
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}
