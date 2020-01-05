package t38

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func countAndSay(n int) string {
	cache := make([]string, n)
	cache[0] = "1"
	for i := 1; i < n; i++ {
		s := cache[i-1]
		l := len(s)

		r := []string{}
		for j := 0; j < l; j++ {
			if j+1 < l && s[j] == s[j+1] {
				c := s[j]
				t := 2
				for {
					if j+t >= l || c != s[j+t] {
						break
					}
					if c == s[j+t] {
						t++
					}
				}
				r = append(r, strconv.Itoa(t))
				r = append(r, string(c))
				j += t - 1
				continue
			}
			r = append(r, "1")
			r = append(r, string(s[j]))
			if i == 9 {
				continue
			}
		}
		cache[i] = strings.Join(r, "")
	}

	return cache[n-1]
}

func TestCountAndSay(t *testing.T) {
	for _, i := range []int{10, 6, 1, 2, 3, 4, 5} {
		fmt.Println(countAndSay(i))
	}
}
