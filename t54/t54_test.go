package t54

import (
	"fmt"
	"testing"
)

func spiralOrder(matrix [][]int) []int {
	mx := len(matrix)
	if mx == 0 {
		return []int{}
	}
	my := len(matrix[0])
	result := []int{}
	cminx, cminy, cmaxx, cmaxy := 0, 0, mx, my
	direct := 0
	for {
		if direct%4 == 0 {
			items := matrix[cminx]
			for y := cminy; y < cmaxy; y++ {
				result = append(result, items[y])
			}
			cminx++
		} else if direct%4 == 1 {
			for x := cminx; x < cmaxx; x++ {
				items := matrix[x]
				result = append(result, items[cmaxy-1])
			}
			cmaxy--
		} else if direct%4 == 2 {
			items := matrix[cmaxx-1]
			for y := cmaxy - 1; y >= cminy; y-- {
				result = append(result, items[y])
			}
			cmaxx--
		} else if direct%4 == 3 {
			for x := cmaxx - 1; x >= cminx; x-- {
				items := matrix[x]
				result = append(result, items[cminy])
			}
			cminy++
		}
		if len(result) == mx*my {
			break
		}
		direct++
	}
	return result
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}
