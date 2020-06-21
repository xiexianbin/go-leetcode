package t1300

import (
	"fmt"
	"math"
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)

	index, arrIndex := 0, 0
	sum := arr[index]
	lessFlag := true

	diff := float64(target - arr[index]*len(arr))
	if diff < 0 {
		lessFlag = false
	}
	value := arr[index]
	for ; value<=arr[len(arr)-1]; value++{
		if value > arr[arrIndex]{
			arrIndex++
			sum += arr[arrIndex]
		}

		for j := 1; j < len(arr); j++ {
			t := sum + arr[j]*(len(arr)-j)
			if diff > math.Abs(float64(target-t)) {
				index = j
				diff = math.Abs(float64(target - t))
				break
			}
			sum += arr[j]
		}
	}

	return value
}

func ExampleT() {
	arrs := [][]int{{4, 9, 3}, {2, 3, 5}, {60864, 25176, 27249, 21296, 20204}}
	targets := []int{10, 10, 56803}
	for i := 0; i < len(arrs); i++ {
		fmt.Println(findBestValue(arrs[i], targets[i]))
	}

	// Output:
	// 3
	// 5
	// 11361
}
