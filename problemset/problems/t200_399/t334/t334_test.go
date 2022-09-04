package t334

import (
	"fmt"
)

func increasingTriplet(nums []int) bool {
	min, mid := 1<<31-1, 1<<31-1
	for i := 0; i < len(nums); i++ {
		if nums[i] <= min {
			min = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}

	return false
}

func ExampleIncreasingTriplet() {
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))

	// Output:
	// false
	// true
}
