package t4

import (
	"fmt"
	"strconv"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	var result float64
	if l1 == 0 || l2 == 0 {
		nums := append(nums1, nums2...)
		if len(nums)%2 == 1 {
			m := (len(nums)) / 2
			result = float64(nums[m])
		} else {
			m := (len(nums)) / 2
			result = (float64(nums[m-1]) + float64(nums[m])) / 2
		}
	} else if (l1+l2)%2 == 1 {
		m := (l1 + l2) / 2
		i, j := 0, 0
		for {
			if i == l1 && j < l2 {
				m--
				result = float64(nums2[j])
				j++
			} else if i < l1 && j == l2 {
				m--
				result = float64(nums1[i])
				i++
			} else if nums1[i] > nums2[j] {
				m--
				result = float64(nums2[j])
				j++
			} else {
				m--
				result = float64(nums1[i])
				i++
			}
			if m < 0 {
				break
			}
		}
	} else {
		m := (l1+l2)/2 - 1
		i, j := 0, 0
		for {
			if i == l1 && j < l2 {
				m--
				result = float64(nums2[j])
				j++
			} else if i < l1 && j == l2 {
				m--
				result = float64(nums1[i])
				i++
			} else if nums1[i] > nums2[j] {
				m--
				result = float64(nums2[j])
				j++
			} else {
				m--
				result = float64(nums1[i])
				i++
			}
			if m < 0 {
				break
			}
		}
		if i == l1 && j < l2 {
			result = (result + float64(nums2[j])) / 2
		} else if i < l1 && j == l2 {
			result = (result + float64(nums1[i])) / 2
		} else if nums1[i] > nums2[j] {
			result = (result + float64(nums2[j])) / 2
		} else {
			result = (result + float64(nums1[i])) / 2
		}
	}

	ans, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(result)), 64)
	return ans
}

func ExampleFindMedianSortedArrays() {
	arr1 := [][]int{{1}, {2, 3}, {100000}, {100001}, {2, 2, 2}, {1, 3}, {1, 2}, {0, 0}, {}, {2}}
	arr2 := [][]int{{2, 3, 4}, {1}, {100001}, {100000}, {2, 2, 2, 2}, {2}, {3, 4}, {0, 0}, {1}, {}}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%.5f", findMedianSortedArrays(arr1[i], arr2[i]))
		fmt.Println()
	}

	// Output:
	// 2.50000
	// 2.00000
	// 100000.50000
	// 100000.50000
	// 2.00000
	// 2.00000
	// 2.50000
	// 0.00000
	// 1.00000
	// 2.00000
}
