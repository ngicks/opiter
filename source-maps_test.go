package opiter_test

import (
	"cmp"
	"iter"
	"slices"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestMapsKeys(t *testing.T) {
	m := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{func() iter.Seq2[string, int] {
			return opiter.MapsKeys(m, slices.Values([]string{"baz", "missing", "foo"}))
		}},
		[]opiter.KV[string, int]{
			{K: "baz", V: 3},
			{K: "missing", V: 0},
			{K: "foo", V: 1},
		},
		nil,
	)
}

func TestMapsSorted(t *testing.T) {
	m := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{func() iter.Seq2[string, int] {
			return opiter.MapsSorted(m)
		}},
		[]opiter.KV[string, int]{
			{K: "bar", V: 2},
			{K: "baz", V: 3},
			{K: "foo", V: 1},
		},
		nil,
	)
}

func TestMapsSortedFunc(t *testing.T) {
	m := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{func() iter.Seq2[string, int] {
			return opiter.MapsSortedFunc(m, func(a, b string) int {
				return cmp.Compare(m[a], m[b])
			})
		}},
		[]opiter.KV[string, int]{
			{K: "foo", V: 1},
			{K: "bar", V: 2},
			{K: "baz", V: 3},
		},
		nil,
	)
}

func TestMapsOverlayContract(t *testing.T) {
	cmpOpt := []goCmp.Option{
		cmpopts.SortSlices(func(a, b opiter.KV[string, int]) bool { return a.K < b.K }),
	}
	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{func() iter.Seq2[string, int] {
			return opiter.MapsOverlay(
				map[string]int{"foo": 1},
				map[string]int{"bar": 2},
				map[string]int{"foo": 3},
			)
		}},
		[]opiter.KV[string, int]{
			{K: "foo", V: 3},
			{K: "bar", V: 2},
		},
		cmpOpt,
	)
}

func TestMapsOverlay(t *testing.T) {
	got := opiter.Collect2(opiter.MapsOverlay(
		map[string]int{"foo": 1, "bar": 2, "baz": 3},
		map[string]int{"foo": 4, "qux": 5},
		map[string]int{"bar": 6, "quux": 7},
	))
	expected := []opiter.KV[string, int]{
		{K: "foo", V: 4},
		{K: "bar", V: 6},
		{K: "baz", V: 3},
		{K: "qux", V: 5},
		{K: "quux", V: 7},
	}
	if diff := goCmp.Diff(
		expected,
		got,
		cmpopts.SortSlices(func(a, b opiter.KV[string, int]) bool { return a.K < b.K }),
	); diff != "" {
		t.Errorf("MapsOverlay mismatch (-want +got):\n%s", diff)
	}
}
