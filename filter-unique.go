package opiter

import "iter"

// UniqueFunc yields the first value for each identifier returned by f.
func UniqueFunc[Seq ~func(yield func(V) bool), V any, Id comparable](
	f func(V) Id,
	seq Seq,
) iter.Seq[V] {
	return func(yield func(V) bool) {
		seen := make(map[Id]struct{})
		seq(func(v V) bool {
			id := f(v)
			if _, ok := seen[id]; ok {
				return true
			}
			seen[id] = struct{}{}
			return yield(v)
		})
	}
}

// Unique yields the first occurrence of every value from seq.
func Unique[Seq ~func(yield func(V) bool), V comparable](seq Seq) iter.Seq[V] {
	return UniqueFunc(func(v V) V { return v }, seq)
}

// UniqueFunc2 yields the first pair for each identifier returned by f.
func UniqueFunc2[Seq ~func(yield func(K, V) bool), K, V any, Id comparable](
	f func(K, V) Id,
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seen := make(map[Id]struct{})
		seq(func(k K, v V) bool {
			id := f(k, v)
			if _, ok := seen[id]; ok {
				return true
			}
			seen[id] = struct{}{}
			return yield(k, v)
		})
	}
}

// Unique2 yields the first occurrence of every pair from seq.
func Unique2[Seq ~func(yield func(K, V) bool), K, V comparable](seq Seq) iter.Seq2[K, V] {
	return UniqueFunc2(func(k K, v V) KV[K, V] { return PackKV(k, v) }, seq)
}

func (seq Seq[V]) UniqueFunc[Id comparable](identifier func(V) Id) Seq[V] {
	return Seq[V](UniqueFunc(identifier, seq))
}

func (seq Seq2[K, V]) UniqueFunc[Id comparable](identifier func(K, V) Id) Seq2[K, V] {
	return Seq2[K, V](UniqueFunc2(identifier, seq))
}
