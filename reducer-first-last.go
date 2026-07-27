package opiter

import "github.com/ngicks/option/opt"

// First returns the first value from seq, or none if seq is empty.
func First[Seq ~func(yield func(V) bool), V any](seq Seq) opt.Option[V] {
	for v := range seq {
		return opt.Some(v)
	}
	return opt.None[V]()
}

// First returns the first value from seq, or none if seq is empty.
func (seq Seq[V]) First() opt.Option[V] {
	return First(seq)
}

// First2 returns the first key-value pair from seq, or none if seq is empty.
func First2[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) opt.Option[KV[K, V]] {
	for k, v := range seq {
		return opt.Some(PackKV(k, v))
	}
	return opt.None[KV[K, V]]()
}

// First2 returns the first key-value pair from seq, or none if seq is empty.
func (seq Seq2[K, V]) First2() opt.Option[KV[K, V]] {
	return First2(seq)
}

// Last returns the last value from seq, or none if seq is empty.
func Last[Seq ~func(yield func(V) bool), V any](seq Seq) opt.Option[V] {
	var last opt.Option[V]
	for v := range seq {
		last = opt.Some(v)
	}
	return last
}

// Last returns the last value from seq, or none if seq is empty.
func (seq Seq[V]) Last() opt.Option[V] {
	return Last(seq)
}

// Last2 returns the last key-value pair from seq, or none if seq is empty.
func Last2[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) opt.Option[KV[K, V]] {
	var last opt.Option[KV[K, V]]
	for k, v := range seq {
		last = opt.Some(PackKV(k, v))
	}
	return last
}

// Last2 returns the last key-value pair from seq, or none if seq is empty.
func (seq Seq2[K, V]) Last2() opt.Option[KV[K, V]] {
	return Last2(seq)
}
