package opiter_test

import (
	"iter"
	"slices"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"github.com/ngicks/option/opt"
	"github.com/ngicks/option/tuple"
)

type namedSeq[V any] func(yield func(V) bool)

type namedSeq2[K, V any] func(yield func(K, V) bool)

func TestZip(t *testing.T) {
	cmpOpt := []goCmp.Option{goCmp.Comparer(func(
		x, y tuple.Tuple2[opt.Option[int], opt.Option[string]],
	) bool {
		return x.V1.EqualFunc(y.V1, func(a, b int) bool { return a == b }) &&
			x.V2.EqualFunc(y.V2, func(a, b string) bool { return a == b })
	})}

	leftOnlyExpected := []tuple.Tuple2[opt.Option[int], opt.Option[string]]{
		{V1: opt.Some(1), V2: opt.Some("a")},
		{V1: opt.Some(2), V2: opt.None[string]()},
	}
	leftOnlyFilters := []func(iter.Seq[int]) iter.Seq[tuple.Tuple2[opt.Option[int], opt.Option[string]]]{
		func(s iter.Seq[int]) iter.Seq[tuple.Tuple2[opt.Option[int], opt.Option[string]]] {
			return opiter.Zip(s, namedSeq[string](slices.Values([]string{"a"})))
		},
		func(s iter.Seq[int]) iter.Seq[tuple.Tuple2[opt.Option[int], opt.Option[string]]] {
			return opiter.WrapSeq(s).Zip(namedSeq[string](slices.Values([]string{"a"}))).Iter()
		},
	}
	testhelper.TestFilterStateless(
		t,
		slices.Values([]int{1, 2}),
		leftOnlyFilters,
		leftOnlyExpected,
		cmpOpt,
	)
	left, right := leftOnlyExpected[0].Unpack()
	if left.IsNone() || right.IsNone() {
		t.Fatal("first pair should be present")
	}
	if left.Value() != 1 || right.Value() != "a" {
		t.Fatalf("Unpack() = %v, %v", left, right)
	}

	rightOnlyExpected := []tuple.Tuple2[opt.Option[int], opt.Option[string]]{
		{V1: opt.Some(1), V2: opt.Some("a")},
		{V1: opt.None[int](), V2: opt.Some("b")},
	}
	rightOnlyFilters := []func(iter.Seq[int]) iter.Seq[tuple.Tuple2[opt.Option[int], opt.Option[string]]]{
		func(s iter.Seq[int]) iter.Seq[tuple.Tuple2[opt.Option[int], opt.Option[string]]] {
			return opiter.Zip(s, namedSeq[string](slices.Values([]string{"a", "b"})))
		},
		func(s iter.Seq[int]) iter.Seq[tuple.Tuple2[opt.Option[int], opt.Option[string]]] {
			return opiter.WrapSeq(s).Zip(namedSeq[string](slices.Values([]string{"a", "b"}))).Iter()
		},
	}
	testhelper.TestFilterStateless(
		t,
		slices.Values([]int{1}),
		rightOnlyFilters,
		rightOnlyExpected,
		cmpOpt,
	)

	namedLeft := namedSeq[int](slices.Values([]int{1, 2}))
	namedRight := namedSeq[int](slices.Values([]int{3, 4}))
	namedExpected := []tuple.Tuple2[opt.Option[int], opt.Option[int]]{
		{V1: opt.Some(1), V2: opt.Some(3)},
		{V1: opt.Some(2), V2: opt.Some(4)},
	}
	namedCmp := goCmp.Comparer(func(
		x, y tuple.Tuple2[opt.Option[int], opt.Option[int]],
	) bool {
		return x.V1.EqualFunc(y.V1, func(a, b int) bool { return a == b }) &&
			x.V2.EqualFunc(y.V2, func(a, b int) bool { return a == b })
	})
	if diff := goCmp.Diff(namedExpected, slices.Collect(opiter.Zip(namedLeft, namedRight)), namedCmp); diff != "" {
		t.Fatalf("Zip(namedSeq, namedSeq) mismatch (-want +got):\n%s", diff)
	}
	if diff := goCmp.Diff(namedExpected, slices.Collect(opiter.WrapSeq(namedLeft).Zip(namedRight).Iter()), namedCmp); diff != "" {
		t.Fatalf("Seq.Zip(namedSeq) mismatch (-want +got):\n%s", diff)
	}
}

