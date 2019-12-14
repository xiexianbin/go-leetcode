package t38

import (
	"fmt"
	"testing"
)

func countAndSay(n int) string {
	r := ""
	//b := fmt.Sprintf("%b", n)
	//array := strings.Split(b, "")
	return r
}

func TestCountAndSay(t *testing.T) {
	for _, i := range []int{1, 2, 3, 4, 5} {
		fmt.Println(countAndSay(i))
	}
}
