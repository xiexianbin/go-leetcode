package t1460

import (
	"fmt"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}

	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < len(target); i++ {
		if target[i] != arr[i] {
			return false
		}
	}

	return true
}

func ExampleCanBeEqual() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))

	// Output:
	// true
}
