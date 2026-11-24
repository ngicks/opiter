package opiter

import (
	"maps"
	"slices"
)

// SlicesAll wraps [slices.All].
// It returns an iterator over index-value pairs in s.
func SlicesAll[S ~[]E, E any](s S) Seq2[int, E] {
	return Seq2[int, E](slices.All(s))
}

// SlicesBackward wraps [slices.Backward].
// It returns an iterator over index-value pairs in s, from the last index to the first.
func SlicesBackward[S ~[]E, E any](s S) Seq2[int, E] {
	return Seq2[int, E](slices.Backward(s))
}

// SlicesValues wraps [slices.Values].
// It returns an iterator over the elements of s.
func SlicesValues[S ~[]E, E any](s S) Seq[E] {
	return Seq[E](slices.Values(s))
}

// SlicesChunk wraps [slices.Chunk].
// It returns an iterator over consecutive subslices of up to n elements of s.
//
// It panics if n is less than 1, as [slices.Chunk] does.
func SlicesChunk[S ~[]E, E any](s S, n int) Seq[S] {
	return Seq[S](slices.Chunk(s, n))
}

// MapsAll wraps [maps.All].
// It returns an iterator over key-value pairs of m in unspecified order.
func MapsAll[M ~map[K]V, K comparable, V any](m M) Seq2[K, V] {
	return Seq2[K, V](maps.All(m))
}

// MapsKeys wraps [maps.Keys].
// It returns an iterator over keys of m in unspecified order.
func MapsKeys[M ~map[K]V, K comparable, V any](m M) Seq[K] {
	return Seq[K](maps.Keys(m))
}

// MapsValues wraps [maps.Values].
// It returns an iterator over values of m in unspecified order.
func MapsValues[M ~map[K]V, K comparable, V any](m M) Seq[V] {
	return Seq[V](maps.Values(m))
}
