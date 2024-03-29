package t189

import (
	"fmt"
)

func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}

	l := len(nums)
	if k >= l {
		k = k % l
	}
	result := append(nums[l-k:], nums[:l-k]...)
	nums = append(nums[:0], result...)
}

func ExampleRotate() {
	nums := [][]int{{1, 2, 3, 4, 5, 6, 7}, {-1, -100, 3, 99}, {-1}, {1, 2}}
	ks := []int{3, 2, 2, 3}
	for i, num := range nums {
		rotate(num, ks[i])
		fmt.Println(num)
	}

	// Output:
	//[5 6 7 1 2 3 4]
	//[3 99 -1 -100]
	//[-1]
	//[2 1]
}
