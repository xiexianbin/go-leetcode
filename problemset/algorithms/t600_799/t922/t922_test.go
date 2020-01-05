package t922

import (
	"fmt"
	"testing"
)

func sortArrayByParityII(A []int) []int {
	h := 0
	for {
		if h%2 == 0 && A[h]%2 == 0 {
			h++
		} else {
			for i := h + 1; i < len(A); i++ {
				if A[i]%2 == 0 {
					A[i], A[h] = A[h], A[i]
					h++
					break
				}
			}
		}
		if h%2 == 1 && A[h]%2 == 1 {
			h++
		} else {
			for i := h + 1; i < len(A); i++ {
				if A[i]%2 == 1 {
					A[i], A[h] = A[h], A[i]
					h++
					break
				}
			}
		}
		if h >= len(A) {
			break
		}
	}

	return A
}

func TestTest(t *testing.T) {
	as := [][]int{{2, 3, 1, 1, 4, 0, 0, 4, 3, 3}, {2, 3}, {4, 2, 5, 7}}
	for _, A := range as {
		fmt.Println(sortArrayByParityII(A))
	}
}
