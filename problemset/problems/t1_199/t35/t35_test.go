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

// binary search
//func searchInsert(nums []int, target int) int {
//	l := len(nums)
//	if target <= nums[0] {
//		return 0
//	}
//	if target > nums[l-1] {
//		return l
//	}
//	left, right := 0, l-1
//	for left <= right {
//		m := (left + right) / 2
//		if nums[m] == target {
//			return m
//		}
//		if nums[m] < target {
//			if target <= nums[m+1] {
//				return m + 1
//			}
//			left = m + 1
//		} else {
//			if nums[m-1] < target {
//				return m
//			}
//			right = m + 1
//		}
//	}
//	return -1
//}

func TestSearchInsert(t *testing.T) {
	//nums := []int{1,3,5,6}
	//target := 5
	//target := 2
	//target := 7
	//target := 0
	nums := []int{1}
	target := 1
	fmt.Println(searchInsert(nums, target))

	// Output:
}
