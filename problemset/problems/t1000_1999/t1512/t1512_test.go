package t1512

import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	count := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

func Example_numIdenticalPairs() {
	inputs := [][]int{{1, 2, 3, 1, 1, 3}, {1, 1, 1, 1}, {1, 2, 3}}
	for _, nums := range inputs {
		fmt.Println(numIdenticalPairs(nums))
	}

	// Output:
	// 4
	// 6
	// 0
}
