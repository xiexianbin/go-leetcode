package t75

import "fmt"

func sortColors(nums []int) {
	pr, pb, l := 0, len(nums)-1, len(nums)
	for i := 0; i < l; i++ {
		if i > pb {
			break
		}
		if nums[i] == 0 {
			nums[i], nums[pr] = nums[pr], nums[i]
			pr++
		} else if nums[i] == 2 {
			nums[i], nums[pb] = nums[pb], nums[i]
			pb--
			i--
		}
	}
}

func ExampleSortColors() {
	numss := [][]int{ {1, 0, 2}, {2, 0, 2, 1, 1, 0}, {2, 2, 1, 0, 1, 1}, {1, 2, 0, 1}}
	for _, nums := range numss {
		sortColors(nums)
		fmt.Println(nums)
	}

	// Output:
	// [0 1 2]
	// [0 0 1 1 2 2]
	// [0 1 1 1 2 2]
	// [0 1 1 2]
}
