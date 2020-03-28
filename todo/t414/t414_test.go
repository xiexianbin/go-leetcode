package t414

import (
	"fmt"
)

func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	first, second, third := ^1<<32, ^1<<32, ^1<<32
	for i := 0; i < len(nums); i++ {
		if first == nums[i] || second == nums[i] || third == nums[i] {
			continue
		}
		if nums[i] > first {
			third = second
			second = first
			first = nums[i]
		} else if nums[i] > second {
			third = second
			second = nums[i]
		} else if nums[i] > third {
			third = nums[i]
		}
	}

	if third == ^1<<32 {
		return first
	}
	return third
}


func ExampleThirdMax(){
	l := [][]int{{3, 2, 1}, {1, 2}, {2, 2, 3, 1}, {1,2,-2147483648}}
	for _, sL := range l {
		fmt.Println(thirdMax(sL))
	}

	// Output:
	//1
	//2
	//1
	//-2147483648
}
