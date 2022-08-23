package t191

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	count := 0

	for {
		if num%2 == 1 {
			count++
		}
		num = num / 2
		if num == 1 {
			count++
			break
		} else if num == 0 {
			break
		}
	}

	return count
}

func ExampleHammingWeight() {
	for _, num := range []uint32{00000000000000000000000000001011, 00000000000000000000000010000000} {
		fmt.Println(hammingWeight(num))
	}

	// Output:
	// 3
	// 1
}
