package t1281

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for {
		if n == 0 {
			break
		}
		x := n % 10
		n = n / 10

		sum += x
		product *= x
	}
	return product - sum
}

func ExampleSubtractProductAndSum() {
	inputs := []int{234, 4421}
	for _, n := range inputs {
		fmt.Println(subtractProductAndSum(n))
	}
	// Output:
	// 15
	// 21
}
