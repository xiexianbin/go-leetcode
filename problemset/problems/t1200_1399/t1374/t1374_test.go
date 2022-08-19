package t1374

import (
	"fmt"
	"strings"
)

func generateTheString(n int) string {
	if n%2 == 1 {
		return strings.Repeat("a", n)
	}
	return strings.Repeat("a", n-1) + "b"
}

func ExampleGenerateTheString() {
	nums := []int{4, 2}
	for _, num := range nums {
		fmt.Println(generateTheString(num))
	}

	// Output:
	// aaab
	// ab
}
