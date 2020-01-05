package t209

import (
	"fmt"
	"math"
	"testing"
)

func minSubArrayLen(s int, nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	min := math.MaxInt64
	isFind := false
	for i := 0; i < len(nums); i++ {
		if nums[i] >= s {
			return 1
		}
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if sum+nums[j] >= s && j-i < min {
				min = j - i
				isFind = true
				break
			}
			sum += nums[j]
		}
	}

	if isFind {
		return min + 1
	}
	return 0
}

func TestTest(t *testing.T) {
	numss := [][]int{{10, 2, 3}, {5, 1, 3, 5, 10, 7, 4, 9, 2, 8}, {1, 4, 4}, {2, 3, 1, 2, 4, 3}}
	for k, s := range []int{6, 15, 4, 7} {
		fmt.Println(minSubArrayLen(s, numss[k]))
	}
}
