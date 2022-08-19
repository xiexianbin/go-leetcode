package t63

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	flag := make(map[int]int)
	point := head
	prev := point
	for {
		if point == nil {
			break
		}
		if _, ok := flag[point.Val]; ok == true {
			prev.Next = point.Next
		} else {
			flag[point.Val] = 1
			prev = point
		}

		point = point.Next
	}

	return head
}

func ExampleMerge() {
	var head, result *ListNode
	//head = &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val:  2,
	//			Next: nil,
	//		},
	//	},
	//}
	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	result = deleteDuplicates(head)
	for {
		if result == nil {
			break
		}
		fmt.Print(fmt.Sprintf("%d ", result.Val))
		result = result.Next
	}

	// Output:
	// 1 2 3
}
