package t70

import (
	"fmt"
	"testing"
)

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	cache := make([]int, n)
	cache[0], cache[1] = 1, 2

	for i := 2; i < n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n-1]
}

func TestTest(t *testing.T) {
	for _, n := range []int{1, 2, 3} {
		fmt.Println(climbStairs(n))
	}
}
