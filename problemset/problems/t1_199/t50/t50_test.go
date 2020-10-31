package t50

import (
	"fmt"
	"testing"
)

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.00000
	}
	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x, n)
}

func TestTest(t *testing.T) {
	d := map[float64]int{}
	d[2.00000] = 10
	d[2.10000] = 3
	d[2.00000] = -2

	for k, v := range d {
		fmt.Println(myPow(k, v))
	}

	// Output:
}
