package t1137

import (
	"fmt"
)

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}

	cache := make([]int, n+1)
	cache[0], cache[1], cache[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2] + cache[i-3]
	}
	return cache[n]
}

func ExampleTribonacci() {
	for _, n := range []int{4, 25} {
		fmt.Println(tribonacci(n))
	}

	// Output:
	//4
	//1389537
}
