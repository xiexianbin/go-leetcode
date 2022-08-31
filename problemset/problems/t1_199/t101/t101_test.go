package t101

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recures(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (right == nil && left != nil) {
		return false
	}

	if left.Val == right.Val && recures(left.Left, right.Right) && recures(left.Right, right.Left) {
		return true
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recures(root.Left, root.Right)
}

func ExampleIsSymmetric() {
	fmt.Println(isSymmetric(&TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 1, Left: nil, Right: nil}}))

	// Output:
	// true
}
