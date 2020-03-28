package t19

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	if p2 == nil {
		return head.Next
	}

	for {
		if p2.Next == nil {
			p1.Next = p1.Next.Next
			break
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	return head
}

func TestTest(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	head = removeNthFromEnd(head, 2)
	cur := head
	for {
		fmt.Printf("%d ", cur.Val)
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	fmt.Println()

	// Output:
}
