package t33

import "fmt"

func search(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}

		m := (left + right) / 2
		if nums[m] == target {
			return m
		}
		if nums[left] < nums[m] { // left order
			if nums[left] < target && target < nums[m] {
				right = m - 1
			} else {
				left = m + 1
			}
		} else { // right order
			if nums[m] < target && target < nums[right] {
				left = m + 1
			} else {
				right = m - 1
			}
		}
	}

	return -1
}

func ExampleSearch() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

	// Output:
	// 4
}
