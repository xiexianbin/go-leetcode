package t7

import (
	"fmt"
)

func reverse(x int) int {
	y := int32(x)
	var z int32
	for y != 0 {
		t := z*10 + y%10
		if t/10 != z {
			return 0
		}
		z = t
		y = y / 10
	}
	return int(z)
}

func ExampleReverse() {
	x := 123

	fmt.Println(reverse(x))

	// Output:
	// 321
}
