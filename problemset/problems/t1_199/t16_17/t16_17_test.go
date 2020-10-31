package t16_17

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0
	for i:=0; i<len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

func ExampleMaxSubArray() {
	nums := [][]int{{-2, 1, -3, 4, -1, 2, 1, -5, 4}, {8, -19, 5, -4, 20}, {1, -1, 1}}
	for _, num := range nums {
		fmt.Println(maxSubArray(num))
	}

	// Output:
	// 6
	// 21
	// 1
}
