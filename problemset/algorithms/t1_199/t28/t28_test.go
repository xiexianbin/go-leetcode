package t28

import (
	"fmt"
	"testing"
)

func strStr(haystack string, needle string) int {
	subLen := len(needle)
	for i := 0; i <= len(haystack)-subLen; i++ {
		subStr := haystack[i : i+subLen]
		if subStr == needle {
			return i
		}
	}
	return -1
}

func TestT28(t *testing.T) {
	//haystack := "hello"
	//needle := "ll"
	//haystack := "aaaaa"
	//needle := "bba"
	haystack := "hello"
	needle := "lo"
	fmt.Println(strStr(haystack, needle))

	// Output:
}
