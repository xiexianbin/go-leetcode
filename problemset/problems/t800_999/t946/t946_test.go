package t946

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		j := 0
		for j < len(popped) {
			if len(stack) > 0 && stack[len(stack)-1] == popped[j] {
				stack = stack[:len(stack)-1]
				j++
				continue
			}
			break
		}
		if j > 0 {
			popped = popped[j:]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func ExampleValidateStackSequences() {
	fmt.Println(validateStackSequences([]int{1, 0}, []int{1, 0}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))

	// Output:
	// true
	// ture
}
