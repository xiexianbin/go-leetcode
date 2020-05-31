package t80

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	count, switchs := 1, 0
	t := nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		if i > length - switchs {
			break
		}
		if nums[i] != t {
			t = nums[i]
			count = 1
			continue
		} else {
			count++
		}

		if count > 2 {
			nums[i], nums[length-switchs-1] = nums[length-switchs-1], nums[i]
			switchs ++
		}
	}

	return len(nums) - switchs
}

func ExampleRemoveDuplicates() {
	numss := [][]int{{0, 0, 1, 1, 1, 1, 2, 3, 3}, {1, 1, 1, 2, 2, 3}}

	for _, nums := range numss {
		fmt.Println(removeDuplicates(nums))
	}

	// Output:
	// 7
	// 5
}