func TestZip2(t *testing.T) {
	type zipped = tuple.Tuple2[
		opt.Option[opiter.KV[int, string]],
		opt.Option[opiter.KV[string, int]],
	]
	cmpOpt := []goCmp.Option{goCmp.Comparer(func(x, y zipped) bool {
		return x.V1.EqualFunc(y.V1, func(a, b opiter.KV[int, string]) bool { return a == b }) &&
			x.V2.EqualFunc(y.V2, func(a, b opiter.KV[string, int]) bool { return a == b })
	})}

	leftOnlyExpected := []zipped{
		{
			V1: opt.Some(opiter.PackKV(1, "a")),
			V2: opt.Some(opiter.PackKV("x", 3)),
		},
		{
			V1: opt.Some(opiter.PackKV(2, "b")),
			V2: opt.None[opiter.KV[string, int]](),
		},
	}
	leftOnlyFilters := []func(iter.Seq2[int, string]) iter.Seq[zipped]{
		func(s iter.Seq2[int, string]) iter.Seq[zipped] {
			return opiter.Zip2(s, namedSeq2[string, int](opiter.Values2([]opiter.KV[string, int]{{"x", 3}})))
		},
		func(s iter.Seq2[int, string]) iter.Seq[zipped] {
			return opiter.WrapSeq2(s).Zip(
				namedSeq2[string, int](opiter.Values2([]opiter.KV[string, int]{{"x", 3}})),
			).Iter()
		},
	}
	testhelper.TestFilterStateless2To1(
		t,
		opiter.Values2([]opiter.KV[int, string]{{1, "a"}, {2, "b"}}),
		leftOnlyFilters,
		leftOnlyExpected,
		cmpOpt,
	)
	leftOpt, rightOpt := leftOnlyExpected[0].Unpack()
	leftKV, rightKV := leftOpt.Value(), rightOpt.Value()
	if leftKV.K != 1 || leftKV.V != "a" || rightKV.K != "x" || rightKV.V != 3 {
		t.Fatalf("Zip2 tuple = %v, %v", leftKV, rightKV)
	}

	rightOnlyExpected := []zipped{
		{
			V1: opt.Some(opiter.PackKV(1, "a")),
			V2: opt.Some(opiter.PackKV("x", 3)),
		},
		{
			V1: opt.None[opiter.KV[int, string]](),
			V2: opt.Some(opiter.PackKV("y", 4)),
		},
	}
	rightOnlyFilters := []func(iter.Seq2[int, string]) iter.Seq[zipped]{
		func(s iter.Seq2[int, string]) iter.Seq[zipped] {
			return opiter.Zip2(s, namedSeq2[string, int](opiter.Values2(
				[]opiter.KV[string, int]{{"x", 3}, {"y", 4}},
			)))
		},
		func(s iter.Seq2[int, string]) iter.Seq[zipped] {
			return opiter.WrapSeq2(s).Zip(namedSeq2[string, int](opiter.Values2(
				[]opiter.KV[string, int]{{"x", 3}, {"y", 4}},
			))).Iter()
		},
	}
	testhelper.TestFilterStateless2To1(
		t,
		opiter.Values2([]opiter.KV[int, string]{{1, "a"}}),
		rightOnlyFilters,
		rightOnlyExpected,
		cmpOpt,
	)

	namedLeft := namedSeq2[int, string](opiter.Values2([]opiter.KV[int, string]{{1, "a"}}))
	namedRight := namedSeq2[int, string](opiter.Values2([]opiter.KV[int, string]{{2, "b"}}))
	namedExpected := []tuple.Tuple2[
		opt.Option[opiter.KV[int, string]],
		opt.Option[opiter.KV[int, string]],
	]{
		{
			V1: opt.Some(opiter.PackKV(1, "a")),
			V2: opt.Some(opiter.PackKV(2, "b")),
		},
	}
	namedCmp := goCmp.Comparer(func(
		x, y tuple.Tuple2[
			opt.Option[opiter.KV[int, string]],
			opt.Option[opiter.KV[int, string]],
		],
	) bool {
		return x.V1.EqualFunc(y.V1, func(a, b opiter.KV[int, string]) bool { return a == b }) &&
			x.V2.EqualFunc(y.V2, func(a, b opiter.KV[int, string]) bool { return a == b })
	})
	if diff := goCmp.Diff(namedExpected, slices.Collect(opiter.Zip2(namedLeft, namedRight)), namedCmp); diff != "" {
		t.Fatalf("Zip2(namedSeq2, namedSeq2) mismatch (-want +got):\n%s", diff)
	}
	if diff := goCmp.Diff(namedExpected, slices.Collect(opiter.WrapSeq2(namedLeft).Zip(namedRight).Iter()), namedCmp); diff != "" {
		t.Fatalf("Seq2.Zip(namedSeq2) mismatch (-want +got):\n%s", diff)
	}
}

func TestZipMethodsReturnSeq(t *testing.T) {
	type zipped = tuple.Tuple2[opt.Option[int], opt.Option[string]]
	var seq opiter.Seq[zipped] = opiter.WrapSeq(slices.Values([]int{1})).
		Zip(namedSeq[string](slices.Values([]string{"a"})))
	if got := slices.Collect(seq.Limit(1).Iter()); len(got) != 1 {
		t.Fatalf("Seq.Zip(...).Limit(1) yielded %d values", len(got))
	}
	if got := slices.Collect(opiter.WrapSeq(slices.Values([]int{1})).
		Zip(namedSeq[string](slices.Values([]string{"a"}))).
		Limit(1).
		Iter()); len(got) != 1 {
		t.Fatalf("direct Seq.Zip(...).Limit(1) yielded %d values", len(got))
	}

	type zipped2 = tuple.Tuple2[
		opt.Option[opiter.KV[int, string]],
		opt.Option[opiter.KV[string, int]],
	]
	var seq2 opiter.Seq[zipped2] = opiter.WrapSeq2(
		opiter.Values2([]opiter.KV[int, string]{{1, "a"}}),
	).Zip(namedSeq2[string, int](
		opiter.Values2([]opiter.KV[string, int]{{"x", 2}}),
	))
	if got := slices.Collect(seq2.Limit(1).Iter()); len(got) != 1 {
		t.Fatalf("Seq2.Zip(...).Limit(1) yielded %d values", len(got))
	}
	if got := slices.Collect(opiter.WrapSeq2(
		opiter.Values2([]opiter.KV[int, string]{{1, "a"}}),
	).Zip(namedSeq2[string, int](
		opiter.Values2([]opiter.KV[string, int]{{"x", 2}}),
	)).Limit(1).Iter()); len(got) != 1 {
		t.Fatalf("direct Seq2.Zip(...).Limit(1) yielded %d values", len(got))
	}
}
