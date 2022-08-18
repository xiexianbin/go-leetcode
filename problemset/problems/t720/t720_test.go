package t720

import (
	"fmt"
	"sort"
)

func longestWord(words []string) string {
	ans := []int{0}
	longestCount := len(words[0])
	for i := 1; i < len(words); i++ {
		if len(words[i]) == longestCount {
			ans = append(ans, i)
		} else if len(words[i]) > longestCount {
			longestCount = len(words[i])
			ans = []int{i}
		}
	}

	newsWords := []string{}
	for i := 0; i < len(ans); i++ {
		newsWords = append(newsWords, words[ans[i]])
	}

	sort.Strings(newsWords)

	return newsWords[0]
}

func ExampleLongestWord() {
	wordss := [][]string{{"w", "wo", "wor", "worl", "world"}, {"a", "banana", "app", "appl", "ap", "apply", "apple"}}

	for _, words := range wordss {
		fmt.Println(longestWord(words))
	}
	// Output:
	// world
	// banana
}
