package t20

import (
	"fmt"
	"strings"
	"testing"
)

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := []string{}
	for _, b := range strings.Split(s, "") {
		lb := ""
		if len(stack) > 0 {
			lb = stack[len(stack)-1]
		}

		if (b == ")" && lb == "(") || b == "]" && lb == "[" || b == "}" && lb == "{" {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, b)
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

func TestValidParentheses(t *testing.T) {
	str := [5]string{"()", "()[]{}", "(]", "([)]", "{[]}"}

	for _, s := range str {
		fmt.Println(isValid(s))
	}
}
