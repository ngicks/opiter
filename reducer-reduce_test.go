package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"gotest.tools/v3/assert"
)

func TestReduce(t *testing.T) {
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) int{
			func(seq iter.Seq[int]) int {
				return opiter.Reduce(func(sum, v int) int { return sum + v }, 0, seq)
			},
		},
		31,
		nil,
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) int{
			func(seq iter.Seq2[string, int]) int {
				return opiter.Reduce2(func(sum int, _ string, v int) int { return sum + v }, 0, seq)
			},
		},
		14,
		nil,
	)
	assert.Equal(t, 42, opiter.Reduce(func(sum, v int) int { return sum + v }, 42, slices.Values([]int{})))
}

func TestSeqReduce(t *testing.T) {
	seq := opiter.WrapSeq(testhelper.SampleSeq())
	got := seq.Reduce(func(sum string, v int) string {
		return sum + string(rune('0'+v))
	}, "")
	assert.Equal(t, "31415926", got)

	seq2 := opiter.WrapSeq2(testhelper.SampleSeq2())
	got2 := seq2.Reduce2(func(sum string, k string, v int) string {
		return sum + k + string(rune('0'+v))
	}, "")
	assert.Equal(t, "foo3bar1baz4foo1qux5", got2)
}
