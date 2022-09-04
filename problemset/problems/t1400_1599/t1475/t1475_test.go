package t1475

import "fmt"

func finalPrices(prices []int) []int {
	result := []int{}
	l := len(prices)
	for i := 0; i < l; i++ {
		flag := false
		for j := i + 1; j < l; j++ {
			if prices[j] <= prices[i] {
				flag = true
				result = append(result, prices[i]-prices[j])
				break
			}
		}

		if flag == false {
			result = append(result, prices[i])
		}
	}

	return result
}

func ExampleFinalPrices() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))

	// Output:
	// [4 2 4 2 3]
}
