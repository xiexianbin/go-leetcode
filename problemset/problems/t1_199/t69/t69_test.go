package t69

import (
	"fmt"
)

func mySqrt(x int) int {
	l, r := 0, x
	ans := 0
	for l < r {
		mid := (l + r + 1) >> 1
		if mid * mid <= x {
			ans = mid
			l = mid
		} else {
			r = mid - 1
		}
	}
	return ans
}

func ExampleMySqrt() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(5))
	fmt.Println(mySqrt(0))

	// Output:
	// 2
	// 2
	// 2
	// 0
}
