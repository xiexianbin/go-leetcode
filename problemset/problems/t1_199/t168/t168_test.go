package t168

import (
	"fmt"
	"strings"
)

func convertToTitle(n int) string {
	ans := []string{}
	for {
		if n > 26 {
			if n%26 == 0 {
				ans = append(ans, "Z")
				n = n/26 - 1
			} else {
				ans = append(ans, string(rune(n%26+64)))
				n = n / 26
			}
		} else {
			ans = append(ans, string(rune(n+64)))
			break
		}
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}

	return strings.Join(ans, "")
}

func ExampleConvertToTitle() {
	arr := []int{26, 52, 1, 28, 701}
	for _, n := range arr {
		fmt.Println(convertToTitle(n))
	}

	// Output:
	// Z
	// AZ
	// A
	// AB
	// ZY
}
