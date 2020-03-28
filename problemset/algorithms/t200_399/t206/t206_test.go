package t206

import (
	"fmt"
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
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for {
		if head == nil {
			break
		}
		cur := head
		head = head.Next
		cur.Next = prev
		prev = cur
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func ExampleReverseList() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	h := reverseList(head)
	//h := reverseList2(head)
	for {
		if h.Next == nil {
			break
		}
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
	fmt.Printf("%d\n", h.Val)

	// Output:
	// 5 4 3 2 1
}
