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

func TestCount(t *testing.T) {
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) int{
			func(seq iter.Seq[int]) int { return opiter.Count(seq) },
			func(seq iter.Seq[int]) int { return opiter.WrapSeq(seq).Count() },
		},
		8,
		nil,
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) int{
			func(seq iter.Seq2[string, int]) int { return opiter.Count2(seq) },
			func(seq iter.Seq2[string, int]) int { return opiter.WrapSeq2(seq).Count() },
		},
		5,
		nil,
	)
	assert.Equal(t, 0, opiter.Count(slices.Values([]int{})))
	assert.Equal(t, 0, opiter.Count2(maps.All(map[string]int{})))
}
