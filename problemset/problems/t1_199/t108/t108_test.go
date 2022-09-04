package t108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}

	l := len(nums)
	m := l / 2
	if l == m+1 {
		return &TreeNode{
			Val:   nums[m],
			Left:  sortedArrayToBST(nums[:m]),
			Right: nil,
		}
	} else {
		return &TreeNode{
			Val:   nums[m],
			Left:  sortedArrayToBST(nums[:m]),
			Right: sortedArrayToBST(nums[m+1:]),
		}
	}
}

func ExampleSortedArrayToBST() {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})

	// Output:
	//
}
