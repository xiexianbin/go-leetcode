package t344

import (
	"fmt"
)

//func reverseString(s []byte) {
//	if len(s) < 2 {
//		return
//	}
//
//	i := 0
//	j := len(s) - 1
//	for i < j {
//		s[i], s[j] = s[j], s[i]
//		i++
//		j--
//	}
//}

func switchByte(s []byte, start, end int) {
	if start >= end {
		return
	}

	s[start], s[end] = s[end], s[start]
	start++
	end--
	switchByte(s, start, end)
}

func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}

	switchByte(s, 0, len(s)-1)
}

func ExampleReverseString() {
	strs := [][]byte{[]byte("hello"), []byte("Hannah")}
	for _, s := range strs {
		reverseString(s)
		fmt.Println(string(s))
	}

	// Output:
	//olleh
	//hannaH
}
