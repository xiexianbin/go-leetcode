package t151

import (
	"fmt"
	"strings"
	"testing"
)

func reverseWords(s string) string {
	sL := strings.Split(s, " ")
	c := len(sL)
	n := 0
	i := 0
	for {
		if i >= c {
			break
		}
		if sL[i] == "" {
			i++
			continue
		} else {
			sL[n], sL[i] = sL[i], sL[n]
			n++
			i++
		}
	}
	for i := 0; i < n/2; i++ {
		sL[i], sL[n-i-1] = sL[n-i-1], sL[i]
	}
	return strings.Join(sL[:n], " ")
}

func TestTest(t *testing.T) {
	for _, s := range []string{"a good   example", "  hello world!  ", "the sky is blue"} {
		fmt.Println(reverseWords(s))
	}
}
