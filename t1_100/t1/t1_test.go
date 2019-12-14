package t1

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func twoSum(nums []int, target int) []int {
	//l := len(nums)
	//for i := 0; i < l; i++ {
	//	for j := 0; j < l; j++ {
	//		if i == j {
	//			continue
	//		}
	//
	//		sum := nums[i] + nums[j]
	//		if sum == target {
	//			return []int{i, j}
	//		}
	//	}
	//}
	//return nil

	l := len(nums)
	for i, num := range nums {
		for j := i + 1; j < l; j++ {
			if target - num == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	assert.Assert(t, is.DeepEqual(twoSum(nums, target), []int{0, 1}))
}
