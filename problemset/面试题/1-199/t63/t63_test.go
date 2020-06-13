package t63

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	if len(prices) < 2 {
		return max
	}

	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				diff := prices[j] - prices[i]
				if diff > max {
					max = diff
				}
			}
		}
	}
	return max
}

func ExampleMaxProfit() {
	stocks := [][]int{{7, 1, 5, 3, 6, 4}, {7, 6, 4, 3, 1}, {}}

	for _, s := range stocks {
		fmt.Println(maxProfit(s))
	}
	// Output:
	// 5
	// 0
	// 0
}
