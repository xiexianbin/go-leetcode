package t118

import (
	"fmt"
	"testing"
)

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}

	return result
}

func TestGenerate(t *testing.T) {
	numRows := 5
	result := generate(numRows)
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Print(result[i][j])
		}
		fmt.Println()
	}
}
