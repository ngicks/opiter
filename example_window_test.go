package opiter_test

import (
	"fmt"
	"iter"
	"slices"

	"github.com/ngicks/opiter"
)

func ExampleWindow_moving_average() {
	data := []int{10, 20, 30, 40, 50, 60}
	averages := opiter.Map(
		func(window []int) float64 {
			return float64(opiter.Sum(slices.Values(window))) / float64(len(window))
		},
		opiter.Window(data, 3),
	)

	fmt.Printf("%.1f\n", slices.Collect(averages))
	// Output:
	// [20.0 30.0 40.0 50.0]
}

func ExampleWindowSeq_moving_average() {
	data := []int{10, 20, 30, 40, 50, 60}
	averages := opiter.Map(
		func(window iter.Seq[int]) float64 {
			return float64(opiter.Sum(window)) / 3
		},
		opiter.WindowSeq(3, slices.Values(data)),
	)

	fmt.Printf("%.1f\n", slices.Collect(averages))
	// Output:
	// [20.0 30.0 40.0 50.0]
}
