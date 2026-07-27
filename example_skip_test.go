package opiter_test

import (
	"fmt"
	"slices"

	"github.com/ngicks/opiter"
)

func ExampleSkipLast() {
	fmt.Println("all: ", slices.Collect(opiter.Range(0, 10)))
	fmt.Println("kept:", slices.Collect(opiter.SkipLast(5, opiter.Range(0, 10))))
	// Output:
	// all:  [0 1 2 3 4 5 6 7 8 9]
	// kept: [0 1 2 3 4]
}
