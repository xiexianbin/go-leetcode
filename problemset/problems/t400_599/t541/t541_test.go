package t541

import (
	"fmt"
	"strings"
)

func reverseString(s []string) string {
	if len(s) < 2 {
		return strings.Join(s, "")
	}

	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return strings.Join(s, "")
}

func reverseStr(s string, k int) string {
	result := []string{}
	strList := strings.Split(s, "")
	i := 0
	for i+2*k <= len(s) {
		s1 := reverseString(strList[i : i+k])
		result = append(result, s1)
		for j := i + k; j < i+2*k; j++ {
			result = append(result, strList[j])
		}
		i += 2 * k
	}

	if len(s)-i <= k {
		result = append(result, reverseString(strList[i:]))
	} else {
		s1 := reverseString(strList[i : i+k])
		result = append(result, s1)
		for j := i + k; j < len(s); j++ {
			result = append(result, strList[j])
		}
	}

	return strings.Join(result, "")
}

func ExampleReverseStr() {
	strs := []string{"abcdefg", "abcdefg", "abcdef"}
	ks := []int{4, 2, 2}
	for k, v := range strs {
		fmt.Println(reverseStr(v, ks[k]))
	}

	// Output:
	//dcbaefg
	//bacdfeg
	//bacdfe
}
