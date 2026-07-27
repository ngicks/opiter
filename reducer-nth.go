package opiter

import "github.com/ngicks/option/opt"

// Nth returns the value at zero-based index n, or none if n is out of range.
func Nth[Seq ~func(yield func(V) bool), V any](n int, seq Seq) opt.Option[V] {
	if n < 0 {
		return opt.None[V]()
	}
	for v := range seq {
		if n == 0 {
			return opt.Some(v)
		}
		n--
	}
	return opt.None[V]()
}

// Nth returns the value at zero-based index n, or none if n is out of range.
func (seq Seq[V]) Nth(n int) opt.Option[V] {
	return Nth(n, seq)
}

// Nth2 returns the key-value pair at zero-based index n, or none if n is out of range.
func Nth2[Seq ~func(yield func(K, V) bool), K, V any](
	n int,
	seq Seq,
) opt.Option[KV[K, V]] {
	if n < 0 {
		return opt.None[KV[K, V]]()
	}
	for k, v := range seq {
		if n == 0 {
			return opt.Some(PackKV(k, v))
		}
		n--
	}
	return opt.None[KV[K, V]]()
}

// Nth2 returns the key-value pair at zero-based index n, or none if n is out of range.
func (seq Seq2[K, V]) Nth2(n int) opt.Option[KV[K, V]] {
	return Nth2(n, seq)
}
