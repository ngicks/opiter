package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestWindowSeq(t *testing.T) {
	filters := []func(iter.Seq[int]) iter.Seq[[]int]{
		func(s iter.Seq[int]) iter.Seq[[]int] {
			return opiter.Map(slices.Collect, opiter.WindowSeq(3, s))
		},
		func(s iter.Seq[int]) iter.Seq[[]int] {
			return opiter.Map(slices.Collect, opiter.WrapSeq(s).Window(3))
		},
	}
	expected := [][]int{{1, 2, 3}, {2, 3, 4}}
	testhelper.TestFilterStateless(t, slices.Values([]int{1, 2, 3, 4}), filters, expected, nil)
}

func TestWindowSeqNonPositive(t *testing.T) {
	filters := []func(iter.Seq[int]) iter.Seq[[]int]{
		func(s iter.Seq[int]) iter.Seq[[]int] {
			return opiter.Map(slices.Collect, opiter.WindowSeq(0, s))
		},
		func(s iter.Seq[int]) iter.Seq[[]int] {
			return opiter.Map(slices.Collect, opiter.WrapSeq(s).Window(-1))
		},
	}
	testhelper.TestFilterStateless(t, testhelper.SampleSeq(), filters, nil, nil)
}

func TestWindowSeqPartialWindow(t *testing.T) {
	head := func(window iter.Seq[int]) int {
		for v := range window {
			return v
		}
		return 0
	}
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.Map(head, opiter.WindowSeq(3, s)) },
	}
	testhelper.TestFilterStateless(t, slices.Values([]int{1, 2, 3, 4}), filters, []int{1, 2}, nil)
}
