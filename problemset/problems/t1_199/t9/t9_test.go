package t9

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	sL := strings.Split(s, "")
	for i := 0; i < len(sL)/2; i++ {
		if sL[i] != sL[len(sL)-1-i] {
			return false
		}
	}
	return true
}

func ExampleIsPalindrome() {
	for _, x := range []int{121, -121, 10} {
		fmt.Println(isPalindrome(x))
	}

	// Output:
	// true
	// false
	// false
}
