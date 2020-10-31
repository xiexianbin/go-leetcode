package t

import (
	"fmt"
	"sort"
	"strings"
)

func CheckPermutation(s1 string, s2 string) bool {
	s1Arry := strings.Split(s1, "")
	s2Arry := strings.Split(s2, "")
	if len(s1Arry) != len(s2Arry) {
		return false
	}
	sort.Strings(s1Arry)
	sort.Strings(s2Arry)
	for i := 0; i < len(s1Arry); i++ {
		if s1Arry[i] != s2Arry[i] {
			return false
		}
	}
	return true
}

func ExampleCheckPermutation() {
	s1 := []string{"abc", "bca"}
	s2 := []string{"bca", "bad"}

	for i := 0; i < len(s1); i++ {
		fmt.Println(CheckPermutation(s1[i], s2[i]))
	}

	// Output:
	// true
	// false
}
