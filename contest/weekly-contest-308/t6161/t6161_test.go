package t6161

import (
	"fmt"
)

func removeStars(s string) string {
	result := ""
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == '*' {
			c := 1
			for {
				i++
				if i >= l {
					break
				}
				if s[i] != '*' {
					break
				}
				c++
			}
			i--
			result = result[:len(result)-c]
		} else {
			c := 1
			for {
				i++
				if i >= l {
					break
				}
				if s[i] == '*' {
					break
				}
				c++
			}
			i--
			result += s[i+1-c : i+1]
		}
	}
	return result
}

func ExampleT() {
	fmt.Println(removeStars("erase*****"))
	fmt.Println(removeStars("leet**cod*e"))

	// Output:
	//
	// lecoe
}
