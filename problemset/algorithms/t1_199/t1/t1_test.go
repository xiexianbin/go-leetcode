package t1

import (
	"fmt"
)

func twoSum2(nums []int, target int) []int {
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
			if target-num == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	ansMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if index, ok := ansMap[nums[i]]; ok {
			return []int{index, i}
		} else {
			ansMap[target-nums[i]] = i
		}
	}

	return []int{-1, -1}
}

func ExampleTwoSum() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))

	// Output:
	// [0 1]
}
