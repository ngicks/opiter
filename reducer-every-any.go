package opiter

// Every reports whether every value in seq satisfies fn.
func Every[Seq ~func(yield func(V) bool), V any](fn func(V) bool, seq Seq) bool {
	for v := range seq {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Every reports whether every value in seq satisfies fn.
func (seq Seq[V]) Every(fn func(V) bool) bool {
	return Every(fn, seq)
}

// Every2 reports whether every key-value pair in seq satisfies fn.
func Every2[Seq ~func(yield func(K, V) bool), K, V any](
	fn func(K, V) bool,
	seq Seq,
) bool {
	for k, v := range seq {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Every2 reports whether every key-value pair in seq satisfies fn.
func (seq Seq2[K, V]) Every2(fn func(K, V) bool) bool {
	return Every2(fn, seq)
}

// Any reports whether any value in seq satisfies fn.
func Any[Seq ~func(yield func(V) bool), V any](fn func(V) bool, seq Seq) bool {
	return FindFunc(fn, seq).IsSome()
}

// Any reports whether any value in seq satisfies fn.
func (seq Seq[V]) Any(fn func(V) bool) bool {
	return Any(fn, seq)
}

// Any2 reports whether any key-value pair in seq satisfies fn.
func Any2[Seq ~func(yield func(K, V) bool), K, V any](
	fn func(K, V) bool,
	seq Seq,
) bool {
	return FindFunc2(fn, seq).IsSome()
}

// Any2 reports whether any key-value pair in seq satisfies fn.
func (seq Seq2[K, V]) Any2(fn func(K, V) bool) bool {
	return Any2(fn, seq)
}
