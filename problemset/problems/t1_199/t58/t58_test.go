package t68

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")

	if len(s) == 0 {
		return 0
	}

	sList := strings.Split(s, " ")
	ls := sList[len(sList)-1]
	for _, v := range ls {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			continue
		} else {
			return 0
		}
	}

	return len(ls)
}

func ExampleLengthOfLastWord() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("Hello World "))
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord(""))

	// Output:
	// 5
	// 5
	// 1
	// 0
}
