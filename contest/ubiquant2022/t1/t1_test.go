// https://leetcode.cn/contest/ubiquant2022/
package t1

import "fmt"

func numberOfPairs(nums []int) int {
	// mirror num
	change := func(num int) int {
		result := 0
		for {
			if num < 10 {
				result = result*10 + num
				break
			} else {
				result = result*10 + num%10
				num = num / 10
			}
		}
		return result
	}

	// per nums count
	numsCount := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if c, ok := numsCount[nums[i]]; ok {
			numsCount[nums[i]] = c + 1
		} else {
			numsCount[nums[i]] = 1
		}
	}

	// diff nums
	numsDiff := make([]int, 0, len(numsCount))
	numsMirror := make(map[int]int)
	for k, _ := range numsCount {
		numsDiff = append(numsDiff, k)
		if _, ok := numsMirror[k]; ok == false {
			numsMirror[k] = change(k)
		}
	}

	l := len(numsDiff)

	ans := 0
	for i := 0; i < l; i++ {
		numi := numsDiff[i]
		for j := i + 1; j < l; j++ {
			numj := numsDiff[j]
			if numi+numsMirror[numj] == numsMirror[numi]+numj {
				ans += numsCount[numi] * numsCount[numj]
			}
		}
	}

	// dumple num count
	for k, v := range numsCount {
		if k+numsMirror[k] == numsMirror[k]+k {
			ans += v * (v - 1) / 2
		}
	}

	max := 1000000007
	if ans > max {
		return ans % max
	}

	return ans
}

func ExampleNumberOfPairs() {
	fmt.Println(numberOfPairs([]int{10, 22, 9, 58, 87, 80, 14, 47, 7, 11, 3, 66, 97, 29, 47, 3, 36, 34, 43, 25}))

	fmt.Println(numberOfPairs([]int{0, 4, 5, 0, 2, 1, 6, 4, 3, 2, 2, 4, 3, 7, 6, 1, 4, 3, 4, 2}))
	fmt.Println(numberOfPairs([]int{71, 60}))
	fmt.Println(numberOfPairs([]int{17, 28, 39, 71}))

	// Output:
	// 39
	// 190
	// 1
	// 3
}