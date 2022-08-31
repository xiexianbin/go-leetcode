package t6162

import (
	"fmt"
	"strings"
)

func gotoback(garbage []string, c string) bool {
	for i := 0; i < len(garbage); i++ {
		if strings.Contains(garbage[i], c) {
			return true
		}
	}
	return false
}

func garbageCollectionOne(garbage []string, c string, travel []int) int {
	t := 0
	clean := true
	for i := 0; i < len(garbage); i++ {
		if strings.Contains(garbage[i], c) {
			clean = false
			if i > 0 {
				t += travel[i-1]
			}
			t += strings.Count(garbage[i], c)
			garbage[i] = strings.ReplaceAll(garbage[i], c, "")
		} else if gotoback(garbage[i:], c) {
			if i > 0 {
				t += travel[i-1]
			}
		}
	}

	if clean {
		return 0
	}

	return t
}

func garbageCollection(garbage []string, travel []int) int {
	answer := 0
	answer += garbageCollectionOne(garbage, "P", travel)
	answer += garbageCollectionOne(garbage, "M", travel)
	answer += garbageCollectionOne(garbage, "G", travel)

	return answer
}

func ExampleT() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))

	// Output:
	// Hello World!
}
