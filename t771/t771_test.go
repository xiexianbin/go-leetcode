package t771

import (
	"fmt"
	"testing"
)

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if i != index {
			if max >= nums[i] * 2 {
				continue
			} else {
				return -1
			}
		}
	}
	return index
}

func TestTest(t *testing.T) {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))

	nums = []int{1, 2, 3, 4}
	fmt.Println(dominantIndex(nums))
}
