package t283

import (
	"fmt"
)

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	count := 0
	for _, num := range nums {
		if num == 0 {
			count++
		}
	}

	for i := 0; i < count; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] == 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func ExampleMoveZeroes() {
	for _, nums := range [][]int{{0, 1, 0, 3, 12}} {
		moveZeroes(nums)
		fmt.Println(nums)
	}

	// Output:
	//[1 3 12 0 0]
}
