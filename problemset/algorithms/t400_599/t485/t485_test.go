package t485

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	maxCount, count := 0, 0

	for _, b := range nums {
		if b == 1 {
			count++
		} else if b == 0 {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}

	if count > maxCount {
		maxCount = count
	}

	return maxCount
}

func ExampleFindMaxConsecutiveOnes() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))

	// Output:
	// 3
}
