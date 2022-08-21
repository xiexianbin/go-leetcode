package t1455

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	strList := strings.Split(sentence, " ")
	for i, str := range strList {
		if strings.HasPrefix(str, searchWord) {
			return i + 1
		}
	}

	return -1
}

func ExampleBusyStudent() {
	ss := []string{"i love eating burger", "this problem is an easy problem", "i am tired"}
	sw := []string{"burg", "pro", "you"}
	for i := 0; i < len(ss); i++ {
		fmt.Println(isPrefixOfWord(ss[i], sw[i]))
	}

	// Output:
	// 4
	// 2
	// -1
}
