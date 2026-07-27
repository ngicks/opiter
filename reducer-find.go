package opiter

import "github.com/ngicks/option/opt"

// Find returns the first occurrence of value and its zero-based index.
func Find[Seq ~func(yield func(V) bool), V comparable](
	value V,
	seq Seq,
) opt.Option[KV[int, V]] {
	return FindFunc(func(v V) bool { return v == value }, seq)
}

// FindFunc returns the first value satisfying fn and its zero-based index.
func FindFunc[Seq ~func(yield func(V) bool), V any](
	fn func(V) bool,
	seq Seq,
) opt.Option[KV[int, V]] {
	i := 0
	for v := range seq {
		if fn(v) {
			return opt.Some(PackKV(i, v))
		}
		i++
	}
	return opt.None[KV[int, V]]()
}

// FindFunc returns the first value satisfying fn and its zero-based index.
func (seq Seq[V]) FindFunc(fn func(V) bool) opt.Option[KV[int, V]] {
	return FindFunc(fn, seq)
}

// Find2 returns the first occurrence of key and value and its zero-based index.
func Find2[Seq ~func(yield func(K, V) bool), K, V comparable](
	key K,
	value V,
	seq Seq,
) opt.Option[KV[int, KV[K, V]]] {
	return FindFunc2(func(k K, v V) bool { return k == key && v == value }, seq)
}

// FindFunc2 returns the first key-value pair satisfying fn and its zero-based index.
func FindFunc2[Seq ~func(yield func(K, V) bool), K, V any](
	fn func(K, V) bool,
	seq Seq,
) opt.Option[KV[int, KV[K, V]]] {
	i := 0
	for k, v := range seq {
		if fn(k, v) {
			return opt.Some(PackKV(i, PackKV(k, v)))
		}
		i++
	}
	return opt.None[KV[int, KV[K, V]]]()
}

// FindFunc2 returns the first key-value pair satisfying fn and its zero-based index.
func (seq Seq2[K, V]) FindFunc2(fn func(K, V) bool) opt.Option[KV[int, KV[K, V]]] {
	return FindFunc2(fn, seq)
}

// FindLast returns the final occurrence of value and its zero-based index.
func FindLast[Seq ~func(yield func(V) bool), V comparable](
	value V,
	seq Seq,
) opt.Option[KV[int, V]] {
	return FindLastFunc(func(v V) bool { return v == value }, seq)
}

// FindLastFunc returns the final value satisfying fn and its zero-based index.
func FindLastFunc[Seq ~func(yield func(V) bool), V any](
	fn func(V) bool,
	seq Seq,
) opt.Option[KV[int, V]] {
	var found opt.Option[KV[int, V]]
	i := 0
	for v := range seq {
		if fn(v) {
			found = opt.Some(PackKV(i, v))
		}
		i++
	}
	return found
}

// FindLastFunc returns the final value satisfying fn and its zero-based index.
func (seq Seq[V]) FindLastFunc(fn func(V) bool) opt.Option[KV[int, V]] {
	return FindLastFunc(fn, seq)
}

// FindLast2 returns the final occurrence of key and value and its zero-based index.
func FindLast2[Seq ~func(yield func(K, V) bool), K, V comparable](
	key K,
	value V,
	seq Seq,
) opt.Option[KV[int, KV[K, V]]] {
	return FindLastFunc2(func(k K, v V) bool { return k == key && v == value }, seq)
}

// FindLastFunc2 returns the final key-value pair satisfying fn and its zero-based index.
func FindLastFunc2[Seq ~func(yield func(K, V) bool), K, V any](
	fn func(K, V) bool,
	seq Seq,
) opt.Option[KV[int, KV[K, V]]] {
	var found opt.Option[KV[int, KV[K, V]]]
	i := 0
	for k, v := range seq {
		if fn(k, v) {
			found = opt.Some(PackKV(i, PackKV(k, v)))
		}
		i++
	}
	return found
}

// FindLastFunc2 returns the final key-value pair satisfying fn and its zero-based index.
func (seq Seq2[K, V]) FindLastFunc2(fn func(K, V) bool) opt.Option[KV[int, KV[K, V]]] {
	return FindLastFunc2(fn, seq)
}
