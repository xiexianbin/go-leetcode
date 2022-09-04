package t98

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recures(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return recures(root.Left, lower, root.Val) && recures(root.Right, root.Val, upper)
}

func isValidBST(root *TreeNode) bool {
	return recures(root, math.MinInt64, math.MaxInt64)
}

func ExampleT() {
	fmt.Println(isValidBST(&TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: nil, Right: nil}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}}))
	fmt.Println(isValidBST(&TreeNode{Val: 0, Left: &TreeNode{Val: -1, Left: nil, Right: nil}, Right: nil}))
	fmt.Println(isValidBST(&TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}))

	// Output:
	// false
	// true
	// true
}
