package t14

import (
	"fmt"
	"strings"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	isCommon := true
	commonCount := 0
	for _, s := range strings.Split(strs[0], "") {
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= commonCount {
				isCommon = false
				break
			}
			strList := strings.Split(strs[i], "")
			if s != strList[commonCount] {
				isCommon = false
			}
		}
		if isCommon == false {
			break
		}
		commonCount++
	}

	return strings.Join(strings.Split(strs[0], "")[:commonCount], "")
}

func commonPrefix(str1, str2 string) string {
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}

	i := 0
	for i < len(str1) {
		if strings.Split(str1, "")[i] == strings.Split(str2, "")[i]{
			i++
		} else {
			break
		}
	}

	return strings.Join(strings.Split(str1, "")[:i], "")
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	str1 := strs[0]
	for i := 1; i < len(strs); i++ {
		str1 = commonPrefix(str1, strs[i])
	}

	return str1
}

func TestTest(t *testing.T) {
	strss := [][]string{{"aa", "a"}, {"flower", "flow", "flight"}, {"dog", "racecar", "car"}}
	for _, strs := range strss {
		fmt.Println(longestCommonPrefix(strs))
		fmt.Println(longestCommonPrefix2(strs))
	}

	// Output:
}
