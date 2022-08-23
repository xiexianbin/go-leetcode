package t278

import (
	"fmt"
)

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 */
func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	left, right := 1, n
	version := 1
	callResult := make(map[int]bool)
	for left <= right {
		mid := (left + right) / 2
		if isBadVersion(mid) == false {
			callResult[mid] = false
			if isBad, ok := callResult[mid+1]; ok && isBad == true {
				return mid + 1
			}
			left = mid + 1
		} else {
			callResult[mid] = true
			if isBad, ok := callResult[mid-1]; ok && isBad == false {
				return mid
			}
			right = mid - 1
		}

	}
	return version
}

func ExampleFirstBadVersion() {
	fmt.Println(firstBadVersion(5))

	// Output:
	// 4
}
