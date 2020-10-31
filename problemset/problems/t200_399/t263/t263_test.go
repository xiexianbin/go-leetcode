package t

import (
	"fmt"
)

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for {
		if num == 1 {
			return true
		} else if num%2 == 0 {
			num = num / 2
		} else if num%3 == 0 {
			num = num / 3
		} else if num%5 == 0 {
			num = num / 5
		} else {
			return false
		}
	}
}

func ExampleIsUgly() {
	inputs := []int{6, 8, 14}
	for _, n := range inputs {
		fmt.Println(isUgly(n))
	}
	// Output:
	// true
	// true
	// false
}
