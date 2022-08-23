package offer03

import (
	"fmt"
	"sort"
)

func findRepeatNumber(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return nums[i]
		}
	}
	return -1
}

func ExampleFindRepeatNumber() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))

	// Output:
	// 2
}
