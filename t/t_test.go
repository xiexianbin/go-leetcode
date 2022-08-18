package t

import (
	"fmt"
	"testing"
)

func T() {
	fmt.Println("Hello World!")
}

func ExampleT() {
	T()

	// Output:
	// Hello World!
}

func TestT(t *testing.T) {
	t.Log("Hello World!")
}
