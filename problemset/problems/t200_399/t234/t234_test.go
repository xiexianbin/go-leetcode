package t234

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	quick, slow := head, head
	for ; quick != nil && quick.Next != nil; {
		quick = quick.Next.Next
		slow = slow.Next
	}
	if quick != nil {
		slow = slow.Next
	}

	slow = reverse(slow)
	quick = head
	for {
		if slow == nil || quick == nil {
			break
		}
		if slow.Val != quick.Val {
			return false
		}
		slow = slow.Next
		quick = quick.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	rev := reverse(next)
	next.Next = head
	head.Next = nil
	return rev
}

func ExampleIsPalindrome() {
	fmt.Println(isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}))

	// Output:
	// false
}
