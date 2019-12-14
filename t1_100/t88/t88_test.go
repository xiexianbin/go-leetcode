package t88

import (
	"fmt"
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	c := m + n
	nums1 = nums1[:c]
	for i := 0; i < c; i++ {
		for j := 0; j < c; j++ {
			if nums1[i] < nums1[j] {
				t := nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = t
			}
		}
	}
}

func TestMerge(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, 3)
	fmt.Println(num1)
}
