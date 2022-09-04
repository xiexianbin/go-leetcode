package t687

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		cLeft, cRight := 0, 0
		if root.Left != nil && root.Val == root.Left.Val {
			cLeft = left + 1
		}
		if root.Right != nil && root.Val == root.Right.Val {
			cRight = right + 1
		}

		ans = max(ans, cLeft+cRight)
		return max(cLeft, cRight)
	}

	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ExampleLongestUnivaluePath() {
	fmt.Println(longestUnivaluePath(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}))

	// Output:
	// 2
}
