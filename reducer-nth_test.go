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

func TestNth(t *testing.T) {
	testhelper.TestReducer(
		t,
		testhelper.SampleSeq(),
		[]func(iter.Seq[int]) opt.Option[int]{
			func(seq iter.Seq[int]) opt.Option[int] { return opiter.Nth(2, seq) },
		},
		opt.Some(4),
		[]goCmp.Option{compareOption[int]()},
	)
	testhelper.TestReducer2(
		t,
		testhelper.SampleSeq2(),
		[]func(iter.Seq2[string, int]) opt.Option[opiter.KV[string, int]]{
			func(seq iter.Seq2[string, int]) opt.Option[opiter.KV[string, int]] {
				return opiter.Nth2(2, seq)
			},
		},
		opt.Some(opiter.PackKV("baz", 4)),
		[]goCmp.Option{compareOption[opiter.KV[string, int]]()},
	)
	assert.Assert(t, opiter.Nth(-1, testhelper.SampleSeq()).IsNone())
	assert.Assert(t, opiter.Nth(len(testhelper.SampleValues), testhelper.SampleSeq()).IsNone())
	assert.Assert(t, opiter.Nth2(-1, testhelper.SampleSeq2()).IsNone())
}

func TestSeqNth(t *testing.T) {
	assert.Equal(t, 4, opiter.WrapSeq(testhelper.SampleSeq()).Nth(2).Value())
	assert.DeepEqual(t, opiter.PackKV("baz", 4), opiter.WrapSeq2(testhelper.SampleSeq2()).Nth2(2).Value())
}
