package lcp1

import (
	"fmt"
)

func game(guess []int, answer []int) int {
	var count int
	for i, v := range guess {
		if v == answer[i] {
			count++
		}
	}
	return count
}

func ExampleGame() {
	guess := [][]int{[]int{1, 2, 3}, []int{2, 2, 3}}
	answer := [][]int{[]int{1, 2, 3}, []int{3, 2, 1}}

	for i, v := range guess {
		fmt.Println(game(v, answer[i]))
	}

	// Output:
	// 3
	// 1
}
