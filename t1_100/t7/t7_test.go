package t7

import (
	"fmt"
	"testing"
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

func TestReverse(t *testing.T) {
	x := 123

	fmt.Println(reverse(x))
}
