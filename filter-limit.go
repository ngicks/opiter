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

// TakeWhile yields values from seq while f returns true.
// If inclusive is true, it also yields the first value for which f returns false.
func TakeWhile[Seq ~func(yield func(V) bool), V any](
	inclusive bool,
	f func(V) bool,
	seq Seq,
) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(v V) bool {
			if !f(v) {
				if inclusive {
					yield(v)
				}
				return false
			}
			return yield(v)
		})
	}
}

// TakeWhile2 yields pairs from seq while f returns true.
// If inclusive is true, it also yields the first pair for which f returns false.
func TakeWhile2[Seq ~func(yield func(K, V) bool), K, V any](
	inclusive bool,
	f func(K, V) bool,
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if !f(k, v) {
				if inclusive {
					yield(k, v)
				}
				return false
			}
			return yield(k, v)
		})
	}
}

func (seq Seq[V]) Limit(n int) Seq[V] {
	return Seq[V](Limit(n, seq))
}

func (seq Seq[V]) TakeWhile(inclusive bool, predicate func(V) bool) Seq[V] {
	return Seq[V](TakeWhile(inclusive, predicate, seq))
}

func (seq Seq2[K, V]) Limit(n int) Seq2[K, V] {
	return Seq2[K, V](Limit2(n, seq))
}

func (seq Seq2[K, V]) TakeWhile(inclusive bool, predicate func(K, V) bool) Seq2[K, V] {
	return Seq2[K, V](TakeWhile2(inclusive, predicate, seq))
}
