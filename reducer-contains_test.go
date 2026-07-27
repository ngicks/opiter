package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"gotest.tools/v3/assert"
)

func TestContains(t *testing.T) {
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) bool{
			func(seq iter.Seq[int]) bool { return opiter.Contains(4, seq) },
			func(seq iter.Seq[int]) bool {
				return opiter.ContainsFunc(func(v int) bool { return v == 4 }, seq)
			},
		},
		true,
		nil,
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) bool{
			func(seq iter.Seq2[string, int]) bool { return opiter.Contains2("baz", 4, seq) },
			func(seq iter.Seq2[string, int]) bool {
				return opiter.ContainsFunc2(func(k string, v int) bool { return k == "baz" && v == 4 }, seq)
			},
		},
		true,
		nil,
	)
	assert.Assert(t, !opiter.Contains(8, testhelper.SampleSeq()))
}

func TestSeqContainsFunc(t *testing.T) {
	assert.Assert(t, opiter.WrapSeq(testhelper.SampleSeq()).ContainsFunc(func(v int) bool { return v == 4 }))
	assert.Assert(t, opiter.WrapSeq2(testhelper.SampleSeq2()).ContainsFunc2(
		func(k string, v int) bool { return k == "baz" && v == 4 },
	))
}
