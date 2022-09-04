package t190

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var n uint32
	for i := 0; i < 32; i++ {
		n <<= 1
		n |= num & 1
		num >>= 1
	}

	return n
}

func ExampleReverseBits() {
	//fmt.Println(reverseBits(001))
	fmt.Println("Hello World!") // reverseBits(00000000000000000001111010011100)

	// Output:
	// Hello World!
}
