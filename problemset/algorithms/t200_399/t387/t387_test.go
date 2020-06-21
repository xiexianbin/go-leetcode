package t387

import (
	"fmt"
)

func firstUniqChar(s string) int {
	findMap := make(map[string]int)

	for _, v := range []rune(s) {
		if _, ok := findMap[string(v)]; ok {
			findMap[string(v)]++
		} else {
			findMap[string(v)] = 1
		}
	}

	for i, v := range []rune(s) {
		c, _ := findMap[string(v)]
		if c == 1{
			return i
		}
	}

	return -1
}

func ExampleFirstUniqChar() {
	ss := []string{"leetcode", "loveleetcode"}

	for i := 0; i < len(ss); i++ {
		fmt.Println(firstUniqChar(ss[i]))
	}

	// Output:
	// 0
	// 2
}
