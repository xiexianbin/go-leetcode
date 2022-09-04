package t461

import (
	"fmt"
)

func hammingDistance(x int, y int) int {
	ans := 0
	for {
		if x == 0 && y == 0 {
			break
		}
		if (x%2 == 0 && y%2 == 1) || (x%2 == 1 && y%2 == 0) {
			ans++
		}
		x = x / 2
		y = y / 2
	}
	return ans
}

func ExampleHammingDistance() {
	fmt.Println(hammingDistance(1, 4))

	// Output:
	// 2
}
