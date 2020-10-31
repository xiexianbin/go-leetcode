package t39

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func ExampleMajorityElement() {
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))

	// Output:
	// 2
}
