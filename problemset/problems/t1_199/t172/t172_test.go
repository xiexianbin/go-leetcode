package t172

import (
	"fmt"
)

func trailingZeroes(n int) int {
	fives := 0
	for {
		fives += n / 5
		n = n / 5
		if n == 0 {
			break
		}
	}
	return fives
}

func ExampleTrailingZeroes() {
	inputs := []int{3, 5}
	for _, n := range inputs {
		fmt.Println(trailingZeroes(n))
	}

	// Output:
	// 0
	// 1
}
