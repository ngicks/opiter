package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestWindow(t *testing.T) {
	got := slices.Collect(opiter.Window([]int{1, 2, 3, 4}, 2))
	expected := [][]int{{1, 2}, {2, 3}, {3, 4}}
	if !slices.EqualFunc(got, expected, slices.Equal) {
		t.Fatalf("Window() = %v", got)
	}
}

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
