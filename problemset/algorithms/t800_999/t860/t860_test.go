package t860

import (
	"fmt"
)

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		if b == 5 {
			five++
		} else if b == 10 {
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		} else {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five > 2 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}

func ExampleLemonadeChange() {
	billss := [][]int{{5, 5, 5, 10, 20}, {5, 5, 10}, {10, 10}, {5, 5, 10, 10, 20}}
	for _, bills := range billss {
		fmt.Println(lemonadeChange(bills))
	}

	// Output:
	// true
	// true
	// false
	// false
}
