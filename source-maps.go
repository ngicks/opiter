package opiter

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

// MapsKeys returns an iterator over the values in m for each key yielded by
// keys. The keys iterator determines the order.
func MapsKeys[M ~map[K]V, Seq ~func(yield func(K) bool), K comparable, V any](
	m M,
	keys Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k := range keys {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}

// MapsSorted returns an iterator over m in ascending key order. It takes a
// snapshot of the keys each time the iterator is invoked.
func MapsSorted[M ~map[K]V, K cmp.Ordered, V any](m M) iter.Seq2[K, V] {
	return MapsSortedFunc(m, cmp.Compare)
}

// MapsSortedFunc is like [MapsSorted], using f to order keys.
func MapsSortedFunc[M ~map[K]V, K comparable, V any](
	m M,
	f func(i, j K) int,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		keys := slices.SortedFunc(maps.Keys(m), f)
		for _, k := range keys {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}

// MapsOverlay returns the unique key-value pairs from mm. For keys present in
// multiple maps, the value from the rightmost map is yielded.
func MapsOverlay[M ~map[K]V, K comparable, V any](mm ...M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seen := make(map[K]struct{})
		for _, m := range slices.Backward(mm) {
			for k, v := range m {
				if _, ok := seen[k]; ok {
					continue
				}
				seen[k] = struct{}{}
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
