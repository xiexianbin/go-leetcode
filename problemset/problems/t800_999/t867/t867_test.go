package t867

import (
	"fmt"
)

func transpose(A [][]int) [][]int {
	B := [][]int{}
	L := len(A)
	W := len(A[0])
	for i := 0; i < W; i++ {
		R := []int{}
		for j := 0; j < L; j++ {
			R = append(R, A[j][i])
		}
		B = append(B, R)
	}

	return B
}

func PrintR(B [][]int) {
	for i := 0; i < len(B); i++ {
		l := B[i]
		for j := 0; j < len(l); j++ {
			fmt.Printf("%v", l[j])
			if j < len(l)-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func ExampleTranspose() {
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	PrintR(transpose(A))

	A = [][]int{{1, 2, 3}, {4, 5, 6}}
	PrintR(transpose(A))

	// Output:
	// 1 4 7
	// 2 5 8
	// 3 6 9
	// 1 4
	// 2 5
	// 3 6
}
