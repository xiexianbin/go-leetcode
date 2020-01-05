package t779

import (
	"fmt"
	"testing"
)

func kthGrammar(N int, K int) int {
	
}

func TestTest(t *testing.T) {
	d := map[int]int{}
	d[1] = 1
	d[2] = 1
	d[2] = 2
	d[4] = 5
	for k, v := range d {
		fmt.Println(kthGrammar(k, v))
	}
}
