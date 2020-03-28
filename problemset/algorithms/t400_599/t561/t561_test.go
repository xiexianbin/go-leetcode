package t561

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func ExampleArrayPairSum() {
	for _, nums := range [][]int{{1, 4, 3, 2}} {
		fmt.Println(arrayPairSum(nums))
	}

	// Output:
	// 4
}
