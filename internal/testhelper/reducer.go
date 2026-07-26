package testhelper

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
)

func testReducerOne[V, R any](
	t *testing.T,
	input iter.Seq[V],
	reduce func(iter.Seq[V]) R,
	expected R,
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	got := reduce(input)
	if d := goCmp.Diff(expected, got, cmpOpt...); d != "" {
		t.Errorf("(-want +got):\n%s", d)
	}

	// Doing same test again: "no use after break" rule.
	invocations := 0
	oneShot := OneShotSeq(slices.Collect(input)...)
	got = reduce(func(yield func(V) bool) {
		invocations++
		if invocations > 1 {
			t.Errorf("input invoked %d times; reducers must work on single-use iterators", invocations)
		}
		oneShot(yield)
	})
	if d := goCmp.Diff(expected, got, cmpOpt...); d != "" {
		t.Errorf("one-shot input: (-want +got):\n%s", d)
	}
}

func testReducer2One[K, V, R any](
	t *testing.T,
	input iter.Seq2[K, V],
	reduce func(iter.Seq2[K, V]) R,
	expected R,
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	got := reduce(input)
	if d := goCmp.Diff(expected, got, cmpOpt...); d != "" {
		t.Errorf("(-want +got):\n%s", d)
	}

	// Doing same test again: "no use after break" rule.
	invocations := 0
	oneShot := OneShotSeq2(opiter.Collect2(input)...)
	got = reduce(func(yield func(K, V) bool) {
		invocations++
		if invocations > 1 {
			t.Errorf("input invoked %d times; reducers must work on single-use iterators", invocations)
		}
		oneShot(yield)
	})
	if d := goCmp.Diff(expected, got, cmpOpt...); d != "" {
		t.Errorf("one-shot input: (-want +got):\n%s", d)
	}
}

func testReducer[V, R any](
	t *testing.T,
	input iter.Seq[V],
	reducers []func(iter.Seq[V]) R,
	expected R,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, reduce := range reducers {
		t.Run(fmt.Sprintf("reducers[%d]", i), func(t *testing.T) {
			testReducerOne(t, input, reduce, expected, cmpOpt)
		})
	}
}

func testReducer2[K, V, R any](
	t *testing.T,
	input iter.Seq2[K, V],
	reducers []func(iter.Seq2[K, V]) R,
	expected R,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, reduce := range reducers {
		t.Run(fmt.Sprintf("reducers[%d]", i), func(t *testing.T) {
			testReducer2One(t, input, reduce, expected, cmpOpt)
		})
	}
}

// TestReducer asserts every reducer in reducers collapses input to expected,
// both on input as-is and on a single-use iterator carrying the same values,
// invoking the latter at most once.
// input must be a stateless, well-behaved iterator such as [SampleSeq].
func TestReducer[V, R any](
	t *testing.T,
	input iter.Seq[V],
	reducers []func(iter.Seq[V]) R,
	expected R,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testReducer(t, input, reducers, expected, cmpOpt)
}

// TestReducer2 is [TestReducer] for iter.Seq2-shaped inputs.
func TestReducer2[K, V, R any](
	t *testing.T,
	input iter.Seq2[K, V],
	reducers []func(iter.Seq2[K, V]) R,
	expected R,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testReducer2(t, input, reducers, expected, cmpOpt)
}
