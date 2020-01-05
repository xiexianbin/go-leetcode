package t67

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func addBinary(a string, b string) string {
	a1 := strings.Split(a, "")
	b1 := strings.Split(b, "")
	s := []int{}
	t := 0
	for i, j := len(a1)-1, len(b1)-1; i >= 0 || j >= 0; {
		var ai, bj int
		if i >= 0 {
			if a1[i] == "1" {
				ai = 1
			}
		}

		if j >= 0 {
			if b1[j] == "1" {
				bj = 1
			}
		}

		s = append(s, (ai+bj+t)%2)
		t = (ai + bj + t)/2
		i--
		j--
	}

	if t > 0 {
		s = append(s, t)
	}

	sum := ""
	for i :=len(s) - 1; i >= 0; i-- {
		sum += strconv.Itoa(s[i])
	}
	return sum
}

func TestAddBinary(t *testing.T) {
	//a := "11"
	//b := "1"
	a := "1010"
	b := "1011"

	fmt.Println(addBinary(a, b))
}
