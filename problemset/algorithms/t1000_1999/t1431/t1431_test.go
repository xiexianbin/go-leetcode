package t1431

import (
	"fmt"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, len(candies))
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			ans[i] = true
		}
	}

	return ans
}

func ExampleKidsWithCandies() {
	candiess := [][]int{{2, 3, 5, 1, 3}, {4, 2, 1, 1, 2}, {12, 1, 12}}
	extraCandiess := []int{3, 1, 10}
	for i := 0; i < len(candiess); i++ {
		fmt.Println(kidsWithCandies(candiess[i], extraCandiess[i]))
	}

	// Output:
	// [true true true false true]
	// [true false false false false]
	// [true false true]
}
