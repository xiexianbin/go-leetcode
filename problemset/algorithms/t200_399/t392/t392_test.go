package t392

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sl, tl := 0, 0
	for sl < len(s) && tl < len(t) {
		if s[sl] == t[tl] {
			sl++
			tl++
		} else {
			tl++
		}
		if sl == len(s) {
			return true
		}
	}

	return false
}

func ExampleT2() {
	ques := [][]string{{"abc", "ahbgdc"}, {"axc", "ahbgdc"}, {"", "abc"}}
	for _, que := range ques {
		s, t := que[0], que[1]
		fmt.Println(isSubsequence(s, t))
	}

	// Output:
	// true
	// false
	// true
}
