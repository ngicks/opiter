package opiter

import "iter"

// Filter returns an iterator over values from seq for which f returns true.
func Filter[Seq ~func(yield func(V) bool), V any](f func(V) bool, seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(v V) bool {
			return !f(v) || yield(v)
		})
	}
}

// Filter2 returns an iterator over pairs from seq for which f returns true.
func Filter2[Seq ~func(yield func(K, V) bool), K, V any](
	f func(K, V) bool,
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			return !f(k, v) || yield(k, v)
		})
	}
}

// Filter returns values from f for which predicate returns true.
func (f Seq[V]) Filter(predicate func(V) bool) Seq[V] {
	return Seq[V](Filter(predicate, f))
}

// Filter returns pairs from f for which predicate returns true.
func (f Seq2[K, V]) Filter(predicate func(K, V) bool) Seq2[K, V] {
	return Seq2[K, V](Filter2(predicate, f))
}
