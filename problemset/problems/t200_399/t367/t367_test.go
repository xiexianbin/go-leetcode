package t

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	x := 1
	for num > 0 {
		num -= x
		x += 2
	}
	if num == 0 {
		return true
	}
	return false
}

func ExampleIsPerfectSquare() {
	arr := []int{16, 14, 1, 2147483647}
	for i := 0; i < len(arr); i++ {
		fmt.Println(isPerfectSquare(arr[i]))
	}

	// Output:
	// true
	// false
	// true
	// false
}
