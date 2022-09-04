package t74

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	i := len(matrix) - 1
	j := 0
	for {
		if i < 0 || j > len(matrix[0])-1 {
			break
		}
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][0] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			break
		}
	}

	return false
}

func ExampleSearchMatrix() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))

	// Output:
	// true
}
