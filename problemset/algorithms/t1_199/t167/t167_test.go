package t167

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	var m, n int
	isFind := false
	for m = 0; m < len(numbers); m++ {
		for n = 0; n < len(numbers); n++ {
			if m == n {
				continue
			}
			if numbers[m] + numbers[n] == target {
				isFind = true
				break
			}
			if numbers[m] + numbers[n] > target {
				break
			}
		}
		if isFind{
			break
		}
	}

	return []int{m + 1, n + 1}
}

func ExampleTwoSum() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 18))

	// Output:
	// [1 2]
	// [2 3]
}
