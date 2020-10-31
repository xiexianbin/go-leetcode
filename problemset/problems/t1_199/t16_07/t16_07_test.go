package t16_07

import (
	"fmt"
	"sort"
)

func maximum(a int, b int) int {
	arr := []int{a, b}
	sort.Ints(arr)
	return arr[1:][0]
}

func ExampleMaximum() {
	fmt.Println(maximum(1,2))

	// Output:
	// 2
}
