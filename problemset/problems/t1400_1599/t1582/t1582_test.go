package t1582

import (
	"fmt"
)

func numSpecial(mat [][]int) int {
	count := 0
	rows := make([]int, len(mat))
	cols := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if rows[i] == 1 && cols[j] == 1 && mat[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

func ExampleT() {
	fmt.Println(numSpecial([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}))

	// Output:
	// 1
}
