package t3

import (
	"fmt"
	"strings"
	"testing"
)

func allDiff(str string, s int, e int) bool {
	str2 := []byte(str)
	var set []byte
	for i := s; i < e; i++{
		b := byte(str2[i])
		if strings.Contains(string(set), string(b)) {
			return false
		}
		set = append(set, b)
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	var max int
	l := len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if max > j - i {
				continue
			}
			if allDiff(s, i, j) {
				t := j - i
				if t > max {
					max = t
				}
			} else {
				break
			}
		}
	}
	return max
}

func TestTest(t *testing.T) {
	for _, s := range []string{"abcabcbb", "bbbbb", "pwwkew"} {
		fmt.Println(lengthOfLongestSubstring(s))
	}
}
