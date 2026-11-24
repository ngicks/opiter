package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func flatMapFilters(n int) []func(iter.Seq[int]) iter.Seq[int] {
	// Odd values map to an empty iterator.
	repeatEven := func(v int) iter.Seq[int] {
		if v%2 != 0 {
			return opiter.Empty[int]().Iter()
		}
		return opiter.Repeat(v, n).Iter()
	}
	return []func(iter.Seq[int]) iter.Seq[int]{
		func(seq iter.Seq[int]) iter.Seq[int] {
			return opiter.FlatMap(repeatEven, seq)
		},
		func(seq iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(seq).FlatMap(repeatEven).Iter()
		},
		func(seq iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(seq).FlatMap(func(v int) opiter.Seq[int] {
				return opiter.WrapSeq(repeatEven(v))
			}).Iter()
		},
	}
}

func TestFlatMap(t *testing.T) {
	testhelper.TestFilterStateless(
		t,
		testhelper.SampleSeq(),
		flatMapFilters(2),
		[]int{4, 4, 2, 2, 6, 6},
		nil,
	)
}

// A break inside an inner iterator drops the rest of that inner iterator,
// so only inner iterators of at most one value follow the stateful contract.
func TestFlatMapStateful(t *testing.T) {
	testhelper.TestFilterStateful(
		t,
		testhelper.SampleSeq(),
		flatMapFilters(1),
		[]int{4, 2, 6},
		nil,
	)
}

func flatMap2Filters(n int) []func(iter.Seq2[string, int]) iter.Seq2[int, string] {
	// Odd values map to an empty iterator.
	repeatEven := func(k string, v int) iter.Seq2[int, string] {
		if v%2 != 0 {
			return opiter.Empty2[int, string]().Iter2()
		}
		return opiter.Repeat2(v, k, n).Iter2()
	}
	return []func(iter.Seq2[string, int]) iter.Seq2[int, string]{
		func(seq iter.Seq2[string, int]) iter.Seq2[int, string] {
			return opiter.FlatMap2(repeatEven, seq)
		},
		func(seq iter.Seq2[string, int]) iter.Seq2[int, string] {
			return opiter.WrapSeq2(seq).FlatMap(repeatEven).Iter2()
		},
		func(seq iter.Seq2[string, int]) iter.Seq2[int, string] {
			return opiter.WrapSeq2(seq).FlatMap(func(k string, v int) opiter.Seq2[int, string] {
				return opiter.WrapSeq2(repeatEven(k, v))
			}).Iter2()
		},
	}
}

func flatMap2Input() iter.Seq2[string, int] {
	return opiter.Values2([]opiter.KV[string, int]{
		{"foo", 3}, {"bar", 4}, {"baz", 1}, {"qux", 2}, {"quux", 5},
	}).Iter2()
}

func TestFlatMap2(t *testing.T) {
	testhelper.TestFilterStateless2(
		t,
		flatMap2Input(),
		flatMap2Filters(2),
		[]opiter.KV[int, string]{{4, "bar"}, {4, "bar"}, {2, "qux"}, {2, "qux"}},
		nil,
	)
}

func TestFlatMap2Stateful(t *testing.T) {
	testhelper.TestFilterStateful2(
		t,
		flatMap2Input(),
		flatMap2Filters(1),
		[]opiter.KV[int, string]{{4, "bar"}, {2, "qux"}},
		nil,
	)
}
