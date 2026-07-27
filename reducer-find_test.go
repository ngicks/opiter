package opiter_test

import (
	"iter"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"github.com/ngicks/option/opt"
	"gotest.tools/v3/assert"
)

func TestFind(t *testing.T) {
	optionCmp := []goCmp.Option{compareOption[opiter.KV[int, int]]()}
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) opt.Option[opiter.KV[int, int]]{
			func(seq iter.Seq[int]) opt.Option[opiter.KV[int, int]] { return opiter.Find(1, seq) },
			func(seq iter.Seq[int]) opt.Option[opiter.KV[int, int]] {
				return opiter.FindFunc(func(v int) bool { return v == 1 }, seq)
			},
		},
		opt.Some(opiter.PackKV(1, 1)),
		optionCmp,
	)
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) opt.Option[opiter.KV[int, int]]{
			func(seq iter.Seq[int]) opt.Option[opiter.KV[int, int]] { return opiter.FindLast(1, seq) },
			func(seq iter.Seq[int]) opt.Option[opiter.KV[int, int]] {
				return opiter.FindLastFunc(func(v int) bool { return v == 1 }, seq)
			},
		},
		opt.Some(opiter.PackKV(3, 1)),
		optionCmp,
	)

	type indexedPair = opiter.KV[int, opiter.KV[string, int]]
	pairOptionCmp := []goCmp.Option{compareOption[indexedPair]()}
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) opt.Option[indexedPair]{
			func(seq iter.Seq2[string, int]) opt.Option[indexedPair] {
				return opiter.Find2("foo", 1, seq)
			},
			func(seq iter.Seq2[string, int]) opt.Option[indexedPair] {
				return opiter.FindFunc2(func(k string, v int) bool { return k == "foo" && v == 1 }, seq)
			},
		},
		opt.Some(opiter.PackKV(3, opiter.PackKV("foo", 1))),
		pairOptionCmp,
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) opt.Option[indexedPair]{
			func(seq iter.Seq2[string, int]) opt.Option[indexedPair] {
				return opiter.FindLast2("foo", 3, seq)
			},
			func(seq iter.Seq2[string, int]) opt.Option[indexedPair] {
				return opiter.FindLastFunc2(func(k string, v int) bool { return k == "foo" && v == 3 }, seq)
			},
		},
		opt.Some(opiter.PackKV(0, opiter.PackKV("foo", 3))),
		pairOptionCmp,
	)
	assert.Assert(t, opiter.Find(8, testhelper.SampleSeq()).IsNone())
	assert.Assert(t, opiter.Find2("missing", 0, testhelper.SampleSeq2()).IsNone())
}

func TestSeqFindFunc(t *testing.T) {
	found := opiter.WrapSeq(testhelper.SampleSeq()).FindFunc(func(v int) bool { return v == 4 })
	assert.DeepEqual(t, opiter.PackKV(2, 4), found.Value())
	last := opiter.WrapSeq(testhelper.SampleSeq()).FindLastFunc(func(v int) bool { return v == 1 })
	assert.DeepEqual(t, opiter.PackKV(3, 1), last.Value())

	found2 := opiter.WrapSeq2(testhelper.SampleSeq2()).FindFunc2(
		func(k string, v int) bool { return k == "foo" && v == 1 },
	)
	assert.DeepEqual(t, opiter.PackKV(3, opiter.PackKV("foo", 1)), found2.Value())
	last2 := opiter.WrapSeq2(testhelper.SampleSeq2()).FindLastFunc2(
		func(k string, _ int) bool { return k == "foo" },
	)
	assert.DeepEqual(t, opiter.PackKV(3, opiter.PackKV("foo", 1)), last2.Value())
}

func TestFindShortCircuits(t *testing.T) {
	consumed := 0
	seq := func(yield func(int) bool) {
		for _, v := range testhelper.SampleValues {
			consumed++
			if !yield(v) {
				return
			}
		}
	}
	opiter.Find(4, seq)
	assert.Equal(t, 3, consumed)
}
