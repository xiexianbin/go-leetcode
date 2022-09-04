package t49

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	cache := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		t := strings.Split(strs[i], "")
		sort.Strings(t)
		key := strings.Join(t, "")
		if v, ok := cache[key]; ok {
			cache[key] = append(v, strs[i])
		} else {
			cache[key] = []string{strs[i]}
		}
	}

	result := [][]string{}
	for _, v := range cache {
		result = append(result, v)
	}
	return result
}

func ExampleGroupAnagrams() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	for i := 0; i < len(result); i++ {
		t := result[i]
		sort.Strings(t)
		if len(t) == 3 {
			fmt.Println(t)
		}
	}

	// Output:
	// [ate eat tea]
}
