package t150

import (
	"fmt"
	"strconv"
	"testing"
)

type Stack struct {
	stack []int
}

func (this *Stack) Push(v int) {
	this.stack = append(this.stack, v)
}

func (this *Stack) Pop() int {
	r := this.stack[len(this.stack)-1]
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
	return r
}

func evalRPN(tokens []string) int {
	stack := Stack{
		stack: make([]int, len(tokens)),
	}
	for i := 0; i < len(tokens); i++ {
		v := tokens[i]
		switch v {
		case "+":
			k1 := stack.Pop()
			k2 := stack.Pop()
			stack.Push(k1 + k2)
		case "-":
			k1 := stack.Pop()
			k2 := stack.Pop()
			stack.Push(k2 - k1)
		case "*":
			k1 := stack.Pop()
			k2 := stack.Pop()
			stack.Push(k1 * k2)
		case "/":
			k1 := stack.Pop()
			k2 := stack.Pop()
			stack.Push(k2 / k1)
		default:
			i, err := strconv.Atoi(v)
			if err != nil {
				return -1
			}
			stack.Push(i)
		}
	}

	return stack.Pop()
}

/* https://leetcode-cn.com/explore/learn/card/queue-stack/218/stack-last-in-first-out-data-structure/880/ */
func TestT880(t *testing.T) {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
