package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"gotest.tools/v3/assert"
)

func TestEqual(t *testing.T) {
	expected := slices.Values(slices.Clone(testhelper.SampleValues))
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) bool{
			func(seq iter.Seq[int]) bool { return opiter.Equal(seq, expected) },
			func(seq iter.Seq[int]) bool {
				return opiter.EqualFunc(func(a, b int) bool { return a == b }, seq, expected)
			},
		},
		true,
		nil,
	)
	expected2 := opiter.Values2(slices.Clone(testhelper.SamplePairs))
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) bool{
			func(seq iter.Seq2[string, int]) bool { return opiter.Equal2(seq, expected2) },
			func(seq iter.Seq2[string, int]) bool {
				return opiter.EqualFunc2(
					func(k1 string, v1 int, k2 string, v2 int) bool {
						return k1 == k2 && v1 == v2
					},
					seq,
					expected2,
				)
			},
		},
		true,
		nil,
	)
	assert.Assert(t, !opiter.Equal(slices.Values([]int{1, 2}), slices.Values([]int{1, 2, 3})))
	assert.Assert(t, !opiter.Equal(slices.Values([]int{1, 3}), slices.Values([]int{1, 2})))
	assert.Assert(t, opiter.Equal(slices.Values([]int{}), slices.Values([]int{})))
}

func TestSeqEqualFunc(t *testing.T) {
	seq := opiter.WrapSeq(testhelper.SampleSeq())
	otherValues := make([]int64, len(testhelper.SampleValues))
	for i, v := range testhelper.SampleValues {
		otherValues[i] = int64(v)
	}
	assert.Assert(t, seq.EqualFunc(func(a int, b int64) bool { return int64(a) == b }, slices.Values(otherValues)))

	seq2 := opiter.WrapSeq2(testhelper.SampleSeq2())
	otherPairs := make([]opiter.KV[string, int64], len(testhelper.SamplePairs))
	for i, kv := range testhelper.SamplePairs {
		otherPairs[i] = opiter.PackKV(kv.K, int64(kv.V))
	}
	assert.Assert(t, seq2.EqualFunc2(
		func(k1 string, v1 int, k2 string, v2 int64) bool {
			return k1 == k2 && int64(v1) == v2
		},
		opiter.Values2(otherPairs),
	))
}
