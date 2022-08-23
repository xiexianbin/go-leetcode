package t125

import (
	"fmt"
)

func isPalindrome(s string) bool {
	l := len(s)
	if l < 2 {
		return true
	}

	t := ""
	for i := 0; i < l; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			t += string(s[i] + 32)
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			t += string(s[i])
		}
		if s[i] >= '0' && s[i] <= '9' {
			t += string(s[i])
		}
	}

	nl := len(t)
	for i := 0; i < nl/2; i++ {
		if t[i] != t[nl-i-1] {
			return false
		}
	}
	return true
}

func ExampleIsPalindrome() {
	ss := []string{"0P", "A man, a plan, a canal: Panama", "race a car"}
	for _, s := range ss {
		fmt.Println(isPalindrome(s))
	}

	// Output:
	// false
	// true
	//false
}
