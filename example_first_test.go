package opiter_test

import (
	"fmt"
	"slices"

	"github.com/ngicks/opiter"
)

func ExampleSeq_First() {
	first := opiter.WrapSeq(slices.Values([]string{"alpha", "beta"})).First()
	missing := opiter.WrapSeq(slices.Values([]string{})).First()

	fmt.Println(first.Value())
	fmt.Println(missing.IsNone())
	// Output:
	// alpha
	// true
}
