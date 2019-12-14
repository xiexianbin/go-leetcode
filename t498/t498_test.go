package t498

import (
	"fmt"
	"testing"
)

func findDiagonalOrder(matrix [][]int) []int {
	result := []int{}
	mx := len(matrix)
	my := len(matrix[0])
	maxIndexSum := mx + my - 2
	for indexSum := 0; indexSum <= maxIndexSum; indexSum++ {
		if indexSum % 2 == 0 {
			y := 0
			x := indexSum - y
			if x >= mx {
				x = mx -1
				y = indexSum -x
			}
			for {
				if y >= my || x < 0 {
					break
				}
				result = append(result, matrix[x][y])
				x--
				y++
			}
		} else {
			x := 0
			y := indexSum - x
			if y >= my {
				y = my -1
				x = indexSum -y
			}
			for {
				if x >= mx || y < 0 {
					break
				}
				result = append(result, matrix[x][y])
				x++
				y--
			}
		}
	}
	return result
}

func TestTest(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(findDiagonalOrder(matrix))

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(findDiagonalOrder(matrix))
}
