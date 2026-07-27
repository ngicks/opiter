package testhelper

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
)

func testSourceStatelessOne[V any](
	t *testing.T,
	source func() iter.Seq[V],
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	seq := source()
	for _, phase := range []string{"first pass", "second pass"} {
		got := slices.Collect(seq)
		assertValues(t, expected, got, cmpOpt, phase)
	}

	got := slices.Collect(source())
	assertValues(t, expected, got, cmpOpt, "fresh instance")

	for cut := range len(expected) {
		seq := source()
		part, rej, rejected := takeN(seq, cut)
		if !rejected {
			t.Errorf("break at %d: iterator ended before yielding element %d", cut, cut)
			continue
		}
		assertValues(t, expected[:cut+1], append(part, rej), cmpOpt, "break at %d", cut)
		got := slices.Collect(seq)
		assertValues(t, expected, got, cmpOpt, "re-invocation after break at %d", cut)
	}
}

func testSourceStatefulOne[V any](
	t *testing.T,
	source func() iter.Seq[V],
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	seq := source()
	got := slices.Collect(seq)
	assertValues(t, expected, got, cmpOpt, "full pass")
	if again := slices.Collect(seq); len(again) != 0 {
		t.Errorf("exhausted iterator yielded %d more value(s)", len(again))
	}

	for cut := range len(expected) {
		seq := source()
		part, rej, rejected := takeN(seq, cut)
		if !rejected {
			t.Errorf("break at %d: iterator ended before yielding element %d", cut, cut)
			continue
		}
		assertValues(t, expected[:cut+1], append(part, rej), cmpOpt, "break at %d", cut)
		rest := slices.Collect(seq)
		assertValues(t, expected[cut+1:], rest, cmpOpt, "resume after break at %d", cut)
	}
}

func testSourceStateless2One[K, V any](
	t *testing.T,
	source func() iter.Seq2[K, V],
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	seq := source()
	for _, phase := range []string{"first pass", "second pass"} {
		got := opiter.Collect2(seq)
		assertValues(t, expected, got, cmpOpt, phase)
	}

	got := opiter.Collect2(source())
	assertValues(t, expected, got, cmpOpt, "fresh instance")

	for cut := range len(expected) {
		seq := source()
		part, rej, rejected := takeN2(seq, cut)
		if !rejected {
			t.Errorf("break at %d: iterator ended before yielding element %d", cut, cut)
			continue
		}
		assertValues(t, expected[:cut+1], append(part, rej), cmpOpt, "break at %d", cut)
		got := opiter.Collect2(seq)
		assertValues(t, expected, got, cmpOpt, "re-invocation after break at %d", cut)
	}
}

func testSourceStateful2One[K, V any](
	t *testing.T,
	source func() iter.Seq2[K, V],
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()

	seq := source()
	got := opiter.Collect2(seq)
	assertValues(t, expected, got, cmpOpt, "full pass")
	if again := opiter.Collect2(seq); len(again) != 0 {
		t.Errorf("exhausted iterator yielded %d more pair(s)", len(again))
	}

	for cut := range len(expected) {
		seq := source()
		part, rej, rejected := takeN2(seq, cut)
		if !rejected {
			t.Errorf("break at %d: iterator ended before yielding element %d", cut, cut)
			continue
		}
		assertValues(t, expected[:cut+1], append(part, rej), cmpOpt, "break at %d", cut)
		rest := opiter.Collect2(seq)
		assertValues(t, expected[cut+1:], rest, cmpOpt, "resume after break at %d", cut)
	}
}

func testSourceStateless[V any, Iter ~func(yield func(V) bool)](
	t *testing.T,
	sources []func() Iter,
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, mk := range sources {
		t.Run(fmt.Sprintf("sources[%d]", i), func(t *testing.T) {
			testSourceStatelessOne(t, func() iter.Seq[V] { return iter.Seq[V](mk()) }, expected, cmpOpt)
		})
	}
}

func testSourceStateful[V any, Iter ~func(yield func(V) bool)](
	t *testing.T,
	sources []func() Iter,
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, mk := range sources {
		t.Run(fmt.Sprintf("sources[%d]", i), func(t *testing.T) {
			testSourceStatefulOne(t, func() iter.Seq[V] { return iter.Seq[V](mk()) }, expected, cmpOpt)
		})
	}
}

func testSourceStateless2[K, V any, Iter ~func(yield func(K, V) bool)](
	t *testing.T,
	sources []func() Iter,
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, mk := range sources {
		t.Run(fmt.Sprintf("sources[%d]", i), func(t *testing.T) {
			testSourceStateless2One(t, func() iter.Seq2[K, V] { return iter.Seq2[K, V](mk()) }, expected, cmpOpt)
		})
	}
}

func testSourceStateful2[K, V any, Iter ~func(yield func(K, V) bool)](
	t *testing.T,
	sources []func() Iter,
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	for i, mk := range sources {
		t.Run(fmt.Sprintf("sources[%d]", i), func(t *testing.T) {
			testSourceStateful2One(t, func() iter.Seq2[K, V] { return iter.Seq2[K, V](mk()) }, expected, cmpOpt)
		})
	}
}

// TestSourceStateless asserts every factory in sources produces a stateless,
// well-behaved iterator yielding exactly expected. See the package doc for the
// contracts. Factories may return the same iterator value on every call.
func TestSourceStateless[V any, Iter ~func(yield func(V) bool)](
	t *testing.T,
	sources []func() Iter,
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testSourceStateless(t, sources, expected, cmpOpt)
}

// TestSourceStateful asserts every factory in sources produces a stateful,
// well-behaved iterator yielding exactly expected. See the package doc for the
// contracts. Each factory call must return a fresh, independent iterator.
func TestSourceStateful[V any, Iter ~func(yield func(V) bool)](
	t *testing.T,
	sources []func() Iter,
	expected []V,
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testSourceStateful(t, sources, expected, cmpOpt)
}

// TestSourceStateless2 is [TestSourceStateless] for iter.Seq2-shaped iterators.
func TestSourceStateless2[K, V any, Iter ~func(yield func(K, V) bool)](
	t *testing.T,
	sources []func() Iter,
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testSourceStateless2(t, sources, expected, cmpOpt)
}

// TestSourceStateful2 is [TestSourceStateful] for iter.Seq2-shaped iterators.
func TestSourceStateful2[K, V any, Iter ~func(yield func(K, V) bool)](
	t *testing.T,
	sources []func() Iter,
	expected []opiter.KV[K, V],
	cmpOpt []goCmp.Option,
) {
	t.Helper()
	testSourceStateful2(t, sources, expected, cmpOpt)
}
