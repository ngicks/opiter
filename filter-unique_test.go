package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestUniqueVariants(t *testing.T) {
	input := slices.Values([]int{1, 2, 1, 3, 2})
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.Unique(s) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.UniqueFunc(func(v int) int { return v }, s) },
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).UniqueFunc(func(v int) int { return v }).Iter()
		},
	}
	testhelper.TestFilterStateless(t, input, filters, []int{1, 2, 3}, nil)
}

func TestUniqueVariants2(t *testing.T) {
	input := opiter.Values2([]opiter.KV[int, int]{{1, 1}, {1, 1}, {2, 2}})
	filters := []func(iter.Seq2[int, int]) iter.Seq2[int, int]{
		func(s iter.Seq2[int, int]) iter.Seq2[int, int] { return opiter.Unique2(s) },
		func(s iter.Seq2[int, int]) iter.Seq2[int, int] {
			return opiter.UniqueFunc2(func(k, v int) opiter.KV[int, int] { return opiter.PackKV(k, v) }, s)
		},
		func(s iter.Seq2[int, int]) iter.Seq2[int, int] {
			return opiter.WrapSeq2(s).UniqueFunc(func(k, v int) opiter.KV[int, int] { return opiter.PackKV(k, v) }).Iter2()
		},
	}
	testhelper.TestFilterStateless2(t, input, filters, []opiter.KV[int, int]{{1, 1}, {2, 2}}, nil)
}
