package t720

import (
	"fmt"
	"sort"
	"strings"
)

func longestWord(words []string) string {
	m := make(map[int][]string)
	maxLength := 0
	for i := 0; i < len(words); i++ {
		k := len(words[i])
		m[k] = append(m[k], words[i])
		if k > maxLength {
			maxLength = k
		}
	}
	if len(m) < 0 {
		return ""
	}

	var matchWords []string
	var ok bool
	if matchWords, ok = m[1]; ok == false {
		return ""
	}
	for i := 2; i <= maxLength; i++ {
		var matchWordsTmp []string
		if value, ok := m[i]; ok {
			for _, s := range value {
				for _, mw := range matchWords {
					if strings.HasPrefix(s, mw) {
						matchWordsTmp = append(matchWordsTmp, s)
					}
				}
			}
			if matchWordsTmp == nil {
				break
			} else {
				matchWords = matchWordsTmp
			}
		} else {
			break
		}
	}

	if len(matchWords) > 0 {
		sort.Strings(matchWords)
		return matchWords[0]
	}

	return ""
}

func ExampleLongestWord() {
	wordss := [][]string{
		{"a", "banana", "app", "appl", "ap", "apply", "apple"},
		{"w", "wo", "wor", "worl", "world"},
	}

	for _, words := range wordss {
		fmt.Println(longestWord(words))
	}
	// Output:
	// apple
	// world
}
