package t17

import (
	"fmt"
	"strconv"
	"strings"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	keys := map[int][]string{
		0: {},
		1: {},
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	var ans []string
	var traceback func(sub string, subdigits []string)
	traceback = func(sub string, subdigits []string) {
		if len(subdigits) == 0 {
			ans = append(ans, sub)
			return
		}
		num, _ := strconv.Atoi(subdigits[0])
		words := keys[num]
		for _, w := range words {
			traceback(fmt.Sprintf("%s%s", sub, w), subdigits[1:])
		}
	}

	subdigits := strings.Split(digits, "")
	traceback("", subdigits)
	return ans
}

func ExampleLetterCombinations() {
	fmt.Println(letterCombinations("23"))

	// Output:
	// [ad ae af bd be bf cd ce cf]
}
