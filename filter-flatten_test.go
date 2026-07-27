package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestFlattenVariants(t *testing.T) {
	nested := slices.Values([][]int{{1, 2}, {3}})
	testhelper.TestFilterStateless(t, nested,
		[]func(iter.Seq[[]int]) iter.Seq[int]{func(s iter.Seq[[]int]) iter.Seq[int] { return opiter.Flatten(s) }},
		[]int{1, 2, 3}, nil)

	inner := slices.Values([]iter.Seq[int]{slices.Values([]int{1, 2}), slices.Values([]int{3})})
	testhelper.TestFilterStateless(t, inner,
		[]func(iter.Seq[iter.Seq[int]]) iter.Seq[int]{func(s iter.Seq[iter.Seq[int]]) iter.Seq[int] { return opiter.FlattenSeq(s) }},
		[]int{1, 2, 3}, nil)

	inner2 := slices.Values([]iter.Seq2[int, int]{opiter.Values2([]opiter.KV[int, int]{{1, 2}}), opiter.Values2([]opiter.KV[int, int]{{3, 4}})})
	testhelper.TestFilterStateless1To2(t, inner2,
		[]func(iter.Seq[iter.Seq2[int, int]]) iter.Seq2[int, int]{func(s iter.Seq[iter.Seq2[int, int]]) iter.Seq2[int, int] { return opiter.FlattenSeq2(s) }},
		[]opiter.KV[int, int]{{1, 2}, {3, 4}}, nil)
}

func TestFlattenPairVariants(t *testing.T) {
	first := opiter.Values2([]opiter.KV[[]int, string]{{[]int{1, 2}, "x"}})
	testhelper.TestFilterStateless2(t, first,
		[]func(iter.Seq2[[]int, string]) iter.Seq2[int, string]{func(s iter.Seq2[[]int, string]) iter.Seq2[int, string] { return opiter.FlattenF(s) }},
		[]opiter.KV[int, string]{{1, "x"}, {2, "x"}}, nil)
	firstSeq := opiter.Values2([]opiter.KV[iter.Seq[int], string]{{slices.Values([]int{1, 2}), "x"}})
	testhelper.TestFilterStateless2(t, firstSeq,
		[]func(iter.Seq2[iter.Seq[int], string]) iter.Seq2[int, string]{func(s iter.Seq2[iter.Seq[int], string]) iter.Seq2[int, string] { return opiter.FlattenSeqF(s) }},
		[]opiter.KV[int, string]{{1, "x"}, {2, "x"}}, nil)
	last := opiter.Values2([]opiter.KV[string, []int]{{"x", []int{1, 2}}})
	testhelper.TestFilterStateless2(t, last,
		[]func(iter.Seq2[string, []int]) iter.Seq2[string, int]{func(s iter.Seq2[string, []int]) iter.Seq2[string, int] { return opiter.FlattenL(s) }},
		[]opiter.KV[string, int]{{"x", 1}, {"x", 2}}, nil)
	lastSeq := opiter.Values2([]opiter.KV[string, iter.Seq[int]]{{"x", slices.Values([]int{1, 2})}})
	testhelper.TestFilterStateless2(t, lastSeq,
		[]func(iter.Seq2[string, iter.Seq[int]]) iter.Seq2[string, int]{func(s iter.Seq2[string, iter.Seq[int]]) iter.Seq2[string, int] { return opiter.FlattenSeqL(s) }},
		[]opiter.KV[string, int]{{"x", 1}, {"x", 2}}, nil)
}

func TestFlattenVariantsStateful(t *testing.T) {
	nested := slices.Values([][]int{{1}, {2}, {3}})
	testhelper.TestFilterStateful(t, nested,
		[]func(iter.Seq[[]int]) iter.Seq[int]{
			func(s iter.Seq[[]int]) iter.Seq[int] { return opiter.Flatten(s) },
		},
		[]int{1, 2, 3}, nil)

	inner := slices.Values([]iter.Seq[int]{
		slices.Values([]int{1}),
		slices.Values([]int{2}),
		slices.Values([]int{3}),
	})
	testhelper.TestFilterStateful(t, inner,
		[]func(iter.Seq[iter.Seq[int]]) iter.Seq[int]{
			func(s iter.Seq[iter.Seq[int]]) iter.Seq[int] { return opiter.FlattenSeq(s) },
		},
		[]int{1, 2, 3}, nil)

	inner2 := slices.Values([]iter.Seq2[int, int]{
		opiter.Values2([]opiter.KV[int, int]{{1, 2}}),
		opiter.Values2([]opiter.KV[int, int]{{3, 4}}),
	})
	testhelper.TestFilterStateful1To2(t, inner2,
		[]func(iter.Seq[iter.Seq2[int, int]]) iter.Seq2[int, int]{
			func(s iter.Seq[iter.Seq2[int, int]]) iter.Seq2[int, int] { return opiter.FlattenSeq2(s) },
		},
		[]opiter.KV[int, int]{{1, 2}, {3, 4}}, nil)
}

func TestFlattenPairVariantsStateful(t *testing.T) {
	first := opiter.Values2([]opiter.KV[[]int, string]{{[]int{1}, "x"}, {[]int{2}, "y"}})
	testhelper.TestFilterStateful2(t, first,
		[]func(iter.Seq2[[]int, string]) iter.Seq2[int, string]{
			func(s iter.Seq2[[]int, string]) iter.Seq2[int, string] { return opiter.FlattenF(s) },
		},
		[]opiter.KV[int, string]{{1, "x"}, {2, "y"}}, nil)

	firstSeq := opiter.Values2([]opiter.KV[iter.Seq[int], string]{
		{slices.Values([]int{1}), "x"},
		{slices.Values([]int{2}), "y"},
	})
	testhelper.TestFilterStateful2(t, firstSeq,
		[]func(iter.Seq2[iter.Seq[int], string]) iter.Seq2[int, string]{
			func(s iter.Seq2[iter.Seq[int], string]) iter.Seq2[int, string] { return opiter.FlattenSeqF(s) },
		},
		[]opiter.KV[int, string]{{1, "x"}, {2, "y"}}, nil)

	last := opiter.Values2([]opiter.KV[string, []int]{{"x", []int{1}}, {"y", []int{2}}})
	testhelper.TestFilterStateful2(t, last,
		[]func(iter.Seq2[string, []int]) iter.Seq2[string, int]{
			func(s iter.Seq2[string, []int]) iter.Seq2[string, int] { return opiter.FlattenL(s) },
		},
		[]opiter.KV[string, int]{{"x", 1}, {"y", 2}}, nil)

	lastSeq := opiter.Values2([]opiter.KV[string, iter.Seq[int]]{
		{"x", slices.Values([]int{1})},
		{"y", slices.Values([]int{2})},
	})
	testhelper.TestFilterStateful2(t, lastSeq,
		[]func(iter.Seq2[string, iter.Seq[int]]) iter.Seq2[string, int]{
			func(s iter.Seq2[string, iter.Seq[int]]) iter.Seq2[string, int] {
				return opiter.FlattenSeqL(s)
			},
		},
		[]opiter.KV[string, int]{{"x", 1}, {"y", 2}}, nil)
}
