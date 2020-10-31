package t119

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	result := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}
	return result[rowIndex]
}

func getRow2(rowIndex int) []int {
	result := []int{1}
	for i := 0; i < rowIndex; i++ {
		result = append(result, 1)
		for j := i; j > 0; j-- {
			result[j] += result[j-1]
		}
	}
	return result
}

func ExampleGenerate() {
	numRows := 3
	fmt.Println(getRow(numRows))
	fmt.Println(getRow2(numRows))

	// Output:
	//[1 3 3 1]
	//[1 3 3 1]
}
