package t779

import (
	"fmt"
)

func kthGrammar(n int, k int) int {
	if k == 1 {
		return 0
	}

	r := kthGrammar(n-1, (k+1)/2)
	if k%2 == 1 {
		return r
	} else {
		return 1 - r
	}
}

func ExampleKthGrammar() {
	d := map[int]int{1: 1, 2: 1, 4: 5}
	for k, v := range d {
		fmt.Println(kthGrammar(k, v))
	}

	// Output:
	// 0
	// 0
	// 1
}
