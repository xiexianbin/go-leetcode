package t42

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

func maxSubArray2(nums []int) int {
	ans := nums[0]
	dp := make([]int, len(nums)+1)
	for i, num := range nums {
		if i > 0 && dp[i-1] > 0 {
			dp[i] = dp[i-1] + num
		} else {
			dp[i] = num
		}

		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func ExampleMaxSubArray() {
	numss := [][]int{{-2, 1, -3, 4, -1, 2, 1, -5, 4}, {-2, 1, -3, 4}, {-1, -2, 0, -3}, {-1, -2, -3}}

	for _, nums := range numss {
		fmt.Println(maxSubArray2(nums))
	}

	// Output:
	// 6
	// 4
	// 0
	// -1
}
