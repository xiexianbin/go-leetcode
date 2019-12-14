package t68

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
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
