package t68

import (
	"fmt"
	"strings"
	"testing"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	s = strings.TrimRight(s, " ")

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

func TestMerge(t *testing.T) {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("Hello Wor2ld"))
	fmt.Println(lengthOfLastWord("Hello World "))
	fmt.Println(lengthOfLastWord("a "))

	// Output:
}
