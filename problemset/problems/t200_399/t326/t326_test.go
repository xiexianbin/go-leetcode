package t326

import (
	"fmt"
)

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			break
		}
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}

	return true
}

func ExampleT() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45))

	// Output:
	// true
	// false
	// true
	// false
}
