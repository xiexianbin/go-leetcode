package t1464

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	first, second := nums[0], math.MinInt64
	for i := 1; i < len(nums); i++ {
		if nums[i] > first {
			second = first
			first = nums[i]
		} else if nums[i] > second {
			second = nums[i]
		}
	}

	return (first - 1) * (second - 1)
}

func ExampleMaxProduct() {
	fmt.Println(maxProduct([]int{2, 4, 5, 1}))

	// Output:
	// 12
}
