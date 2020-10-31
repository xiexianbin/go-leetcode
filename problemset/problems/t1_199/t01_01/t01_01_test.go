package t01_01

import (
	"fmt"
	"sort"
)

func isUnique(astr string) bool {
	s := []rune(astr)
	if len(s) == 1 {
		return true
	}

	b := make([]int, len(astr))
	for i := 0; i < len(s); i++ {
		b[i] = int(s[i])
	}

	sort.Ints(b)
	for i := 1; i < len(b); i++ {
		if b[i-1] == b[i] {
			return false
		}
	}
	return true
}

func ExampleIsUnique() {
	ss := []string{"kzwunahkiz", "leetcode", "abc"}

	for _, s := range ss {
		fmt.Println(isUnique(s))
	}
	// Output:
	// false
	// false
	// true
}
