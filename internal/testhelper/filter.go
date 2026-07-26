package testhelper

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
)

func testFilterStatelessOne[U, V any](
	t *testing.T,
	input iter.Seq[U],
	filter func(iter.Seq[U]) iter.Seq[V],
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	testSourceStatelessOne(t, func() iter.Seq[V] { return filter(input) }, expected, cmpOpt)

	for cut := range len(expected) {
		var stopped bool
		upstream := iter.Seq[U](func(yield func(U) bool) {
			input(func(u U) bool {
				if stopped {
					t.Errorf("break at %d: filter consumed upstream after downstream break", cut)
					return false
				}
				return yield(u)
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

func testFilterStatefulOne[U, V any](
	t *testing.T,
	input iter.Seq[U],
	filter func(iter.Seq[U]) iter.Seq[V],
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	values := slices.Collect(input)
	testSourceStatefulOne(t, func() iter.Seq[V] { return filter(OneShotSeq(values...)) }, expected, cmpOpt)
}

func testFilterStateless1To2One[U, K, V any](
	t *testing.T,
	input iter.Seq[U],
	filter func(iter.Seq[U]) iter.Seq2[K, V],
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	testSourceStateless2One(t, func() iter.Seq2[K, V] { return filter(input) }, expected, cmpOpt)

	for cut := range len(expected) {
		var stopped bool
		upstream := iter.Seq[U](func(yield func(U) bool) {
			input(func(u U) bool {
				if stopped {
					t.Errorf("break at %d: filter consumed upstream after downstream break", cut)
					return false
				}
				return yield(u)
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

func testFilterStateful1To2One[U, K, V any](
	t *testing.T,
	input iter.Seq[U],
	filter func(iter.Seq[U]) iter.Seq2[K, V],
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	values := slices.Collect(input)
	testSourceStateful2One(t, func() iter.Seq2[K, V] { return filter(OneShotSeq(values...)) }, expected, cmpOpt)
}

func testFilterStateless[U, V any](
	t *testing.T,
	input iter.Seq[U],
	filters []func(iter.Seq[U]) iter.Seq[V],
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, filter := range filters {
		t.Run(fmt.Sprintf("filters[%d]", i), func(t *testing.T) {
			testFilterStatelessOne(t, input, filter, expected, cmpOpt)
		})
	}
}

func testFilterStateful[U, V any](
	t *testing.T,
	input iter.Seq[U],
	filters []func(iter.Seq[U]) iter.Seq[V],
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, filter := range filters {
		t.Run(fmt.Sprintf("filters[%d]", i), func(t *testing.T) {
			testFilterStatefulOne(t, input, filter, expected, cmpOpt)
		})
	}
}

func testFilterStateless1To2[U, K, V any](
	t *testing.T,
	input iter.Seq[U],
	filters []func(iter.Seq[U]) iter.Seq2[K, V],
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, filter := range filters {
		t.Run(fmt.Sprintf("filters[%d]", i), func(t *testing.T) {
			testFilterStateless1To2One(t, input, filter, expected, cmpOpt)
		})
	}
}

func testFilterStateful1To2[U, K, V any](
	t *testing.T,
	input iter.Seq[U],
	filters []func(iter.Seq[U]) iter.Seq2[K, V],
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, filter := range filters {
		t.Run(fmt.Sprintf("filters[%d]", i), func(t *testing.T) {
			testFilterStateful1To2One(t, input, filter, expected, cmpOpt)
		})
	}
}

// TestFilterStateless asserts every filter in filters, applied over input,
// produces a stateless, well-behaved iterator yielding exactly expected, and
// stops consuming the upstream once the consumer breaks.
// input must be a stateless, well-behaved iterator such as [SampleSeq].
func TestFilterStateless[U, V any](
	t *testing.T,
	input iter.Seq[U],
	filters []func(iter.Seq[U]) iter.Seq[V],
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testFilterStateless(t, input, filters, expected, cmpOpt)
}

// TestFilterStateful asserts every filter in filters transparently passes
// through single-use upstreams: applied over a one-shot iterator carrying
// input's values, the output follows the stateful contract described in the
// package doc. Only streaming filters satisfy this; buffering filters
// (e.g. reverse, sort) should use [TestFilterStateless] only.
// input itself must be stateless; the helper derives one-shot upstreams from it.
func TestFilterStateful[U, V any](
	t *testing.T,
	input iter.Seq[U],
	filters []func(iter.Seq[U]) iter.Seq[V],
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testFilterStateful(t, input, filters, expected, cmpOpt)
}

// TestFilterStateless1To2 is [TestFilterStateless] for iter.Seq -> iter.Seq2 filters.
func TestFilterStateless1To2[U, K, V any](
	t *testing.T,
	input iter.Seq[U],
	filters []func(iter.Seq[U]) iter.Seq2[K, V],
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testFilterStateless1To2(t, input, filters, expected, cmpOpt)
}

// TestFilterStateful1To2 is [TestFilterStateful] for iter.Seq -> iter.Seq2 filters.
func TestFilterStateful1To2[U, K, V any](
	t *testing.T,
	input iter.Seq[U],
	filters []func(iter.Seq[U]) iter.Seq2[K, V],
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testFilterStateful1To2(t, input, filters, expected, cmpOpt)
}
