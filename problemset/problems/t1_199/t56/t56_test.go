package t56

import (
	"fmt"
)

func merge(intervals [][]int) [][]int {
	l := len(intervals)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	ans := [][]int{intervals[0]}
	for i := 1; i < l; i++ {
		if ans[len(ans)-1][1] >= intervals[i][0] {
			if ans[len(ans)-1][1] < intervals[i][1] {
				ans[len(ans)-1][1] = intervals[i][1]
			}
		} else {
			ans = append(ans, intervals[i])
		}
	}

	return ans
}

func ExampleMerge() {
	fmt.Println(merge([][]int{{1, 4}, {5, 6}}))
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))

	// Output:
	// [[1 4] [5 6]]
	// [[1 6] [8 10] [15 18]]
}
