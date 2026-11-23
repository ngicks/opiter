package opiter

import (
	"iter"
	"slices"
)

// Free function counterparts are not defined here:
// calling collector(seq) directly does the same as Collect,
// the standard library has slices.Collect and basic-kv.go has Collect2 for slices.

// Collect passes seq to collector and returns its result.
//
// It lets a method chain end with any collecting function, e.g. slices.Collect or slices.Sorted.
func (seq Seq[V]) Collect[S ~func(yield func(V) bool), U any](collector func(S) U) U {
	return collector(S(seq))
}

// CollectSlice collects values from seq into a new slice and returns it.
func (seq Seq[V]) CollectSlice() []V {
	return seq.Collect(slices.Collect)
}

// AppendSeq appends the values from seq to s and returns the extended slice.
func (seq Seq[V]) AppendSeq[S ~[]V](s S) S {
	return slices.AppendSeq(s, iter.Seq[V](seq))
}

// Collect passes seq to collector and returns its result.
//
// It lets a method chain end with any collecting function, e.g. maps.Collect or [Collect2].
func (seq Seq2[K, V]) Collect[S ~func(yield func(K, V) bool), U any](collector func(S) U) U {
	return collector(S(seq))
}

// CollectSlice collects key-value pairs from seq into a new KV slice and returns it.
func (seq Seq2[K, V]) CollectSlice() []KV[K, V] {
	return seq.Collect(Collect2[iter.Seq2[K, V]])
}

// AppendSeq appends the key-value pairs from seq to the KV slice and returns the extended slice.
func (seq Seq2[K, V]) AppendSeq[S ~[]KV[K, V]](s S) S {
	return AppendSeq2(s, seq)
}
