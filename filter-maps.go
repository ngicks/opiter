package opiter

import "iter"

// MapsBy returns an iterator over the values in m for each key yielded by
// keys. The keys iterator determines the order.
func MapsBy[M ~map[K]V, Seq ~func(yield func(K) bool), K comparable, V any](
	m M,
	keys Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		keys(func(k K) bool {
			return yield(k, m[k])
		})
	}
}
