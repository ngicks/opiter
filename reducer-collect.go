package opiter

import (
	"iter"
	"slices"
)

// Free function counterparts are not defined here:
// the standard library has slices.Collect and basic-kv.go has Collect2.

// Collect collects values from seq into a new slice and returns it.
func (seq Seq[V]) Collect() []V {
	return slices.Collect(iter.Seq[V](seq))
}

// Collect2 collects key-value pairs from seq into a new KV slice and returns it.
func (seq Seq2[K, V]) Collect2() []KV[K, V] {
	return Collect2(seq)
}
