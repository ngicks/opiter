package opiter

// ForEach calls fn with every value in seq.
func ForEach[Seq ~func(yield func(V) bool), V any](fn func(V), seq Seq) {
	for v := range seq {
		fn(v)
	}
}

// ForEach calls fn with every value in seq.
func (seq Seq[V]) ForEach(fn func(V)) {
	ForEach(fn, seq)
}

// ForEach2 calls fn with every key-value pair in seq.
func ForEach2[Seq ~func(yield func(K, V) bool), K, V any](fn func(K, V), seq Seq) {
	for k, v := range seq {
		fn(k, v)
	}
}

// ForEach2 calls fn with every key-value pair in seq.
func (seq Seq2[K, V]) ForEach2(fn func(K, V)) {
	ForEach2(fn, seq)
}

// Discard fully consumes seq.
func Discard[Seq ~func(yield func(V) bool), V any](seq Seq) {
	for range seq {
	}
}

// Discard fully consumes seq.
func (seq Seq[V]) Discard() {
	Discard(seq)
}

// Discard2 fully consumes seq.
func Discard2[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) {
	for range seq {
	}
}

// Discard2 fully consumes seq.
func (seq Seq2[K, V]) Discard2() {
	Discard2(seq)
}
