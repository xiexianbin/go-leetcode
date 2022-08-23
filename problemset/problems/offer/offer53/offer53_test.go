package offer53

import (
	"fmt"
)

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}

func ExampleMissingNumber() {
	inputs := [][]int{{0}, {0, 1, 3}, {0, 1, 2, 3, 4, 5, 6, 7, 9}}
	for _, input := range inputs {
		fmt.Println(missingNumber(input))
	}

	// Output:
	// 1
	// 2
	// 8
}
