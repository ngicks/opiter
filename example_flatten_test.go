package opiter_test

import (
	"fmt"
	"slices"

	"github.com/ngicks/opiter"
)

func ExampleFlatten() {
	grouped := [][]string{
		{"apple", "banana"},
		{"carrot"},
		{"dog", "elephant", "fox"},
	}

	fmt.Println(slices.Collect(opiter.Flatten(slices.Values(grouped))))
	// Output:
	// [apple banana carrot dog elephant fox]
}

func ExampleFlattenF() {
	categories := [][]string{{"fruits", "vegetables"}, {"proteins"}}
	counts := []int{10, 5}

	for category, count := range opiter.FlattenF(
		opiter.Pairs(slices.Values(categories), slices.Values(counts)),
	) {
		fmt.Printf("%s: %d\n", category, count)
	}
	// Output:
	// fruits: 10
	// vegetables: 10
	// proteins: 5
}

func ExampleFlattenL() {
	products := []string{"laptop", "phone"}
	prices := [][]int{{999, 1099, 899}, {599, 649}}

	for product, price := range opiter.FlattenL(
		opiter.Pairs(slices.Values(products), slices.Values(prices)),
	) {
		fmt.Printf("%s: %d\n", product, price)
	}
	// Output:
	// laptop: 999
	// laptop: 1099
	// laptop: 899
	// phone: 599
	// phone: 649
}
