package t6167

import (
	"fmt"
)

func checkDistances(s string, distance []int) bool {
	cache := map[byte]int{}
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if firstIndex, ok := cache[index]; ok {
			if i-firstIndex-1 != distance[index] {
				return false
			} else {
				distance[index] = 0
			}
		} else {
			cache[index] = i
		}
	}

	return true
}

func ExampleT() {
	fmt.Println(checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))

	// Output:
	// true
}
