package t202

import (
	"fmt"
)

func isHappy(n int) bool {
	for {
		sum := 0
		for {
			sum += (n % 10) * (n % 10)
			n = n / 10
			if n < 10 {
				sum += n * n
				n = sum
				break
			}
		}
		if n == 1 || n == 7 {
			return true
		}
		if n > 1 && n < 10 {
			return false
		}
	}
}

func ExampleIsHappy() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(20))
	fmt.Println(isHappy(21))
	fmt.Println(isHappy(1111111))

	// Output:
	// true
	// false
	// false
	// true
}
