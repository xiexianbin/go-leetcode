package t6

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	source := strings.Split(s, "")
	result := []string{}
	resultMap := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		resultMap[i] = make([]string, len(s))
	}

	iAddFlag := true
	i := 0
	j := 0
	for n := 0; n < len(s); n++ {
		resultMap[i][j] = source[n]

		if iAddFlag == true && i%numRows == numRows-1 {
			iAddFlag = false
			i--
			j++
		} else if iAddFlag {
			i++
		} else {
			i--
			j++
		}
		if i == 0 {
			iAddFlag = true
		}
	}

	for i := 0; i < numRows; i++ {
		newRow := resultMap[i]
		result = append(result, strings.Join(newRow, ""))
	}
	return strings.Join(result, "")
}

func ExampleConvert() {
	ss := []string{"AB", "LEETCODEISHIRING", "LEETCODEISHIRING"}
	numRowss := []int{1, 3, 4}
	for i, s := range ss {
		fmt.Println(convert(s, numRowss[i]))
	}

	// Output:
	// AB
	// LCIRETOESIIGEDHN
	// LDREOEIIECIHNTSG
}
