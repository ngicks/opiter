package opiter_test

import (
	"iter"
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
				return opiter.WrapSeq(seq).Collect()
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
				return opiter.WrapSeq2(seq).Collect2()
			},
		},
		testhelper.SamplePairs,
		nil,
	)
}

func TestSeqCollectEmpty(t *testing.T) {
	assert.Assert(t, opiter.WrapSeq(opiter.Empty[int]()).Collect() == nil)
	assert.Assert(t, opiter.WrapSeq2(opiter.Empty2[string, int]()).Collect2() == nil)
}
