package opiter_test

import (
	"iter"
	"slices"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"github.com/ngicks/option/opt"
	"gotest.tools/v3/assert"
)

func TestFirstLast(t *testing.T) {
	optionCmp := []goCmp.Option{compareOption[int]()}
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) opt.Option[int]{
			func(seq iter.Seq[int]) opt.Option[int] { return opiter.First(seq) },
		},
		opt.Some(3),
		optionCmp,
	)
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) opt.Option[int]{
			func(seq iter.Seq[int]) opt.Option[int] { return opiter.Last(seq) },
		},
		opt.Some(6),
		optionCmp,
	)
	pairCmp := []goCmp.Option{compareOption[opiter.KV[string, int]]()}
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) opt.Option[opiter.KV[string, int]]{
			func(seq iter.Seq2[string, int]) opt.Option[opiter.KV[string, int]] {
				return opiter.First2(seq)
			},
		},
		opt.Some(opiter.PackKV("foo", 3)),
		pairCmp,
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) opt.Option[opiter.KV[string, int]]{
			func(seq iter.Seq2[string, int]) opt.Option[opiter.KV[string, int]] {
				return opiter.Last2(seq)
			},
		},
		opt.Some(opiter.PackKV("qux", 5)),
		pairCmp,
	)
	assert.Assert(t, opiter.First(slices.Values([]int{})).IsNone())
	assert.Assert(t, opiter.Last(slices.Values([]int{})).IsNone())
}

func TestSeqFirstLast(t *testing.T) {
	seq := opiter.WrapSeq(testhelper.SampleSeq())
	assert.Equal(t, 3, seq.First().Value())
	assert.Equal(t, 6, seq.Last().Value())

	seq2 := opiter.WrapSeq2(testhelper.SampleSeq2())
	assert.DeepEqual(t, opiter.PackKV("foo", 3), seq2.First2().Value())
	assert.DeepEqual(t, opiter.PackKV("qux", 5), seq2.Last2().Value())
}

func TestFirstShortCircuits(t *testing.T) {
	consumed := 0
	seq := func(yield func(int) bool) {
		for _, v := range testhelper.SampleValues {
			consumed++
			if !yield(v) {
				return
			}
		}
	}
	opiter.First(seq)
	assert.Equal(t, 1, consumed)
}
