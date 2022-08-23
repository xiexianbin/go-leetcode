package t257

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}

	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		if path == "" {
			path = fmt.Sprintf("%d", root.Val)
		} else {
			path = fmt.Sprintf("%s->%d", path, root.Val)
		}

		if root.Left == nil && root.Right == nil {
			ans = append(ans, path)
		}

		if root.Left != nil {
			dfs(root.Left, path)
		}
		if root.Right != nil {
			dfs(root.Right, path)
		}
	}

	dfs(root, "")
	return ans
}

func ExampleBinaryTreePaths() {
	fmt.Println(binaryTreePaths(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val: 5, Left: nil, Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}))

	// Output:
	// [1->2->5 1->3]
}
