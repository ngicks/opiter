package opiter_test

import (
	"errors"
	"fmt"
	"slices"

	"github.com/ngicks/opiter"
)

func ExampleTryFind() {
	input := opiter.Pairs(
		slices.Values([]int{1, 2, 3, 4, 5}),
		opiter.Concat(
			opiter.Repeat(error(nil), 2),
			opiter.Once(errors.New("processing error")),
			opiter.Repeat(error(nil), 2),
		),
	)

	found, err := opiter.TryFind(func(n int) bool { return n > 2 }, input)
	fmt.Println("found:", found.IsSome())
	fmt.Println("error:", err)
	// Output:
	// found: false
	// error: processing error
}

func ExampleTryForEach() {
	input := opiter.Pairs(
		slices.Values([]string{"apple", "banana", "cherry", "date"}),
		opiter.Concat(
			opiter.Repeat(error(nil), 2),
			opiter.Once(errors.New("bad fruit")),
			opiter.Once(error(nil)),
		),
	)

	count := 0
	err := opiter.TryForEach(func(fruit string) {
		count++
		fmt.Println("processing:", fruit)
	}, input)
	fmt.Printf("stopped after %d: %v\n", count, err)
	// Output:
	// processing: apple
	// processing: banana
	// stopped after 2: bad fruit
}
