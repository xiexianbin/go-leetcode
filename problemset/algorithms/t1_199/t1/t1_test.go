package t1

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	//l := len(nums)
	//for i := 0; i < l; i++ {
	//	for j := 0; j < l; j++ {
	//		if i == j {
	//			continue
	//		}
	//
	//		sum := nums[i] + nums[j]
	//		if sum == target {
	//			return []int{i, j}
	//		}
	//	}
	//}
	//return nil

	l := len(nums)
	for i, num := range nums {
		for j := i + 1; j < l; j++ {
			if target - num == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func ExampleTwoSum() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))

	// Output:
	// [0 1]
}
