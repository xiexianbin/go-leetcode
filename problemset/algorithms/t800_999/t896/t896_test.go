package t896

import (
	"fmt"
	"testing"
)

func isMonotonic(A []int) bool {
	if len(A) < 2 {
		return true
	}

	increaseFlag := 0

	for i := 1; i < len(A); i++ {
		if increaseFlag == 0 {
			if A[i-1] > A[i] {
				increaseFlag = -1
			} else if A[i-1] < A[i] {
				increaseFlag = 1
			}
		}

		if increaseFlag == 1 && A[i-1] > A[i] {
			return false
		}
		if increaseFlag == -1 && A[i-1] < A[i] {
			return false
		}
	}

	return true
}

func TestTest(t *testing.T) {
	for _, A := range [][]int{{1, 1, 0}, {1, 1, 1}, {1, 2, 2, 3}, {6, 5, 4, 4}, {1, 3, 2}, {1, 2, 4, 5}, {1, 1, 1}} {
		fmt.Println(isMonotonic(A))
	}
}
