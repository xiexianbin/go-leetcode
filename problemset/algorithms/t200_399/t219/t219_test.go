package t219

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := indexMap[nums[i]]; ok {
			if i-j <= k {
				return true
			} else {
				indexMap[nums[i]] = i
			}
		} else {
			indexMap[nums[i]] = i
		}
	}
	return false
}

func ExampleContainsNearbyDuplicate() {
	numss := [][]int{{1, 2, 3, 1}, {1, 0, 1, 1}, {1, 2, 3, 1, 2, 3}}
	ks := []int{3, 1, 2}
	for i := 0; i < len(numss); i++ {
		fmt.Println(containsNearbyDuplicate(numss[i], ks[i]))
	}
	// Output:
	// true
	// true
	// false
}
