package t122

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	if len(prices) < 2 {
		return max
	}

	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			max += prices[i] - prices[i-1]
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
	// 7
	// 0
	// 0
}
