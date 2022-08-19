package t5

import "fmt"

func longestPalindrome(s string) string {
	size := len(s)
	if size < 2 {
		return s
	}

	maxLen := 1
	begin := 0
	var dp = make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
		dp[i][i] = true
	}

	for l := 2; l <= size; l++ {
		for i := 0; i < size; i++ {
			j := l + i - 1
			// right index j out of index
			if j >= size {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] == true && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}

func ExampleLongestPalindrome() {
	ss := []string{"xaabacxcabaaxcabaax", "babad", "cbbd"}
	for _, s := range ss {
		fmt.Println(longestPalindrome(s))
	}

	// Output:
	// xaabacxcabaax
	// bab
	// bb
}
