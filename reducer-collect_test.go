package opiter_test

import (
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"gotest.tools/v3/assert"
)

func TestSeqCollect(t *testing.T) {
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) []int{
			func(seq iter.Seq[int]) []int {
				return opiter.WrapSeq(seq).Collect(slices.Sorted)
			},
			func(seq iter.Seq[int]) []int {
				return opiter.WrapSeq(seq).Collect(func(seq opiter.Seq[int]) []int {
					return slices.Sorted(seq.Iter())
				})
			},
		},
		slices.Sorted(testhelper.SampleSeq()),
		nil,
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) map[string]int{
			func(seq iter.Seq2[string, int]) map[string]int {
				return opiter.WrapSeq2(seq).Collect(maps.Collect)
			},
			func(seq iter.Seq2[string, int]) map[string]int {
				return opiter.WrapSeq2(seq).Collect(
					func(seq opiter.Seq2[string, int]) map[string]int {
						return maps.Collect(seq.Iter2())
					},
				)
			},
		},
		maps.Collect(testhelper.SampleSeq2()),
		nil,
	)
}

func TestSeqCollectSlice(t *testing.T) {
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) []int{
			func(seq iter.Seq[int]) []int {
				return opiter.WrapSeq(seq).CollectSlice()
			},
		},
		testhelper.SampleValues,
		nil,
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) []opiter.KV[string, int]{
			func(seq iter.Seq2[string, int]) []opiter.KV[string, int] {
				return opiter.WrapSeq2(seq).CollectSlice()
			},
		},
		testhelper.SamplePairs,
		nil,
	)
}

func TestSeqCollectSliceEmpty(t *testing.T) {
	assert.Assert(t, opiter.WrapSeq(opiter.Empty[int]()).CollectSlice() == nil)
	assert.Assert(t, opiter.WrapSeq2(opiter.Empty2[string, int]()).CollectSlice() == nil)
}
