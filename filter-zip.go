package opiter

import (
	"iter"

	"github.com/ngicks/option/opt"
	"github.com/ngicks/option/tuple"
)

// Zip combines two sequences, retaining unmatched values as optional sides.
func Zip[
	X ~func(yield func(L) bool),
	Y ~func(yield func(R) bool),
	L, R any,
](x X, y Y) iter.Seq[tuple.Tuple2[opt.Option[L], opt.Option[R]]] {
	return func(yield func(tuple.Tuple2[opt.Option[L], opt.Option[R]]) bool) {
		next, stop := iter.Pull(iter.Seq[R](y))
		defer stop()
		right, rightOK := next()
		stopped := false
		x(func(left L) bool {
			if !yield(tuple.Tuple2[opt.Option[L], opt.Option[R]]{
				V1: opt.Some(left),
				V2: opt.FromOk(right, rightOK),
			}) {
				stopped = true
				return false
			}
			right, rightOK = next()
			return true
		})
		if stopped {
			return
		}
		for rightOK {
			if !yield(tuple.Tuple2[opt.Option[L], opt.Option[R]]{
				V1: opt.None[L](),
				V2: opt.Some(right),
			}) {
				return
			}
			right, rightOK = next()
		}
	}
}

// Zip2 combines two pair sequences, retaining unmatched pairs as optional sides.
func Zip2[
	X ~func(yield func(K1, V1) bool),
	Y ~func(yield func(K2, V2) bool),
	K1, V1, K2, V2 any,
](
	x X,
	y Y,
) iter.Seq[tuple.Tuple2[opt.Option[KV[K1, V1]], opt.Option[KV[K2, V2]]]] {
	return func(yield func(tuple.Tuple2[opt.Option[KV[K1, V1]], opt.Option[KV[K2, V2]]]) bool) {
		next, stop := iter.Pull2(iter.Seq2[K2, V2](y))
		defer stop()
		k2, v2, rightOK := next()
		stopped := false
		x(func(k1 K1, v1 V1) bool {
			if !yield(tuple.Tuple2[opt.Option[KV[K1, V1]], opt.Option[KV[K2, V2]]]{
				V1: opt.Some(PackKV(k1, v1)),
				V2: opt.FromOk(PackKV(k2, v2), rightOK),
			}) {
				stopped = true
				return false
			}
			k2, v2, rightOK = next()
			return true
		})
		if stopped {
			return
		}
		for rightOK {
			if !yield(tuple.Tuple2[opt.Option[KV[K1, V1]], opt.Option[KV[K2, V2]]]{
				V1: opt.None[KV[K1, V1]](),
				V2: opt.Some(PackKV(k2, v2)),
			}) {
				return
			}
			k2, v2, rightOK = next()
		}
	}
}

func (seq Seq[V]) Zip[
	Other ~func(yield func(U) bool),
	Z tuple.Tuple2[opt.Option[V], opt.Option[U]],
	U any,
](
	other Other,
) Seq[Z] {
	return func(yield func(Z) bool) {
		Zip(seq.Iter(), other)(func(z tuple.Tuple2[opt.Option[V], opt.Option[U]]) bool {
			return yield(Z(z))
		})
	}
}

func (seq Seq2[K, V]) Zip[
	Other ~func(yield func(K2, V2) bool),
	Z tuple.Tuple2[opt.Option[KV[K, V]], opt.Option[KV[K2, V2]]],
	K2, V2 any,
](
	other Other,
) Seq[Z] {
	return func(yield func(Z) bool) {
		Zip2(seq.Iter2(), other)(func(
			z tuple.Tuple2[opt.Option[KV[K, V]], opt.Option[KV[K2, V2]]],
		) bool {
			return yield(Z(z))
		})
	}
}
