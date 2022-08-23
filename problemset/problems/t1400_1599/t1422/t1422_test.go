package t1422

import "fmt"

func maxScore(s string) int {
	l := len(s)
	max := 0
	for i := 1; i <= l-1; i++ {
		l, r, ls, rs := s[:i], s[i:], 0, 0
		for j := 0; j < len(l); j++ {
			if l[j] == '0' {
				ls++
			}
		}
		for j := 0; j < len(r); j++ {
			if r[j] == '1' {
				rs++
			}
		}

		if ls+rs > max {
			max = ls + rs
		}
	}
	return max
}

func ExampleMaxScore() {
	ss := []string{"00", "011101", "00111", "1111"}
	for _, s := range ss {
		fmt.Println(maxScore(s))
	}

	// Output:
	// 1
	// 5
	// 5
	// 3
}
