package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"github.com/ngicks/option/tuple"
	"gotest.tools/v3/assert"
)

func TestKV(t *testing.T) {
	kv := opiter.PackKV("foo", 3)

	k, v := kv.Unpack()
	assert.Equal(t, k, "foo")
	assert.Equal(t, v, 3)

	assert.Equal(t, kv.ToTuple2(), tuple.Tuple2[string, int]{V1: "foo", V2: 3})
}

func TestAssembleKV(t *testing.T) {
	filters := []func(iter.Seq2[string, int]) iter.Seq[opiter.KV[string, int]]{
		func(s iter.Seq2[string, int]) iter.Seq[opiter.KV[string, int]] {
			return opiter.AssembleKV(s)
		},
	}
	testhelper.TestFilterStateless2To1(
		t,
		testhelper.SampleSeq2(),
		filters,
		testhelper.SamplePairs,
		nil,
	)
	testhelper.TestFilterStateful2To1(
		t,
		testhelper.SampleSeq2(),
		filters,
		testhelper.SamplePairs,
		nil,
	)
}

func TestDisassembleKV(t *testing.T) {
	filters := []func(iter.Seq[opiter.KV[string, int]]) iter.Seq2[string, int]{
		func(s iter.Seq[opiter.KV[string, int]]) iter.Seq2[string, int] {
			return opiter.DisassembleKV(s)
		},
	}
	input := opiter.AssembleKV(testhelper.SampleSeq2())
	testhelper.TestFilterStateless1To2(t, input, filters, testhelper.SamplePairs, nil)
	testhelper.TestFilterStateful1To2(t, input, filters, testhelper.SamplePairs, nil)
}

func TestKVPairs(t *testing.T) {
	pairs := opiter.KVPairs[string, int](testhelper.SamplePairs)
	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{pairs.Iter2},
		testhelper.SamplePairs,
		nil,
	)
}
