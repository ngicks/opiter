package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"gotest.tools/v3/assert"
)

func TestSum(t *testing.T) {
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) int{
			func(seq iter.Seq[int]) int { return opiter.Sum(seq) },
			func(seq iter.Seq[int]) int {
				return opiter.SumOf(func(v int) int { return v }, seq)
			},
		},
		31,
		nil,
	)
	assert.Equal(t, "", opiter.Sum(slices.Values([]string{})))
}

func TestSeqSumOf(t *testing.T) {
	got := opiter.WrapSeq(testhelper.SampleSeq()).SumOf(func(v int) int64 { return int64(v) })
	assert.Equal(t, int64(31), got)
}
