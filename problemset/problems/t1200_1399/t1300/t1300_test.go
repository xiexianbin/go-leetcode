package t1300

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	max := arr[n-1]
	ans, diff := 0, target
	for i := 1; i <= max; i++ {
		index := sort.SearchInts(arr, i)
		if index < 0 {
			index = -index - 1
		}

		cur := prefix[index] + (n-index)*i
		if abs(cur-target) < diff {
			ans = i
			diff = abs(cur - target)
		}
	}

	return ans
}

func ExampleFindBestValue() {
	arrs := [][]int{{4, 9, 3}, {2, 3, 5}, {60864, 25176, 27249, 21296, 20204}}
	targets := []int{10, 10, 56803}
	for i := 0; i < len(arrs); i++ {
		fmt.Println(findBestValue(arrs[i], targets[i]))
	}

	// Output:
	// 3
	// 5
	// 11361
}
