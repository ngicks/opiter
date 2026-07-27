package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"gotest.tools/v3/assert"
)

func TestEveryAny(t *testing.T) {
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) bool{
			func(seq iter.Seq[int]) bool { return opiter.Every(func(v int) bool { return v > 0 }, seq) },
			func(seq iter.Seq[int]) bool { return opiter.Any(func(v int) bool { return v == 4 }, seq) },
		},
		true,
		nil,
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) bool{
			func(seq iter.Seq2[string, int]) bool {
				return opiter.Every2(func(k string, v int) bool { return k != "" && v > 0 }, seq)
			},
			func(seq iter.Seq2[string, int]) bool {
				return opiter.Any2(func(k string, v int) bool { return k == "baz" && v == 4 }, seq)
			},
		},
		true,
		nil,
	)
	empty := slices.Values([]int{})
	assert.Assert(t, opiter.Every(func(int) bool { return false }, empty))
	assert.Assert(t, !opiter.Any(func(int) bool { return true }, empty))
}

func TestSeqEveryAny(t *testing.T) {
	assert.Assert(t, opiter.WrapSeq(testhelper.SampleSeq()).Every(func(v int) bool { return v > 0 }))
	assert.Assert(t, opiter.WrapSeq(testhelper.SampleSeq()).Any(func(v int) bool { return v == 4 }))
	assert.Assert(t, opiter.WrapSeq2(testhelper.SampleSeq2()).Every2(
		func(k string, v int) bool { return k != "" && v > 0 },
	))
	assert.Assert(t, opiter.WrapSeq2(testhelper.SampleSeq2()).Any2(
		func(k string, v int) bool { return k == "baz" && v == 4 },
	))
}

func TestEveryAnyShortCircuit(t *testing.T) {
	type testCase struct {
		name   string
		reduce func(iter.Seq[int])
	}
	tests := []testCase{
		{"Every", func(seq iter.Seq[int]) { opiter.Every(func(v int) bool { return v < 4 }, seq) }},
		{"Any", func(seq iter.Seq[int]) { opiter.Any(func(v int) bool { return v == 4 }, seq) }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			consumed := 0
			seq := func(yield func(int) bool) {
				for _, v := range testhelper.SampleValues {
					consumed++
					if !yield(v) {
						return
					}
				}
			}
			tc.reduce(seq)
			assert.Equal(t, 3, consumed)
		})
	}
}
