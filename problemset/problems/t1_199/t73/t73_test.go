package t73

import "fmt"

func setZeroes(matrix [][]int) {
	line := make([]bool, len(matrix))
	row := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				line[i] = true
				row[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if line[i] == true || row[j] == true {
				matrix[i][j] = 0
			}
		}
	}
}

func ExampleSetZeroes() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)

	// Output:
	// [[1 0 1] [0 0 0] [1 0 1]]
}
