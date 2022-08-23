package t655

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxHeight(root *TreeNode) int {
	h := 0
	if root.Left != nil {
		h = maxHeight(root.Left) + 1
	}
	if root.Right != nil {
		h = max(h, maxHeight(root.Right)+1)
	}
	return h
}

func printTree(root *TreeNode) [][]string {
	// calculate binary tree height
	height := maxHeight(root)
	m := height + 1
	n := 1<<m - 1

	// init array
	res := make([][]string, m)
	for i := 0; i < m; i++ {
		res[i] = make([]string, n)
	}

	// dfs
	var dfs func(root *TreeNode, r, c int)
	dfs = func(root *TreeNode, r, c int) {
		res[r][c] = fmt.Sprintf("%d", root.Val)
		if root.Left != nil {
			dfs(root.Left, r+1, c-1<<(height-r-1))
		}
		if root.Right != nil {
			dfs(root.Right, r+1, c+1<<(height-r-1))
		}
	}
	dfs(root, 0, (n-1)/2)

	return res
}

func ExamplePrintTree() {
	res := printTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	})
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] == "" {
				fmt.Printf("-")
			} else {
				fmt.Printf("%s", res[i][j])
			}

		}
		fmt.Println()
	}

	// Output:
	// -1-
	// 2--
}
