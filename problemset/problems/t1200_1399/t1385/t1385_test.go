package t1385

import (
	"fmt"
	"math"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	count := 0
	for i := 0; i < len(arr1); i++ {
		flag := true
		for j := 0; j < len(arr2); j++ {
			if math.Abs(float64(arr1[i]-arr2[j])) <= math.Abs(float64(d)) {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}

	return count
}

func ExampleFindTheDistanceValue() {
	arr1s := [][]int{{4, 5, 8}, {1, 4, 2, 3}, {2, 1, 100, 3}}
	arr2s := [][]int{{10, 9, 1, 8}, {-4, -3, 6, 10, 20, 30}, {-5, -2, 10, -3, 7}}
	ds := []int{2, 3, 6}

	for i := 0; i < len(arr1s); i++ {
		fmt.Println(findTheDistanceValue(arr1s[i], arr2s[i], ds[i]))
	}

	// Output:
	// 2
	// 2
	// 1
}
