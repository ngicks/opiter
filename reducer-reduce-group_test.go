package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReduceGroup(t *testing.T) {
	expected := map[string]int{"foo": 4, "bar": 1, "baz": 4, "qux": 5}
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) map[string]int{
			func(seq iter.Seq2[string, int]) map[string]int {
				return opiter.ReduceGroup(func(sum, v int) int { return sum + v }, 0, seq)
			},
		},
		expected,
		nil,
	)
	expectedInserted := map[string]int{"existing": 10, "foo": 4, "bar": 1, "baz": 4, "qux": 5}
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) map[string]int{
			func(seq iter.Seq2[string, int]) map[string]int {
				return opiter.InsertReduceGroup(
					map[string]int{"existing": 10},
					func(sum, v int) int { return sum + v },
					0,
					seq,
				)
			},
		},
		expectedInserted,
		nil,
	)
	inserted := opiter.InsertReduceGroup(
		map[string]int{"existing": 10},
		func(sum, v int) int { return sum + v },
		0,
		testhelper.SampleSeq2(),
	)
	assert.Check(t, is.DeepEqual(expectedInserted, inserted))
}
