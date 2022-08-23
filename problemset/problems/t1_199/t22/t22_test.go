package t22

import "fmt"

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(left, right int, path string)
	dfs = func(left, right int, path string) {
		if left == 0 && right == 0 {
			ans = append(ans, path)
		}
		if left > 0 {
			dfs(left-1, right, fmt.Sprintf("%s(", path))
		}
		if left < right {
			dfs(left, right-1, fmt.Sprintf("%s)", path))
		}
	}
	dfs(n, n, "")
	return ans
}

func ExampleGenerateParenthesis() {
	for i := 1; i <= 3; i++ {
		fmt.Println(len(generateParenthesis(i)))
	}

	// Output:
	// 1
	// 2
	// 5
}
