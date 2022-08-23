package t1417

import (
	"fmt"
)

func reformat(s string) string {
	nc, wc, ns, ws := 0, 0, "", ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			nc++
			ns += string(s[i])
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			wc++
			ws += string(s[i])
		}
	}
	ans := ""
	if nc == wc {
		for i := 0; i < nc; i++ {
			ans += string(ns[i])
			ans += string(ws[i])
		}
	} else if nc == wc+1 {
		for i := 0; i < wc; i++ {
			ans += string(ns[i])
			ans += string(ws[i])
		}
		ans += string(ns[nc-1])
	} else if wc == nc+1 {
		for i := 0; i < nc; i++ {
			ans += string(ws[i])
			ans += string(ns[i])
		}
		ans += string(ws[wc-1])
	}

	return ans
}

func ExampleReformat() {
	ss := []string{"s", "a0b1c2", "leetcode", "123456789"}
	for _, s := range ss {
		fmt.Println(reformat(s))
	}

	// Output:
	// s
	// 0a1b2c
}
