package t1470

import (
	"fmt"
)

func shuffle(nums []int, n int) []int {
	t := make([]int, 0, n-1)
	for i := 1; i < n; i++ {
		t = append(t, nums[i])
	}

	j := 0
	for i := 1; i < 2*n-1; i++ {
		if i%2 == 1 {
			nums[i] = nums[n+i/2]
		} else {
			nums[i] = t[j]
			j++
		}
	}

	return nums
}

func ExampleShuffle() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))

	// Output:
	// [2 3 5 4 1 7]
}
