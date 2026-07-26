package testhelper

import (
	"iter"
	"slices"

	"github.com/ngicks/opiter"
)

// Sample inputs for iterator tests: unsorted, with duplicate values
// (SamplePairs also duplicates a key).
var (
	SampleValues = []int{3, 1, 4, 1, 5, 9, 2, 6}
	SamplePairs  = []opiter.KV[string, int]{
		{K: "foo", V: 3},
		{K: "bar", V: 1},
		{K: "baz", V: 4},
		{K: "foo", V: 1},
		{K: "qux", V: 5},
	}
)

// SampleSeq returns a stateless iterator over SampleValues.
func SampleSeq() iter.Seq[int] {
	return slices.Values(SampleValues)
}

// SampleSeq2 returns a stateless iterator over SamplePairs.
func SampleSeq2() iter.Seq2[string, int] {
	return opiter.Values2(SamplePairs)
}

// OneShotSeq returns a single-use iterator over values,
// following the stateful contract described in the package doc.
func OneShotSeq[V any](values ...V) iter.Seq[V] {
	i := 0
	return func(yield func(V) bool) {
		// i advances before yield so that a rejected value counts as consumed.
		for i < len(values) {
			v := values[i]
			i++
			if !yield(v) {
				return
			}
		}
	}
}

// OneShotSeq2 returns a single-use iterator over pairs,
// following the stateful contract described in the package doc.
func OneShotSeq2[K, V any](pairs ...opiter.KV[K, V]) iter.Seq2[K, V] {
	i := 0
	return func(yield func(K, V) bool) {
		for i < len(pairs) {
			kv := pairs[i]
			i++
			if !yield(kv.K, kv.V) {
				return
			}
		}
	}
}
