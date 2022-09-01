package t48

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func ExampleRotate() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)

	// Output:
	// [[7 4 1] [8 5 2] [9 6 3]]
}
