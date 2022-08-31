package t977

import (
	"fmt"
)

//func sortedSquares(A []int) []int {
//	for i := 0; i < len(A); i++ {
//		A[i] = A[i] * A[i]
//	}
//	sort.Ints(A)
//	return A
//}

func sortedSquares(nums []int) []int {
	lastNagativeIndex := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			lastNagativeIndex = i
		} else {
			break
		}
	}

	ans := make([]int, 0, n)
	i, j := lastNagativeIndex, lastNagativeIndex+1
	for ; 0 <= i || j < n; {
		if i < 0 && j < n {
			ans = append(ans, nums[j]*nums[j])
			j++
			continue
		}
		if j >= n && i >= 0 {
			ans = append(ans, nums[i]*nums[i])
			i--
			continue
		}
		di := nums[i] * nums[i]
		dj := nums[j] * nums[j]
		if di > dj {
			ans = append(ans, dj)
			j++
		} else {
			ans = append(ans, di)
			i--
		}
	}

	return ans
}

func ExampleSortedSquares() {
	for _, A := range [][]int{{-4, -1, 0, 3, 10}, {-7, -3, 2, 3, 11}} {
		fmt.Println(sortedSquares(A))
	}

	// Output:
	//[0 1 9 16 100]
	//[4 9 9 49 121]
}
