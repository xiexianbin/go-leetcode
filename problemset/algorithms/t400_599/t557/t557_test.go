package t557

import (
	"fmt"
	"strings"
)

func reverseOneWord(s string) string {
	sL := strings.Split(s, "")
	l := len(sL)
	for i := 0; i < l/2; i++ {
		sL[i], sL[l-i-1] = sL[l-i-1], sL[i]
	}
	return strings.Join(sL, "")
}

func reverseWords(s string) string {
	if s == "" {
		return s
	}
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	sL := strings.Split(s, " ")
	for i := 0; i < len(sL); i++ {
		sL[i] = reverseOneWord(sL[i])
	}

	return strings.Join(sL, " ")
}

func ExampleReverseWords() {
	strs := []string{"Let's take LeetCode contest"}
	for _, s := range strs {
		fmt.Println(reverseWords(s))
	}

	// Output:
	// s'teL ekat edoCteeL tsetnoc
}
