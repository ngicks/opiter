package opiter

import "iter"

// Limit returns an iterator over at most n values from seq.
func Limit[Seq ~func(yield func(V) bool), V any](n int, seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		remaining := n
		if remaining <= 0 {
			return
		}
		seq(func(v V) bool {
			if !yield(v) {
				return false
			}
			remaining--
			return remaining > 0
		})
	}
}

// Limit2 returns an iterator over at most n pairs from seq.
func Limit2[Seq ~func(yield func(K, V) bool), K, V any](n int, seq Seq) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		remaining := n
		if remaining <= 0 {
			return
		}
		seq(func(k K, v V) bool {
			if !yield(k, v) {
				return false
			}
			remaining--
			return remaining > 0
		})
	}
}

// LimitUntil yields values from seq until f returns false.
func LimitUntil[Seq ~func(yield func(V) bool), V any](f func(V) bool, seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(v V) bool {
			if !f(v) {
				return false
			}
			return yield(v)
		})
	}
}

// LimitUntil2 yields pairs from seq until f returns false.
func LimitUntil2[Seq ~func(yield func(K, V) bool), K, V any](
	f func(K, V) bool,
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if !f(k, v) {
				return false
			}
			return yield(k, v)
		})
	}
}

// LimitAfter is like LimitUntil but includes the first value for which f returns false.
func LimitAfter[Seq ~func(yield func(V) bool), V any](f func(V) bool, seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(v V) bool {
			if !f(v) {
				yield(v)
				return false
			}
			return yield(v)
		})
	}
}

// LimitAfter2 is like LimitUntil2 but includes the first pair for which f returns false.
func LimitAfter2[Seq ~func(yield func(K, V) bool), K, V any](
	f func(K, V) bool,
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if !f(k, v) {
				yield(k, v)
				return false
			}
			return yield(k, v)
		})
	}
}

func (f Seq[V]) Limit(n int) Seq[V] {
	return Seq[V](Limit(n, f))
}

func (f Seq[V]) LimitUntil(predicate func(V) bool) Seq[V] {
	return Seq[V](LimitUntil(predicate, f))
}

func (f Seq[V]) LimitAfter(predicate func(V) bool) Seq[V] {
	return Seq[V](LimitAfter(predicate, f))
}

func (f Seq2[K, V]) Limit(n int) Seq2[K, V] {
	return Seq2[K, V](Limit2(n, f))
}

func (f Seq2[K, V]) LimitUntil(predicate func(K, V) bool) Seq2[K, V] {
	return Seq2[K, V](LimitUntil2(predicate, f))
}

func (f Seq2[K, V]) LimitAfter(predicate func(K, V) bool) Seq2[K, V] {
	return Seq2[K, V](LimitAfter2(predicate, f))
}
