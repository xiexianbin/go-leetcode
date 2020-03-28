package t628

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	l := len(nums)
	max := nums[l-1] * nums[l-2] * nums[l-3]

	if nums[0] < 0 && nums[1] < 0 {
		c := nums[0] * nums[1] * nums[l-1]
		if c > max {
			max = c
		}
	}
	return max
}

func ExampleMaximumProduct() {
	for _, nums := range [][]int{{-4, -3, -2, -1, 60}, {1, 2, 3}, {1, 2, 3, 4}} {
		fmt.Println(maximumProduct(nums))
	}

	// Output:
	// 720
	// 6
	// 24
}
