package opiter

import "iter"

// Enumerate adds zero-based indices to the values from seq.
func Enumerate[Seq ~func(yield func(V) bool), V any](seq Seq) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := 0
		seq(func(v V) bool {
			ok := yield(i, v)
			i++
			return ok
		})
	}
}

// Pairs combines x and y, stopping when either sequence stops.
func Pairs[X ~func(yield func(K) bool), Y ~func(yield func(V) bool), K, V any](
	x X,
	y Y,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull(iter.Seq[V](y))
		defer stop()
		x(func(k K) bool {
			v, ok := next()
			return ok && yield(k, v)
		})
	}
}

// Pairs2 combines x and y, stopping when either sequence stops.
func Pairs2[
	X ~func(yield func(K1, V1) bool),
	Y ~func(yield func(K2, V2) bool),
	K1, V1, K2, V2 any,
](
	x X,
	y Y,
) iter.Seq2[KV[K1, V1], KV[K2, V2]] {
	return func(yield func(KV[K1, V1], KV[K2, V2]) bool) {
		next, stop := iter.Pull2(iter.Seq2[K2, V2](y))
		defer stop()
		x(func(k1 K1, v1 V1) bool {
			k2, v2, ok := next()
			return ok && yield(PackKV(k1, v1), PackKV(k2, v2))
		})
	}
}

// Transpose swaps the two components of each pair from seq.
func Transpose[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) iter.Seq2[V, K] {
	return func(yield func(V, K) bool) {
		seq(func(k K, v V) bool {
			return yield(v, k)
		})
	}
}

// OmitL drops the latter component of each pair from seq.
func OmitL[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) iter.Seq[K] {
	return func(yield func(K) bool) {
		seq(func(k K, _ V) bool {
			return yield(k)
		})
	}
}

// OmitF drops the former component of each pair from seq.
func OmitF[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(_ K, v V) bool {
			return yield(v)
		})
	}
}

// Omit drops every value while preserving the number of yielded elements.
func Omit[Seq ~func(yield func(V) bool), V any](seq Seq) func(yield func() bool) {
	return func(yield func() bool) {
		seq(func(V) bool {
			return yield()
		})
	}
}

// Omit2 drops every pair while preserving the number of yielded elements.
func Omit2[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) func(yield func() bool) {
	return func(yield func() bool) {
		seq(func(K, V) bool {
			return yield()
		})
	}
}

// Assemble maps each pair from seq to a single value.
func Assemble[Seq ~func(yield func(K, V) bool), K, V, U any](
	fn func(K, V) U,
	seq Seq,
) iter.Seq[U] {
	return func(yield func(U) bool) {
		seq(func(k K, v V) bool {
			return yield(fn(k, v))
		})
	}
}

// Disassemble maps each value from seq to a pair.
func Disassemble[Seq ~func(yield func(U) bool), K, V, U any](
	fn func(U) (K, V),
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(u U) bool {
			return yield(fn(u))
		})
	}
}

func (seq Seq[V]) Enumerate() Seq2[int, V] {
	return Seq2[int, V](Enumerate(seq))
}

func (seq Seq[V]) Pairs[Other ~func(yield func(U) bool), U any](other Other) Seq2[V, U] {
	return Seq2[V, U](Pairs(seq, other))
}

func (seq Seq[V]) Disassemble[K, U any](fn func(V) (K, U)) Seq2[K, U] {
	return Seq2[K, U](Disassemble(fn, seq))
}

func (seq Seq[V]) Omit() func(yield func() bool) {
	return Omit(seq)
}

func (seq Seq2[K, V]) Pairs[Other ~func(yield func(K2, V2) bool), K2, V2 any](
	other Other,
) iter.Seq2[KV[K, V], KV[K2, V2]] {
	return Pairs2(seq, other)
}

func (seq Seq2[K, V]) Transpose() Seq2[V, K] {
	return Seq2[V, K](Transpose(seq))
}

func (seq Seq2[K, V]) OmitF() Seq[V] {
	return Seq[V](OmitF(seq))
}

func (seq Seq2[K, V]) OmitL() Seq[K] {
	return Seq[K](OmitL(seq))
}

func (seq Seq2[K, V]) Assemble[U any](fn func(K, V) U) Seq[U] {
	return Seq[U](Assemble(fn, seq))
}

func (seq Seq2[K, V]) Omit() func(yield func() bool) {
	return Omit2(seq)
}
