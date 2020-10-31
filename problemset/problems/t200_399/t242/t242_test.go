package t242

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	targeMap := make(map[string]int)
	for _, v := range []rune(s) {
		if _, ok := targeMap[string(v)]; ok {
			targeMap[string(v)]++
		} else {
			targeMap[string(v)] = 1
		}
	}

	for _, v := range []rune(t) {
		if _, ok := targeMap[string(v)]; ok {
			targeMap[string(v)]--
			if targeMap[string(v)] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func ExampleT() {
	ss := []string{"anagram", "rat"}
	ts := []string{"nagaram", "cat"}

	for i := 0; i < len(ss); i++ {
		fmt.Println(isAnagram(ss[i], ts[i]))
	}

	// Output:
	// true
	// false
}
