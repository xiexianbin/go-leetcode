package t905

import (
	"fmt"
	"testing"
)

func sortArrayByParity(A []int) []int {
	if len(A) < 2 {
		return A
	}

	h, t := 0, len(A)-1
	for {
		if A[h]%2 == 0 {
			h++
		}
		if A[t]%2 == 1 {
			t--
		}
		if h > t {
			break
		}
		if A[h]%2 == 1 && A[t]%2 == 0 {
			A[h], A[t] = A[t], A[h]
			h++
			t--
		}
	}

	return A
}

func TestTest(t *testing.T) {
	as := [][]int{{0}, {0, 2}, {3, 1, 2, 4}}
	for _, A := range as {
		fmt.Println(sortArrayByParity(A))
	}
}
