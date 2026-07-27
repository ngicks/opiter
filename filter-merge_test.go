package opiter_test

import (
	"cmp"
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestMergeVariants(t *testing.T) {
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.Merge(s, slices.Values([]int{2, 4, 6})) },
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.MergeFunc(cmp.Compare[int], s, slices.Values([]int{2, 4, 6}))
		},
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).MergeFunc(cmp.Compare[int], slices.Values([]int{2, 4, 6})).Iter()
		},
	}
	input := slices.Values([]int{1, 3, 5})
	expected := []int{1, 2, 3, 4, 5, 6}
	testhelper.TestFilterStateless(t, input, filters, expected, nil)
}

func TestMergeVariants2(t *testing.T) {
	right := opiter.Values2([]opiter.KV[int, string]{{2, "b"}, {4, "d"}})
	filters := []func(iter.Seq2[int, string]) iter.Seq2[int, string]{
		func(s iter.Seq2[int, string]) iter.Seq2[int, string] { return opiter.Merge2(s, right) },
		func(s iter.Seq2[int, string]) iter.Seq2[int, string] {
			return opiter.MergeFunc2(cmp.Compare[int], s, right)
		},
		func(s iter.Seq2[int, string]) iter.Seq2[int, string] {
			return opiter.WrapSeq2(s).MergeFunc(cmp.Compare[int], right).Iter2()
		},
	}
	input := opiter.Values2([]opiter.KV[int, string]{{1, "a"}, {3, "c"}})
	expected := []opiter.KV[int, string]{{1, "a"}, {2, "b"}, {3, "c"}, {4, "d"}}
	testhelper.TestFilterStateless2(t, input, filters, expected, nil)
}

func TestMergeVariantsStateful(t *testing.T) {
	empty := slices.Values([]int(nil))
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.Merge(s, empty) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.MergeFunc(cmp.Compare[int], s, empty) },
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).MergeFunc(cmp.Compare[int], empty).Iter()
		},
	}
	testhelper.TestFilterStateful(t, testhelper.SampleSeq(), filters, testhelper.SampleValues, nil)
}

func TestMergeVariants2Stateful(t *testing.T) {
	empty := opiter.Values2([]opiter.KV[string, int](nil))
	filters := []func(iter.Seq2[string, int]) iter.Seq2[string, int]{
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.Merge2(s, empty) },
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.MergeFunc2(cmp.Compare[string], s, empty)
		},
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).MergeFunc(cmp.Compare[string], empty).Iter2()
		},
	}
	testhelper.TestFilterStateful2(t, testhelper.SampleSeq2(), filters, testhelper.SamplePairs, nil)
}
