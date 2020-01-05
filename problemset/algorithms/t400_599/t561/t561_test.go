package t561

import (
	"fmt"
	"sort"
	"testing"
)

func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func TestTest(t *testing.T) {
	for _, nums := range [][]int{{1, 4, 3, 2}} {
		fmt.Println(arrayPairSum(nums))
	}
}
