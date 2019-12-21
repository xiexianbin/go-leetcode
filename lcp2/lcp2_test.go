package lcp2

import (
	"fmt"
	"testing"
)

func gcd(m, n int) []int {
	if n > 1 {
		skip := false
		for {
			for i := 2; i <= n; i++ {
				if m%i == 0 && n%i == 0 {
					m, n = m/i, n/i
				}
				if i == n {
					skip = true
					break
				}
				if i >= (n+1)/2 {
					i = n - 1
				}
			}
			if skip {
				break
			}
		}
	}
	return []int{m, n}
}

func fraction(cont []int) []int {
	m := cont[len(cont)-1]
	n := 1
	for i := len(cont) - 2; i >= 0; i-- {
		n = cont[i]*m + n
		m, n = n, m
	}
	m, n = n, m
	return gcd(n, m)
}

func TestTest(t *testing.T) {
	conts := [][]int{{0, 2147483647}, {3}, {12, 11}, {3, 2, 0, 2}, {0, 0, 3}}

	for _, v := range conts {
		fmt.Println(fraction(v))
	}
}
