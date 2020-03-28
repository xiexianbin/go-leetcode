package t66

import (
	"fmt"
	"testing"
)

func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return []int{1}
	}

	digits[l-1] = digits[l-1] + 1
	for i := l - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i-1] = digits[i-1] + digits[i]/10
			digits[i] = digits[i] % 10
		}
	}

	t := digits[0]
	if t > 9 {
		digits[0] = t % 10
		return append([]int{t / 10}, digits[0:]...)
	}
	return digits
}

func TestPlusOne(t *testing.T) {
	//digits := []int{1, 2, 3}
	//digits := []int{9}
	digits := []int{9, 9}
	//digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}

	fmt.Println(plusOne(digits))

	// Output:
}
