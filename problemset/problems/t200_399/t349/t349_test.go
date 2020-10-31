package t349

import (
	"fmt"
)

func duplicateRemoval(nums []int) []int {
	result := []int{}
	if len(nums) > 0 {
		result = append(result, nums[0])
	}
	for i := 1; i < len(nums); i++ {
		j := 0
		for j < len(result) {
			if nums[i] == result[j] {
				break
			}
			j++
		}
		if j == len(result) {
			result = append(result, nums[i])
		}
	}
	return result
}

func intersection(nums1 []int, nums2 []int) []int {
	nums1 = duplicateRemoval(nums1)
	nums2 = duplicateRemoval(nums2)
	result := []int{}

	if len(nums1) <= len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				result = append(result, nums1[i])
				break
			}
		}
	}

	return result
}

func ExampleIntersection() {
	num1 := [][]int{{1, 2, 2, 1}, {4, 9, 5}}
	num2 := [][]int{{2, 2}, {9, 4, 9, 8, 4}}

	for i, v := range num1 {
		fmt.Println(intersection(v, num2[i]))
	}

	// Output:
	//[2]
	//[9 4]
}
