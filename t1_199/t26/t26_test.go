package t26

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	low := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[low] != nums[fast] {
			low++
			nums[low] = nums[fast]
		}
	}
	return low + 1
}

func removeDuplicates2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[l-1] {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}

	return l
}

func TestT26(t *testing.T) {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
