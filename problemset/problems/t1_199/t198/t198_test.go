package t198

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	ans := nums[0]
	for i := 1; i <= len(nums); i++ {
		if i < 2 {
			dp[i] = nums[i-1]
		} else if i == 2 {
			dp[i] = dp[0] + nums[i-1]
		} else {
			dp[i] = max(dp[i-2], dp[i-3]) + nums[i-1]
		}

		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

func ExampleT() {
	fmt.Println(rob([]int{1, 2, 3, 1}))

	// Output:
	// 4
}
