package t203

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for {
		if head == nil {
			break
		}
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
	}
	p := head
	for {
		if p == nil || p.Next == nil {
			break
		}
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return head
}

func TestRemoveElements(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	removeElements(head, 6)
}
