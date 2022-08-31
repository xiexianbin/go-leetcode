package t662

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func dfs(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	lh := dfs(root.Left)
//	rh := dfs(root.Right)
//	return max(lh, rh) + 1
//}

type pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	ans := 1
	line := []pair{{root, 1}}
	for len(line) > 0 {
		ans = max(ans, line[len(line)-1].index-line[0].index+1)
		t := []pair{}
		for _, n := range line {
			if n.node.Left != nil {
				t = append(t, pair{n.node.Left, n.index * 2})
			}
			if n.node.Right != nil {
				t = append(t, pair{n.node.Right, n.index*2 + 1})
			}
		}
		line = t
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ExampleT() {
	fmt.Println(widthOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}))

	// Output:
	// 2
}
