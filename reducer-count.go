package opiter

// Count returns the number of values in seq.
func Count[Seq ~func(yield func(V) bool), V any](seq Seq) int {
	n := 0
	for range seq {
		n++
	}
	return n
}

// Count returns the number of values in seq.
func (seq Seq[V]) Count() int {
	return Count(seq)
}

// Count2 returns the number of key-value pairs in seq.
func Count2[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) int {
	n := 0
	for range seq {
		n++
	}
	return n
}

// Count returns the number of key-value pairs in seq.
func (seq Seq2[K, V]) Count() int {
	return Count2(seq)
}
