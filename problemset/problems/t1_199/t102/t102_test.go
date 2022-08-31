package t102

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func level(cur []*TreeNode) ([]*TreeNode, []int) {
	chird := []*TreeNode{}
	ans := []int{}
	flag := false
	for _, t := range cur {
		if t != nil {
			flag = true
			chird = append(chird, t.Left, t.Right)
			ans = append(ans, t.Val)
		}
	}
	if flag {
		return chird, ans
	} else {
		return nil, nil
	}
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{{root.Val}}

	chird := []*TreeNode{root.Left, root.Right}
	var ans []int
	for {
		chird, ans = level(chird)
		if ans == nil {
			break
		}
		result = append(result, ans)
	}

	return result
}

func ExampleLevelOrder() {
	fmt.Println(levelOrder(&TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: nil, Right: nil}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 10, Left: nil, Right: nil}, Right: nil}}))

	// Output:
	// [[3] [9 9] [10]]
}
