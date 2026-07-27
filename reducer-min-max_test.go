package opiter_test

import (
	"cmp"
	"iter"
	"slices"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"github.com/ngicks/option/opt"
	"gotest.tools/v3/assert"
)

func TestMinMax(t *testing.T) {
	optionCmp := []goCmp.Option{compareOption[int]()}
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) opt.Option[int]{
			func(seq iter.Seq[int]) opt.Option[int] { return opiter.Min(seq) },
			func(seq iter.Seq[int]) opt.Option[int] { return opiter.MinFunc(cmp.Compare, seq) },
		},
		opt.Some(1),
		optionCmp,
	)
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) opt.Option[int]{
			func(seq iter.Seq[int]) opt.Option[int] { return opiter.Max(seq) },
			func(seq iter.Seq[int]) opt.Option[int] { return opiter.MaxFunc(cmp.Compare, seq) },
		},
		opt.Some(9),
		optionCmp,
	)
	assert.Assert(t, opiter.Min(slices.Values([]int{})).IsNone())
	assert.Assert(t, opiter.Max(slices.Values([]int{})).IsNone())
}

func TestSeqMinMaxFunc(t *testing.T) {
	seq := opiter.WrapSeq(testhelper.SampleSeq())
	assert.Equal(t, 1, seq.MinFunc(cmp.Compare).Value())
	assert.Equal(t, 9, seq.MaxFunc(cmp.Compare).Value())
}
