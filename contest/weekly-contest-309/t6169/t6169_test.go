package t6169

import (
	"fmt"
	"testing"
)

func longestNiceSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	max := 1
	for i := 0; i < len(nums); i++ {
		cache := []int{nums[i]}
		for j := i + 1; j < len(nums) && i < len(nums); j++ {
			flag := false
			for x := 0; x < len(cache); x++ {
				if cache[x]&nums[j] != 0 {
					flag = false
					break
				}
				flag = true
			}
			if flag {
				cache = append(cache, nums[j])
			} else {
				break
			}
		}
		if len(cache) > max {
			max = len(cache)
		}
	}

	return max
}

func ExampleT() {
	fmt.Println(longestNiceSubarray([]int{45106826, 547958667, 823366125, 332020148, 611677524, 510346561, 555831456, 436600904, 12594192, 127206768, 540754485, 201997978, 473116514, 233000361, 538246458, 729745279, 343417143, 892046691, 376031730}))
	fmt.Println(longestNiceSubarray([]int{904163577, 321202512, 470948612, 490925389, 550193477, 87742556, 151890632, 655280661, 4, 263168, 32, 573703555, 886743681, 937599702, 120293650, 725712231, 257119393}))
	fmt.Println(longestNiceSubarray([]int{744437702, 379056602, 145555074, 392756761, 560864007, 934981918, 113312475, 1090, 16384, 33, 217313281, 117883195, 978927664}))
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10}))

	// Output:
	// 3
	// 3
	// 3
	// 3
}

func TestT(t *testing.T) {
	t.Log("Hello World!")
}
