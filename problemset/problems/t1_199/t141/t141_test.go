package t141

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	flag := make(map[*ListNode]bool)
	for ; head != nil; {
		if _, ok := flag[head]; ok {
			return true
		}
		flag[head] = true
		head = head.Next
	}

	return false
}

func ExampleHasCycle() {
	fmt.Println(hasCycle(&ListNode{Val: 3, Next: &ListNode{Val: 2}}))

	// Output:
	// false
}
