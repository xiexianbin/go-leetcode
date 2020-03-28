package t104

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	mR := maxDepth(root.Left) + 1
	rR := maxDepth(root.Right) + 1
	if mR > rR {
		return mR
	} else {
		return rR
	}
}

func TestTest(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: nil, Right: nil}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}}
	fmt.Println(maxDepth(root))

	// Output:
}
