package t171

import (
	"fmt"
	"strings"
)

func titleToNumber(s string) int {
	num := 0
	nums := strings.Split(s, "")
	for i := 0; i < len(nums); i++ {
		r := []rune(nums[i])
		c := r[0] - 65 + 1
		num = num*26 + int(c)
	}
	return num
}

func ExampleTitleToNumber() {
	lines := []string{"A", "AB", "ZY"}
	for _, line := range lines {
		fmt.Println(titleToNumber(line))
	}

	// Output:
	// 1
	// 28
	// 701
}
