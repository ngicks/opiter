package opiter

import "iter"

// Repeat returns an iterator over v repeated n times. A negative n repeats
// forever.
func Repeat[V any](v V, n int) iter.Seq[V] {
	if n < 0 {
		return func(yield func(V) bool) {
			for yield(v) {
			}
		}
	}
	return func(yield func(V) bool) {
		for range n {
			if !yield(v) {
				return
			}
		}
	}
}

// Repeat2 returns an iterator over k and v repeated n times. A negative n
// repeats forever.
func Repeat2[K, V any](k K, v V, n int) iter.Seq2[K, V] {
	if n < 0 {
		return func(yield func(K, V) bool) {
			for yield(k, v) {
			}
		}
	}
	return func(yield func(K, V) bool) {
		for range n {
			if !yield(k, v) {
				return
			}
		}
	}
}

// RepeatFunc returns an iterator that calls fnV and yields its result n times.
// A negative n repeats forever.
func RepeatFunc[V any](fnV func() V, n int) iter.Seq[V] {
	if n < 0 {
		return func(yield func(V) bool) {
			for yield(fnV()) {
			}
		}
	}
	return func(yield func(V) bool) {
		for range n {
			if !yield(fnV()) {
				return
			}
		}
	}
}

// RepeatFunc2 returns an iterator that calls fnK and fnV and yields their
// results n times. A negative n repeats forever.
func RepeatFunc2[K, V any](fnK func() K, fnV func() V, n int) iter.Seq2[K, V] {
	if n < 0 {
		return func(yield func(K, V) bool) {
			for yield(fnK(), fnV()) {
			}
		}
	}
	return func(yield func(K, V) bool) {
		for range n {
			if !yield(fnK(), fnV()) {
				return
			}
		}
	}
}
