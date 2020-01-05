package t350

import (
	"fmt"
	"testing"
)

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for {
		i := nums1[0]
		for j := 0; j < len(nums2); j++ {
			if i == nums2[j] {
				result = append(result, i)
				nums2 = append(nums2[:j], nums2[j+1:]...)
				break
			}
		}

		nums1 = nums1[1:]
		if len(nums1) == 0 {
			break
		}
	}

	return result
}

func TestTest(t *testing.T) {
	num1 := [][]int{{1, 2, 2, 1}, {4, 9, 5}, {3, 1, 2}}
	num2 := [][]int{{2, 2}, {9, 4, 9, 8, 4}, {1, 1}}

	for i, v := range num1 {
		fmt.Println(intersect(v, num2[i]))
	}
}
