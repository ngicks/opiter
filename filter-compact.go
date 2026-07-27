package opiter

import "iter"

// Compact skips consecutive equal values from seq.
func Compact[Seq ~func(yield func(V) bool), V comparable](seq Seq) iter.Seq[V] {
	return CompactFunc(func(x, y V) bool { return x == y }, seq)
}

// CompactFunc skips consecutive values that eq reports equal.
func CompactFunc[Seq ~func(yield func(V) bool), V any](
	eq func(V, V) bool,
	seq Seq,
) iter.Seq[V] {
	return func(yield func(V) bool) {
		var prev V
		first := true
		seq(func(v V) bool {
			different := first || !eq(prev, v)
			first, prev = false, v
			return !different || yield(v)
		})
	}
}

// Compact2 skips consecutive equal pairs from seq.
func Compact2[Seq ~func(yield func(K, V) bool), K, V comparable](seq Seq) iter.Seq2[K, V] {
	return CompactFunc2(func(k1 K, v1 V, k2 K, v2 V) bool {
		return k1 == k2 && v1 == v2
	}, seq)
}

// CompactFunc2 skips consecutive pairs that eq reports equal.
func CompactFunc2[Seq ~func(yield func(K, V) bool), K, V any](
	eq func(K, V, K, V) bool,
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var prevK K
		var prevV V
		first := true
		seq(func(k K, v V) bool {
			different := first || !eq(prevK, prevV, k, v)
			first, prevK, prevV = false, k, v
			return !different || yield(k, v)
		})
	}
}

func (f Seq[V]) CompactFunc(eq func(V, V) bool) Seq[V] {
	return Seq[V](CompactFunc(eq, f))
}

func (f Seq2[K, V]) CompactFunc(eq func(K, V, K, V) bool) Seq2[K, V] {
	return Seq2[K, V](CompactFunc2(eq, f))
}
