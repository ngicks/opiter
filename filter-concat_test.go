package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestConcat(t *testing.T) {
	input := slices.Values([]int{1, 2})
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(seq iter.Seq[int]) iter.Seq[int] { return opiter.Concat(seq, slices.Values([]int{3, 4})) },
		func(seq iter.Seq[int]) iter.Seq[int] {
			return opiter.Concat(opiter.WrapSeq(seq), opiter.WrapSeq(slices.Values([]int{3, 4})))
		},
		func(seq iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(seq).Concat(slices.Values([]int{3, 4})).Iter()
		},
		func(seq iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(seq).
				Concat(opiter.WrapSeq(slices.Values([]int{3})), opiter.WrapSeq(slices.Values([]int{4}))).
				Iter()
		},
	}
	testhelper.TestFilterStateless(t, input, filters, []int{1, 2, 3, 4}, nil)
}

func TestConcat2(t *testing.T) {
	input := opiter.Values2([]opiter.KV[int, int]{{1, 2}})
	other := opiter.Values2([]opiter.KV[int, int]{{3, 4}})
	filters := []func(iter.Seq2[int, int]) iter.Seq2[int, int]{
		func(seq iter.Seq2[int, int]) iter.Seq2[int, int] {
			return opiter.Concat2(seq, other)
		},
		func(seq iter.Seq2[int, int]) iter.Seq2[int, int] {
			return opiter.Concat2(opiter.WrapSeq2(seq), opiter.WrapSeq2(other))
		},
		func(seq iter.Seq2[int, int]) iter.Seq2[int, int] {
			return opiter.WrapSeq2(seq).Concat(other).Iter2()
		},
		func(seq iter.Seq2[int, int]) iter.Seq2[int, int] {
			return opiter.WrapSeq2(seq).Concat(opiter.WrapSeq2(other)).Iter2()
		},
	}
	expected := []opiter.KV[int, int]{{1, 2}, {3, 4}}
	testhelper.TestFilterStateless2(t, input, filters, expected, nil)
}
