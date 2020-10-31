package t10_01

import (
	"fmt"
)

func merge(A []int, m int, B []int, n int) {
	for {
		if n == 0 {
			for {
				A[m+n-1] = A[m-1]
				m--
				if m == 0 {
					break
				}
			}
			break
		}
		if m == 0 {
			for {
				A[m+n-1] = B[n-1]
				n--
				if n == 0 {
					break
				}
			}
			break
		}
		if A[m-1] > B[n-1] {
			A[m+n-1] = A[m-1]
			m--
		} else {
			A[m+n-1] = B[n-1]
			n--
		}
	}
}

func ExampleMerge() {
	A := []int{1, 2, 3, 0, 0, 0}
	m := 3
	B := []int{2, 5, 6}
	n := 3

	merge(A, m, B, n)
	fmt.Println(A)

	A = []int{2, 2, 3, 0, 0, 0}
	m = 3
	B = []int{1, 5, 6}
	n = 3

	merge(A, m, B, n)
	fmt.Println(A)

	A = []int{1}
	m = 1
	B = []int{}
	n = 0

	merge(A, m, B, n)
	fmt.Println(A)

	// Output:
	// [1 2 2 3 5 6]
	// [1 2 2 3 5 6]
	// [1]
}
