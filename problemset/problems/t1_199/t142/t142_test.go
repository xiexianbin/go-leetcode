package t142

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := visited[head]; ok {
			return head
		} else {
			visited[head] = struct{}{}
		}

		head = head.Next
	}

	return nil
}

func TestDetectCycle(t *testing.T) {
	detectCycle(nil)
}
