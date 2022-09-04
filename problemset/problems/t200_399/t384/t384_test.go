package t384

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums     []int
	original []int
}

func Constructor(nums []int) Solution {
	return Solution{
		original: append([]int{}, nums...),
		nums:     nums,
	}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.original)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	result := make([]int, 0, len(this.nums))
	for {
		if len(this.nums) == 0 {
			break
		}
		i := rand.Intn(len(this.nums))
		result = append(result, this.nums[i])
		this.nums = append(this.nums[:i], this.nums[i+1:]...)
	}
	this.nums = result
	return result
}

func ExampleConstructor() {
	/**
	 * Your Solution object will be instantiated and called as such:
	 */
	obj := Constructor([]int{1, 3, 2})
	fmt.Println(obj.Reset())
	for i := 0; i < 100; i++ {
		obj.Reset()
		obj.Shuffle()
	}

	// Output:
	// [1 3 2]
}
