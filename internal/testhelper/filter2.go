package testhelper

import (
	"fmt"
	"iter"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
)

func testFilterStateless2One[K, V, K2, V2 any](
	t *testing.T,
	input iter.Seq2[K, V],
	filter func(iter.Seq2[K, V]) iter.Seq2[K2, V2],
	expected []opiter.KV[K2, V2],
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	testSourceStateless2One(t, func() iter.Seq2[K2, V2] { return filter(input) }, expected, cmpOpt)

	for cut := range len(expected) {
		var stopped bool
		upstream := iter.Seq2[K, V](func(yield func(K, V) bool) {
			input(func(k K, v V) bool {
				if stopped {
					t.Errorf("break at %d: filter consumed upstream after downstream break", cut)
					return false
				}
				return yield(k, v)
			})
		})
		n := 0
		for range filter(upstream) {
			if n >= cut {
				stopped = true
				break
			}
			n++
		}
	}
}

func testFilterStateful2One[K, V, K2, V2 any](
	t *testing.T,
	input iter.Seq2[K, V],
	filter func(iter.Seq2[K, V]) iter.Seq2[K2, V2],
	expected []opiter.KV[K2, V2],
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	pairs := opiter.Collect2(input)
	testSourceStateful2One(t, func() iter.Seq2[K2, V2] { return filter(OneShotSeq2(pairs...)) }, expected, cmpOpt)
}

func testFilterStateless2To1One[K, V, U any](
	t *testing.T,
	input iter.Seq2[K, V],
	filter func(iter.Seq2[K, V]) iter.Seq[U],
	expected []U,
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	testSourceStatelessOne(t, func() iter.Seq[U] { return filter(input) }, expected, cmpOpt)

	for cut := range len(expected) {
		var stopped bool
		upstream := iter.Seq2[K, V](func(yield func(K, V) bool) {
			input(func(k K, v V) bool {
				if stopped {
					t.Errorf("break at %d: filter consumed upstream after downstream break", cut)
					return false
				}
				return yield(k, v)
			})
		})
		n := 0
		for range filter(upstream) {
			if n >= cut {
				stopped = true
				break
			}
			n++
		}
	}
}

func testFilterStateful2To1One[K, V, U any](
	t *testing.T,
	input iter.Seq2[K, V],
	filter func(iter.Seq2[K, V]) iter.Seq[U],
	expected []U,
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	pairs := opiter.Collect2(input)
	testSourceStatefulOne(t, func() iter.Seq[U] { return filter(OneShotSeq2(pairs...)) }, expected, cmpOpt)
}

func testFilterStateless2[K, V, K2, V2 any](
	t *testing.T,
	input iter.Seq2[K, V],
	filters []func(iter.Seq2[K, V]) iter.Seq2[K2, V2],
	expected []opiter.KV[K2, V2],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, filter := range filters {
		t.Run(fmt.Sprintf("filters[%d]", i), func(t *testing.T) {
			testFilterStateless2One(t, input, filter, expected, cmpOpt)
		})
	}
}

func testFilterStateful2[K, V, K2, V2 any](
	t *testing.T,
	input iter.Seq2[K, V],
	filters []func(iter.Seq2[K, V]) iter.Seq2[K2, V2],
	expected []opiter.KV[K2, V2],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, filter := range filters {
		t.Run(fmt.Sprintf("filters[%d]", i), func(t *testing.T) {
			testFilterStateful2One(t, input, filter, expected, cmpOpt)
		})
	}
}

func testFilterStateless2To1[K, V, U any](
	t *testing.T,
	input iter.Seq2[K, V],
	filters []func(iter.Seq2[K, V]) iter.Seq[U],
	expected []U,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, filter := range filters {
		t.Run(fmt.Sprintf("filters[%d]", i), func(t *testing.T) {
			testFilterStateless2To1One(t, input, filter, expected, cmpOpt)
		})
	}
}

func testFilterStateful2To1[K, V, U any](
	t *testing.T,
	input iter.Seq2[K, V],
	filters []func(iter.Seq2[K, V]) iter.Seq[U],
	expected []U,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, filter := range filters {
		t.Run(fmt.Sprintf("filters[%d]", i), func(t *testing.T) {
			testFilterStateful2To1One(t, input, filter, expected, cmpOpt)
		})
	}
}

// TestFilterStateless2 is [TestFilterStateless] for iter.Seq2 -> iter.Seq2 filters.
func TestFilterStateless2[K, V, K2, V2 any](
	t *testing.T,
	input iter.Seq2[K, V],
	filters []func(iter.Seq2[K, V]) iter.Seq2[K2, V2],
	expected []opiter.KV[K2, V2],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testFilterStateless2(t, input, filters, expected, cmpOpt)
}

// TestFilterStateful2 is [TestFilterStateful] for iter.Seq2 -> iter.Seq2 filters.
func TestFilterStateful2[K, V, K2, V2 any](
	t *testing.T,
	input iter.Seq2[K, V],
	filters []func(iter.Seq2[K, V]) iter.Seq2[K2, V2],
	expected []opiter.KV[K2, V2],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testFilterStateful2(t, input, filters, expected, cmpOpt)
}

// TestFilterStateless2To1 is [TestFilterStateless] for iter.Seq2 -> iter.Seq filters.
func TestFilterStateless2To1[K, V, U any](
	t *testing.T,
	input iter.Seq2[K, V],
	filters []func(iter.Seq2[K, V]) iter.Seq[U],
	expected []U,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testFilterStateless2To1(t, input, filters, expected, cmpOpt)
}

// TestFilterStateful2To1 is [TestFilterStateful] for iter.Seq2 -> iter.Seq filters.
func TestFilterStateful2To1[K, V, U any](
	t *testing.T,
	input iter.Seq2[K, V],
	filters []func(iter.Seq2[K, V]) iter.Seq[U],
	expected []U,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testFilterStateful2To1(t, input, filters, expected, cmpOpt)
}
