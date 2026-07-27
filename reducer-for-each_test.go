package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"gotest.tools/v3/assert"
)

func TestForEachDiscard(t *testing.T) {
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) []int{
			func(seq iter.Seq[int]) (got []int) {
				opiter.ForEach(func(v int) { got = append(got, v) }, seq)
				return got
			},
			func(seq iter.Seq[int]) (got []int) {
				opiter.Discard(func(yield func(int) bool) {
					seq(func(v int) bool {
						got = append(got, v)
						return yield(v)
					})
				})
				return got
			},
		},
		testhelper.SampleValues,
		nil,
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) []opiter.KV[string, int]{
			func(seq iter.Seq2[string, int]) (got []opiter.KV[string, int]) {
				opiter.ForEach2(func(k string, v int) { got = append(got, opiter.PackKV(k, v)) }, seq)
				return got
			},
			func(seq iter.Seq2[string, int]) (got []opiter.KV[string, int]) {
				opiter.Discard2(func(yield func(string, int) bool) {
					seq(func(k string, v int) bool {
						got = append(got, opiter.PackKV(k, v))
						return yield(k, v)
					})
				})
				return got
			},
		},
		testhelper.SamplePairs,
		nil,
	)
}

func TestSeqForEachDiscard(t *testing.T) {
	var got []int
	opiter.WrapSeq(testhelper.SampleSeq()).ForEach(func(v int) { got = append(got, v) })
	assert.DeepEqual(t, testhelper.SampleValues, got)
	opiter.WrapSeq(testhelper.SampleSeq()).Discard()

	var got2 []opiter.KV[string, int]
	opiter.WrapSeq2(testhelper.SampleSeq2()).ForEach2(func(k string, v int) {
		got2 = append(got2, opiter.PackKV(k, v))
	})
	assert.DeepEqual(t, testhelper.SamplePairs, got2)
	opiter.WrapSeq2(testhelper.SampleSeq2()).Discard2()
}
