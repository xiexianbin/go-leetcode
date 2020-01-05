package t977

import (
	"fmt"
	"sort"
	"testing"
)

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}
	sort.Ints(A)
	return A
}

func TestTest(t *testing.T) {
	for _, A := range [][]int{{-4, -1, 0, 3, 10}, {-7, -3, 2, 3, 11}} {
		fmt.Println(sortedSquares(A))
	}
}
