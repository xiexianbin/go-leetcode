package t347

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	h := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := h[nums[i]]; ok {
			h[nums[i]] = v + 1
		} else {
			h[nums[i]] = 1
		}
	}

	t := []int{}
	reverseH := make(map[int]int)
	for key, value := range h {
		t = append(t, value)
		reverseH[value] = key
	}
	sort.Ints(t)

	minNum := t[len(t)-k]

	ans := make([]int, 0, k)
	for key, value := range h {
		if value >= minNum {
			ans = append(ans, key)
		}
	}
	return ans
}

func ExampleTopKFrequent() {
	r1 := topKFrequent([]int{1, 2}, 2)
	sort.Ints(r1)
	fmt.Println(r1)
	r2 := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	sort.Ints(r2)
	fmt.Println(r2)

	// Output:
	// [1 2]
	// [1 2]
}
