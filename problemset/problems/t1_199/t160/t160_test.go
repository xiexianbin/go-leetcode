package t160

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]struct{})
	for headA != nil {
		visited[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := visited[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

func TestGetIntersectionNode(t *testing.T) {
	getIntersectionNode(nil, nil)
}
