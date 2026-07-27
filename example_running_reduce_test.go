package opiter_test

import (
	"fmt"
	"math"

	"github.com/ngicks/opiter"
)

func ExampleRunningReduce() {
	for i, sum := range opiter.Enumerate(
		opiter.RunningReduce(
			func(accum, next int64, _ int) int64 { return accum + next },
			int64(0),
			opiter.Range[int64](1, math.MaxInt64),
		),
	) {
		if i >= 5 {
			break
		}
		fmt.Println(sum)
	}
	// Output:
	// 1
	// 3
	// 6
	// 10
	// 15
}
