package t191

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	count := 0
	fmt.Println(num)
	num = num&(-num)
	fmt.Println(num)

	return count
}

func ExampleHammingWeight(){
	for _, num := range []uint32{00000000000000000000000000001011, 00000000000000000000000010000000} {
		fmt.Println(hammingWeight(num))
	}

	// Output:
	//521
	//1
	//0
	//2097152
	//2097152
	//0
}
