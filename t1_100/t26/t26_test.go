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

func TestT26(t *testing.T) {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
