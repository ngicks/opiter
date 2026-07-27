package opiter_test

import (
	"fmt"

	"github.com/ngicks/opiter"
)

func ExampleRange_prevent_off_by_one() {
	for i := range opiter.LimitUntil(
		func(i int) bool { return i < 50 },
		opiter.Map(
			func(i int) int { return i * 7 },
			opiter.Range(0, 10),
		),
	) {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(i)
	}
	// Output:
	// 0 7 14 21 28 35 42 49
}
