package t8

import (
	"fmt"
)

func myAtoi(s string) int {
	min32Int, max32Int := -1<<31, 1<<31-1
	i, flag, sign, result := 0, false, 1, 0
	l := len(s)
	for ; i < l && s[i] == ' '; i++ {
	}
	if l == i {
		return 0
	}

	for ; i < l; i++ {
		k := string(rune(s[i]))
		if k == "-" && flag == false {
			sign = -1
			flag = true
		} else if k == "+" && flag == false {
			sign = 1
			flag = true
		} else if k == "." {
			return result
		} else if k >= "0" && k <= "9" {
			v := int(s[i] - '0')
			if flag == false {
				if v < 0 {
					sign = -1
				}
				flag = true
			}
			result = sign * (sign*result*10 + v)
		} else {
			break
		}

		if sign == -1 && result < min32Int {
			return min32Int
		}
		if sign == 1 && result > max32Int {
			return max32Int
		}
	}

	return result
}

func ExampleMyAtoi() {
	ss := []string{"00000-42a1234", " ", " b11228552307", "-5-", "  0000000000012345678", "   -42", "20000000000000000000", "+-12", ".1", "3.14159", "words and 987", "42", "4193 with words"}
	for _, s := range ss {
		fmt.Println(myAtoi(s))
	}

	// Output:
	// 0
	// 0
	// 0
	// -5
	// 12345678
	// -42
	// 2147483647
	// 0
	// 0
	// 3
	// 0
	// 42
	// 4193
}
