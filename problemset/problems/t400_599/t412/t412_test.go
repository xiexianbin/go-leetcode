package t412

import (
	"fmt"
)

func fizzBuzz(n int) []string {
	ans := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, fmt.Sprintf("%d", i))
		}
	}
	return ans
}

func ExampleFizzBuzz() {
	fmt.Println(fizzBuzz(5))

	// Output:
	// [1 2 Fizz 4 Buzz]
}
