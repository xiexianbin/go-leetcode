package t1137

import (
	"fmt"
	"testing"
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

func TestTest(t *testing.T) {
	for _, n := range []int{4, 25} {
		fmt.Println(tribonacci(n))
	}
}
