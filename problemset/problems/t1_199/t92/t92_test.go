package t92

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if right == left {
		return head
	}

	ans, point, index, flag := head, head, 1, false
	next, newHead, newEnd, cutNode := new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	for {
		if left == 1 && flag == false {
			flag = true
			cutNode = nil
			newEnd = point
			newHead = newEnd

			point = point.Next
			index++
			continue
		} else if index+1 == left && flag == false {
			next = point.Next
			flag = true
			cutNode = point // cut
			newEnd = next
			newHead = newEnd

			point = next
			index++
			continue
		}
		if flag == true && index < right {
			next = point.Next
			t := newHead
			newHead = point
			newHead.Next = t
			if index == left {
				newEnd.Next = nil
				if cutNode != nil {
					cutNode.Next = nil
				}
			}

			point = next
			index++
			continue
		}
		if right == index {
			next = point.Next
			t := newHead
			newHead = point
			newHead.Next = t

			newEnd.Next = next
			if cutNode == nil {
				ans = newHead
			} else {
				cutNode.Next = newHead
			}

			break
		}

		point = point.Next
		index++
	}

	return ans
}

func TestReverseBetween(t *testing.T) {
	nodes := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	//l, r := 2, 4
	l, r := 1, 2
	nl := reverseBetween(nodes, l, r)
	for {
		if nl != nil {
			fmt.Println(nl.Val)
			nl = nl.Next
		} else {
			break
		}
	}
}
