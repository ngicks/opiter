package opiter

// Reduce combines the values in seq using fn, starting with initial.
func Reduce[Seq ~func(yield func(V) bool), V, Sum any](
	fn func(Sum, V) Sum,
	initial Sum,
	seq Seq,
) Sum {
	sum := initial
	for v := range seq {
		sum = fn(sum, v)
	}
	return sum
}

// Reduce combines the values in seq using fn, starting with initial.
func (seq Seq[V]) Reduce[Sum any](fn func(Sum, V) Sum, initial Sum) Sum {
	return Reduce(fn, initial, seq)
}

// Reduce2 combines the key-value pairs in seq using fn, starting with initial.
func Reduce2[Seq ~func(yield func(K, V) bool), K, V, Sum any](
	fn func(Sum, K, V) Sum,
	initial Sum,
	seq Seq,
) Sum {
	sum := initial
	for k, v := range seq {
		sum = fn(sum, k, v)
	}
	return sum
}

// Reduce2 combines the key-value pairs in seq using fn, starting with initial.
func (seq Seq2[K, V]) Reduce2[Sum any](fn func(Sum, K, V) Sum, initial Sum) Sum {
	return Reduce2(fn, initial, seq)
}
