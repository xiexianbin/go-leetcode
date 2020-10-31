package t136

import (
	"fmt"
)

func singleNumber(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}

func ExampleSingleNumber() {
	numss := [][]int{{2, 2, 1}, {4, 1, 2, 1, 2}}
	for _, nums := range numss {
		fmt.Println(singleNumber(nums))
	}

	// Output:
	// 1
	// 4
}
