package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestTranslateSeqShapes(t *testing.T) {
	testhelper.TestFilterStateless1To2(t, slices.Values([]int{7, 8}),
		[]func(iter.Seq[int]) iter.Seq2[int, int]{
			func(s iter.Seq[int]) iter.Seq2[int, int] { return opiter.Enumerate(s) },
			func(s iter.Seq[int]) iter.Seq2[int, int] { return opiter.WrapSeq(s).Enumerate().Iter2() },
		}, []opiter.KV[int, int]{{0, 7}, {1, 8}}, nil)
	testhelper.TestFilterStateless1To2(t, slices.Values([]int{1, 2}),
		[]func(iter.Seq[int]) iter.Seq2[int, int]{
			func(s iter.Seq[int]) iter.Seq2[int, int] { return opiter.Pairs(s, slices.Values([]int{3, 4})) },
			func(s iter.Seq[int]) iter.Seq2[int, int] {
				return opiter.WrapSeq(s).Pairs(slices.Values([]int{3, 4})).Iter2()
			},
		},
		[]opiter.KV[int, int]{{1, 3}, {2, 4}}, nil)
	testhelper.TestFilterStateless1To2(t, slices.Values([]int{12, 34}),
		[]func(iter.Seq[int]) iter.Seq2[int, int]{
			func(s iter.Seq[int]) iter.Seq2[int, int] {
				return opiter.Divide(func(v int) (int, int) { return v / 10, v % 10 }, s)
			},
			func(s iter.Seq[int]) iter.Seq2[int, int] {
				return opiter.WrapSeq(s).Divide(func(v int) (int, int) { return v / 10, v % 10 }).Iter2()
			},
		},
		[]opiter.KV[int, int]{{1, 2}, {3, 4}}, nil)
}

func TestTranslatePairShapes(t *testing.T) {
	input := opiter.Values2([]opiter.KV[int, string]{{1, "a"}, {2, "b"}})
	testhelper.TestFilterStateless2(t, input,
		[]func(iter.Seq2[int, string]) iter.Seq2[string, int]{
			func(s iter.Seq2[int, string]) iter.Seq2[string, int] { return opiter.Transpose(s) },
			func(s iter.Seq2[int, string]) iter.Seq2[string, int] { return opiter.WrapSeq2(s).Transpose().Iter2() },
		}, []opiter.KV[string, int]{{"a", 1}, {"b", 2}}, nil)
	testhelper.TestFilterStateless2To1(t, input,
		[]func(iter.Seq2[int, string]) iter.Seq[int]{
			func(s iter.Seq2[int, string]) iter.Seq[int] { return opiter.OmitL(s) },
			func(s iter.Seq2[int, string]) iter.Seq[int] { return opiter.WrapSeq2(s).OmitL().Iter() },
		}, []int{1, 2}, nil)
	testhelper.TestFilterStateless2To1(t, input,
		[]func(iter.Seq2[int, string]) iter.Seq[string]{
			func(s iter.Seq2[int, string]) iter.Seq[string] { return opiter.OmitF(s) },
			func(s iter.Seq2[int, string]) iter.Seq[string] {
				return opiter.Unify(func(k int, v string) string { return v }, s)
			},
			func(s iter.Seq2[int, string]) iter.Seq[string] {
				return opiter.WrapSeq2(s).Unify(func(k int, v string) string { return v }).Iter()
			},
		}, []string{"a", "b"}, nil)
}

func TestPairs2AndOmit(t *testing.T) {
	got := opiter.Collect2(opiter.Pairs2(
		opiter.Values2([]opiter.KV[int, int]{{1, 2}}),
		opiter.Values2([]opiter.KV[int, int]{{3, 4}}),
	))
	if !slices.Equal(got, []opiter.KV[opiter.KV[int, int], opiter.KV[int, int]]{{opiter.PackKV(1, 2), opiter.PackKV(3, 4)}}) {
		t.Fatalf("Pairs2() = %v", got)
	}
	methodGot := opiter.Collect2(opiter.WrapSeq2(
		opiter.Values2([]opiter.KV[int, int]{{1, 2}}),
	).Pairs(opiter.Values2([]opiter.KV[int, int]{{3, 4}})))
	if !slices.Equal(methodGot, got) {
		t.Fatalf("Pairs method = %v", methodGot)
	}
	count := 0
	for range opiter.WrapSeq(slices.Values([]int{1, 2})).Omit() {
		count++
	}
	for range opiter.WrapSeq2(opiter.Values2([]opiter.KV[int, int]{{1, 2}})).Omit() {
		count++
	}
	if count != 3 {
		t.Fatalf("omit count = %d", count)
	}
}
