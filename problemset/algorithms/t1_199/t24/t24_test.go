package t24

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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next.Next = swapPairs(head.Next.Next)
	c := head
	head = head.Next
	c.Next = head.Next
	head.Next = c
	return head
}

func TestTest(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	head = swapPairs(head)
	for {
		if head == nil {
			break
		}
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
