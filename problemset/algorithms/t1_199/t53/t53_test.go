package t53

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

func ExampleMaxSubArray() {
	numss := [][]int{{-2, 1, -3, 4, -1, 2, 1, -5, 4}, {-2, 1, -3, 4}, {-1, -2, 0, -3}, {-1, -2, -3}}

	for _, nums := range numss {
		fmt.Println(maxSubArray(nums))
	}

	// Output:
	// 6
	// 4
	// 0
	// -1
}
