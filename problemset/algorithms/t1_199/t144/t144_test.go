package t144

import (
	"fmt"
	"testing"
)

type Stack struct {
	Trees []*TreeNode
}

func (this *Stack) Pop() *TreeNode {
	if len(this.Trees) > 0 {
		r := this.Trees[len(this.Trees)-1]
		this.Trees = this.Trees[:len(this.Trees)-1]
		return r
	}
	return nil
}

func (this *Stack) Push(v *TreeNode) {
	this.Trees = append(this.Trees, v)
}

func (this *Stack) IsEmpty() bool {
	if len(this.Trees) > 0 {
		return false
	}
	return true
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	result := []int{}
	stack := &Stack{
		Trees: []*TreeNode{},
	}
	stack.Push(root)
	for stack.IsEmpty() == false {
		tn := stack.Pop()
		result = append(result, tn.Val)
		if tn.Right != nil {
			stack.Push(tn.Right)
		}
		if tn.Left != nil {
			stack.Push(tn.Left)
		}
	}

	return result
}

func TestT144(t *testing.T) {
	root := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}
	fmt.Println(preorderTraversal(root))
}
