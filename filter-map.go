package opiter

import "iter"

// Map returns an iterator over f applied to seq.
func Map[Seq ~func(yield func(V1) bool), V1, V2 any](f func(V1) V2, seq Seq) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		seq(func(v V1) bool {
			return yield(f(v))
		})
	}
}

// Map2 returns an iterator over f applied to seq.
func Map2[Seq ~func(yield func(K1, V1) bool), K1, V1, K2, V2 any](
	f func(K1, V1) (K2, V2),
	seq Seq,
) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		seq(func(k K1, v V1) bool {
			return yield(f(k, v))
		})
	}
}

// Map returns an iterator over fn applied to f.
func (f Seq[V]) Map[U any](fn func(V) U) Seq[U] {
	return Seq[U](Map(fn, f))
}

// Map returns an iterator over fn applied to f.
func (f Seq2[K, V]) Map[K2, V2 any](fn func(K, V) (K2, V2)) Seq2[K2, V2] {
	return Seq2[K2, V2](Map2(fn, f))
}
