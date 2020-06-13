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
		if i >= length-switchs {
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
			// nums[i], nums[length-switchs-1] = nums[length-switchs-1], nums[i]
			for j := i; j < length-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			i--
			switchs++
		}
	}

	return length - switchs
}

func ExampleRemoveDuplicates() {
	numss := [][]int{{1, 1, 1, 1, 2}, {1, 1, 1}, {0, 0, 1, 1, 1, 1, 2, 3, 3}, {1, 1, 1, 2, 2, 3}}

	for _, nums := range numss {
		c := removeDuplicates(nums)
		fmt.Println(nums[:c])
	}

	// Output:
	// [1 1 2]
	// [1 1]
	// [0 0 1 1 2 3 3]
	// [1 1 2 2 3]
}
