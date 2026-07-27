package opiter_test

import (
	"encoding/json"
	"fmt"
	"maps"

	"github.com/ngicks/opiter"
)

func ExampleReduceGroup() {
	first := map[int]string{
		0: "foo",
		1: "bar",
		2: "baz",
	}
	second := map[int]string{
		0: "foo",
		2: "zab",
		3: "good",
	}

	grouped := opiter.ReduceGroup(
		func(values []string, value string) []string { return append(values, value) },
		nil,
		opiter.Concat2(maps.All(first), maps.All(second)),
	)
	output, _ := json.MarshalIndent(grouped, "", "  ")
	fmt.Println(string(output))
	// Output:
	// {
	//   "0": [
	//     "foo",
	//     "foo"
	//   ],
	//   "1": [
	//     "bar"
	//   ],
	//   "2": [
	//     "baz",
	//     "zab"
	//   ],
	//   "3": [
	//     "good"
	//   ]
	// }
}
