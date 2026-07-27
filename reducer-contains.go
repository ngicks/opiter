package opiter

// Contains reports whether value is present in seq.
func Contains[Seq ~func(yield func(V) bool), V comparable](value V, seq Seq) bool {
	return Find(value, seq).IsSome()
}

// ContainsFunc reports whether at least one value in seq satisfies fn.
func ContainsFunc[Seq ~func(yield func(V) bool), V any](fn func(V) bool, seq Seq) bool {
	return FindFunc(fn, seq).IsSome()
}

// ContainsFunc reports whether at least one value in seq satisfies fn.
func (seq Seq[V]) ContainsFunc(fn func(V) bool) bool {
	return ContainsFunc(fn, seq)
}

// Contains2 reports whether the key-value pair is present in seq.
func Contains2[Seq ~func(yield func(K, V) bool), K, V comparable](
	key K,
	value V,
	seq Seq,
) bool {
	return Find2(key, value, seq).IsSome()
}

// ContainsFunc2 reports whether at least one key-value pair in seq satisfies fn.
func ContainsFunc2[Seq ~func(yield func(K, V) bool), K, V any](
	fn func(K, V) bool,
	seq Seq,
) bool {
	return FindFunc2(fn, seq).IsSome()
}

// ContainsFunc2 reports whether at least one key-value pair in seq satisfies fn.
func (seq Seq2[K, V]) ContainsFunc2(fn func(K, V) bool) bool {
	return ContainsFunc2(fn, seq)
}
