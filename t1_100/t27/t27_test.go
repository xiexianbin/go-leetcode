package t27

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	c := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if val == nums[i] {
			for true {
				if l-c-1 < 0 || c+i+1 >= l {
					break
				}
				if nums[l-c-1] != val {
					nums[i] = nums[l-c-1]
					nums[l-c-1] = val
					break
				}
				c++
			}
			c++
		}
		if c+i+1 >= l {
			break
		}
	}
	return l - c
}

func TestT27(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3

	//nums := []int{0,1,2,2,3,0,4,2}
	//val := 2

	//nums := []int{1}
	//val := 1

	l := removeElement(nums, val)
	fmt.Println(l)
	for i := 0; i < l; i++ {
		fmt.Print(nums[i])
	}
}
