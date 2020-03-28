package t448

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	result := []int{}

	for i, v := range nums {
		for {
			if i+1 != v && nums[v-1] != v {
				nums[i], nums[v-1] = nums[v-1], nums[i]
				v = nums[i]
			}
			if i+1 == v || nums[v-1] == v {
				break
			}
		}
	}

	for i, v := range nums {
		if i+1 != v {
			result = append(result, i+1)
		}
	}

	return result
}

func ExampleFindDisappearedNumbers() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))

	// Output:
	// [5 6]
}
