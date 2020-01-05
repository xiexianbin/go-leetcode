package t268

import (
	"fmt"
	"sort"
	"testing"
)

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return -1
}

func TestTest(t *testing.T) {
	for _, nums := range [][]int{{}, {0}, {3, 0, 1}, {9, 6, 4, 2, 3, 5, 7, 0, 1}, {1, 2, 3}} {
		fmt.Println(missingNumber(nums))
	}
}
