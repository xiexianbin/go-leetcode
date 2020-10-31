package t13

import (
	"fmt"
	"strings"
	"testing"
)

func romanToInt(s string) int {
	d := map[string]int{}
	d["I"] = 1
	d["V"] = 5
	d["X"] = 10
	d["L"] = 50
	d["C"] = 100
	d["D"] = 500
	d["M"] = 1000
	d["IV"] = 4
	d["IX"] = 9
	d["XL"] = 40
	d["XC"] = 90
	d["CD"] = 400
	d["CM"] = 900

	result := 0
	sL := strings.Split(s, "")
	for i := 0; i < len(sL); i++ {
		if i+1 < len(sL) {
			k := fmt.Sprint(sL[i], sL[i+1])
			if v, ok := d[k]; ok {
				result += v
				i++
				continue
			}
		}
		if v, ok := d[sL[i]]; ok {
			result += v
		}
	}

	return result
}

func TestTest(t *testing.T) {
	for _, s := range []string{"III", "IV"} {
		fmt.Println(romanToInt(s))
	}

	// Output:
}
