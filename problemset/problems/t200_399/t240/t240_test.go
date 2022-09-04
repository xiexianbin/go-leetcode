package t240

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	i := 0
	j := len(matrix[0]) - 1
	for {
		if i > len(matrix)-1 || j < 0 {
			break
		}
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			break
		}
	}

	return false
}

func ExampleSearchMatrix() {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 40))

	// Output:
	// false
}
