package opiter

import "iter"

// Empty returns an iterator over no values.
func Empty[V any]() iter.Seq[V] {
	return func(yield func(V) bool) {}
}

// Empty2 returns an iterator over no key-value pairs.
func Empty2[K, V any]() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {}
}
