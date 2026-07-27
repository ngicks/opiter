package opiter_test

import (
	"fmt"
	"slices"

	"github.com/ngicks/opiter"
)

func ExampleCompact() {
	merged := opiter.Merge(
		opiter.Map(func(i int) int { return 2 * i }, opiter.Range(1, 11)),
		opiter.Map(func(i int) int { return 1 << i }, opiter.Range(1, 11)),
	)

	fmt.Println(slices.Collect(opiter.Compact(merged)))
	// Output:
	// [2 4 6 8 10 12 14 16 18 20 32 64 128 256 512 1024]
}

func ExampleCompactFunc2() {
	type record struct {
		Key  string
		Data string
	}
	records := []record{
		{"foo", "yay"},
		{"foo", "nay"},
		{"foo", "mah"},
		{"bar", "yay"},
		{"baz", "yay"},
		{"baz", "nay"},
	}

	for i, v := range opiter.CompactFunc2(
		func(_ int, a record, _ int, b record) bool { return a.Key == b.Key },
		opiter.Enumerate(slices.Values(records)),
	) {
		fmt.Printf("%d: %v\n", i, v)
	}
	// Output:
	// 0: {foo yay}
	// 3: {bar yay}
	// 4: {baz yay}
}
