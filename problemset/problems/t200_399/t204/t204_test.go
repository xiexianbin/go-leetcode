package t204

import (
	"fmt"
)

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	count := 0
	primes := make([]bool, n)

	for i := 2; i < n; i++ {
		if primes[i] == true {
			continue
		}
		for j := i + i; j < n; j += i {
			primes[j] = true
		}
		count++
	}

	return count
}

func ExampleCountPrimes() {
	nums := []int{10, 0, 1}
	for _, n := range nums {
		fmt.Println(countPrimes(n))
	}

	// Output:
	// 4
	// 0
	// 0
}
