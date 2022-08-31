package t6160

import (
	"fmt"
	"sort"
)

func subSummary(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

// return sub array length
func queryOne(nums []int, query int) int {
	result := -1
	l := len(nums)
	for i := 0; i < l; i++ {
		subSum := subSummary(nums[0 : i+1])
		if subSum <= query {
			if i+1 > result {
				result = i + 1
			}
		} else {
			break
		}
	}
	return result
}

func answerQueries(nums []int, queries []int) []int {
	answer := []int{}
	sort.Ints(nums)
	for i := 0; i < len(queries); i++ {
		result := queryOne(nums, queries[i])
		if result != -1 {
			answer = append(answer, result)
		} else {
			answer = append(answer, 0)
		}
	}

	return answer
}

func ExampleAnswerQueries() {
	fmt.Println(answerQueries([]int{624082}, []int{972985, 564269, 607119, 693641, 787608, 46517, 500857, 140097}))
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))

	// Output:
	// [1 0 0 1 1 0 0 0]
	// [2 3 4]
}
