package opiter

import (
	"iter"

	"github.com/ngicks/option/tuple"
)

// KV is pair of K and V.
// It is useful if paris of k and v needs to be sent through
// somewhere it only accepts a single vlaue, e.g. channel.
//
// There are some utilities that help you consume KV easier:
//
//   - [PackKV] : creates a KV which saves you some key strokes by inferring types.
//   - [Values2] : converts []KV[K, V] based types to iter.Seq2.
//   - [AppendSeq2], [Collect2] : collects iter.Seq2 to []KV or types based on it.
//   - [AssembleKV], [DisassembleKV] : converts iter.Seq2[K, V] to/from iter.Seq[KV[K, V]].
type KV[K, V any] struct {
	K K
	V V
}

func PackKV[K, V any](k K, v V) KV[K, V] {
	return KV[K, V]{K: k, V: v}
}

func (kv KV[K, V]) Unpack() (K, V) {
	return kv.K, kv.V
}

func (kv KV[K, V]) ToTuple2() tuple.Tuple2[K, V] {
	return tuple.Tuple2[K, V]{V1: kv.K, V2: kv.V}
}

// Values2 returns iterator that yields the KV slice elements in order.
func Values2[S ~[]KV[K, V], K, V any](s S) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, kv := range s {
			if !yield(kv.K, kv.V) {
				return
			}
		}
	}
}

// AppendSeq2 appends the values from seq to the KV slice and
// returns the extended slice.
func AppendSeq2[S ~[]KV[K, V], Seq ~func(yield func(K, V) bool), K, V any](s S, seq Seq) S {
	for k, v := range seq {
		s = append(s, KV[K, V]{k, v})
	}
	return s
}

// Collect2 collects values from seq into a new KV slice and returns it.
func Collect2[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) []KV[K, V] {
	return AppendSeq2[[]KV[K, V]](nil, seq)
}

// AssembleKV converts iter.Seq2[K, V] to iter.Seq[KV[K, V]].
func AssembleKV[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) iter.Seq[KV[K, V]] {
	return func(yield func(KV[K, V]) bool) {
		seq(func(k K, v V) bool {
			return yield(PackKV(k, v))
		})
	}
}

// DisassembleKV converts iter.Seq2[K, V] from iter.Seq[KV[K, V]].
func DisassembleKV[Seq ~func(yield func(KV[K, V]) bool), K, V any](seq Seq) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(kv KV[K, V]) bool {
			return yield(kv.K, kv.V)
		})
	}
}

var _ Iterable2[any, any] = KVPairs[any, any](nil)

// KVPairs adds the Iter2 method to slice of KV.
type KVPairs[K, V any] []KV[K, V]

func (v KVPairs[K, V]) Iter2() iter.Seq2[K, V] {
	return Values2(v)
}
