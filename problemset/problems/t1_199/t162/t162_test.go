package t162

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			return i
		}
	}

	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums) - 1
	}
	return -1
}

func ExampleFindPeakElement() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))

	// Output:
	// 2
}
