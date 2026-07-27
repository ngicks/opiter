package opiter

import "iter"

// Once returns an iterator that yields v once.
func Once[V any](v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		yield(v)
	}
}

// Once2 returns an iterator that yields k and v once.
func Once2[K, V any](k K, v V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		yield(k, v)
	}
}
