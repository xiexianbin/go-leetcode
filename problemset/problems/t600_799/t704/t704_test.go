package t704

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}

func ExampleSearch() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))

	// Output:
	// 4
}
