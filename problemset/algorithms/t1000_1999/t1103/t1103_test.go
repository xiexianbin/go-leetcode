package t1103

import (
	"fmt"
)

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)

	count := 1
	ans[0] = 1
	candies -= 1
	for {
		if candies <= count {
			ans[count%num_people] += candies
			break
		}
		count++
		ans[(count-1)%num_people] += count
		candies -= count
	}

	return ans
}

func ExampleDistributeCandies() {
	fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))

	// Output:
	// [1 2 3 1]
	// [5 2 3]
}
