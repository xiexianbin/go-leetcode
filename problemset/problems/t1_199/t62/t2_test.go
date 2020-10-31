package t62

import (
	"fmt"
)

func f(n, m int) int {
	if n == 1 {
		return 0
	}
	x := f(n-1, m)
	return (m + x) % n
}

func lastRemaining(n int, m int) int {
	return f(n, m)
}

func ExampleLastRemaining() {
	test := [][]int{{5, 3}, {10, 17}}
	for _, t := range test {
		fmt.Println(lastRemaining(t[0], t[1]))
	}

	// Output:
	// 3
	// 2
}
