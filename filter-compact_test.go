package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestCompactVariants(t *testing.T) {
	input := slices.Values([]int{1, 1, 2, 2, 1})
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.Compact(s) },
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.CompactFunc(func(x, y int) bool { return x == y }, s)
		},
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).CompactFunc(func(x, y int) bool { return x == y }).Iter()
		},
	}
	testhelper.TestFilterStateless(t, input, filters, []int{1, 2, 1}, nil)
}

func TestCompactVariants2(t *testing.T) {
	input := opiter.Values2([]opiter.KV[int, int]{{1, 1}, {1, 1}, {2, 2}})
	eq := func(k1, v1, k2, v2 int) bool { return k1 == k2 && v1 == v2 }
	filters := []func(iter.Seq2[int, int]) iter.Seq2[int, int]{
		func(s iter.Seq2[int, int]) iter.Seq2[int, int] { return opiter.Compact2(s) },
		func(s iter.Seq2[int, int]) iter.Seq2[int, int] { return opiter.CompactFunc2(eq, s) },
		func(s iter.Seq2[int, int]) iter.Seq2[int, int] { return opiter.WrapSeq2(s).CompactFunc(eq).Iter2() },
	}
	expected := []opiter.KV[int, int]{{1, 1}, {2, 2}}
	testhelper.TestFilterStateless2(t, input, filters, expected, nil)
}
