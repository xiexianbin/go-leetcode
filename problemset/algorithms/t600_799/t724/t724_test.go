package t724

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	leftsum := 0
	for i := 0; i < len(nums); i++ {
		if leftsum == sum - nums[i] - leftsum {
			return i
		}
		leftsum += nums[i]
	}
	return -1
}

func ExamplePivotIndex() {
	//nums := []int{1, 7, 3, 6, 5, 6}
	nums := []int{1, 2, 3}
	fmt.Println(pivotIndex(nums))

	// Output:
	// -1
}
