package t977

import (
	"fmt"
	"sort"
)

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}
	sort.Ints(A)
	return A
}

func ExampleSortedSquares() {
	for _, A := range [][]int{{-4, -1, 0, 3, 10}, {-7, -3, 2, 3, 11}} {
		fmt.Println(sortedSquares(A))
	}

	// Output:
	//[0 1 9 16 100]
	//[4 9 9 49 121]
}
