package t35

import (
	"fmt"
	"testing"
)

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 || target <= nums[0] {
		return 0
	}
	for i := 1; i < l; i++ {
		if nums[i-1] < target && nums[i] >= target {
			return i
		}
	}
	return l
}

func TestSearchInsert(t *testing.T) {
	//nums := []int{1,3,5,6}
	//target := 5
	//target := 2
	//target := 7
	//target := 0
	nums := []int{1}
	target := 1
	fmt.Println(searchInsert(nums, target))
}
