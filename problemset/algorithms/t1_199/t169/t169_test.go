package t169

import (
	"fmt"
	"sort"
	"testing"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func TestTest(t *testing.T) {
	for _, nums := range [][]int{{3,2,3}, {2,2,1,1,1,2,2,}} {
		fmt.Println(majorityElement(nums))
	}

	// Output:
}
