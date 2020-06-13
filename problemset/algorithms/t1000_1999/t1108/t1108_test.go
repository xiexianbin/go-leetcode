package t1108

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", 4)
}

func ExampleDefangIPaddr() {
	addresses := []string{"1.1.1.1", "255.100.50.0"}

	for _, address := range addresses {
		fmt.Println(defangIPaddr(address))
	}

	// Output:
	// 1[.]1[.]1[.]1
	// 255[.]100[.]50[.]0
}
