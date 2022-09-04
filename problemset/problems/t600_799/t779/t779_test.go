package t779

import (
	"fmt"
	"sort"
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
	r := []int{}
	for k, v := range d {
		r = append(r, kthGrammar(k, v))
	}
	sort.Ints(r)
	fmt.Println(r)

	// Output:
	// [0 0 1]
}
