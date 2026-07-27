package opiter

import "iter"

// Skip returns an iterator over seq after skipping n values.
func Skip[Seq ~func(yield func(V) bool), V any](n int, seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		remaining := n
		seq(func(v V) bool {
			if remaining > 0 {
				remaining--
				return true
			}
			return yield(v)
		})
	}
}

// Skip2 returns an iterator over seq after skipping n pairs.
func Skip2[Seq ~func(yield func(K, V) bool), K, V any](n int, seq Seq) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		remaining := n
		seq(func(k K, v V) bool {
			if remaining > 0 {
				remaining--
				return true
			}
			return yield(k, v)
		})
	}
}

// SkipLast returns an iterator over seq without its last n values.
func SkipLast[Seq ~func(yield func(V) bool), V any](n int, seq Seq) iter.Seq[V] {
	if n <= 0 {
		return iter.Seq[V](seq)
	}
	return func(yield func(V) bool) {
		buf := make([]V, n)
		cursor, full := 0, false
		seq(func(v V) bool {
			if !full {
				buf[cursor] = v
				cursor++
				if cursor == n {
					cursor, full = 0, true
				}
				return true
			}
			old := buf[cursor]
			if !yield(old) {
				return false
			}
			buf[cursor] = v
			cursor = (cursor + 1) % n
			return true
		})
	}
}

// SkipLast2 returns an iterator over seq without its last n pairs.
func SkipLast2[Seq ~func(yield func(K, V) bool), K, V any](n int, seq Seq) iter.Seq2[K, V] {
	if n <= 0 {
		return iter.Seq2[K, V](seq)
	}
	return func(yield func(K, V) bool) {
		buf := make([]KV[K, V], n)
		cursor, full := 0, false
		seq(func(k K, v V) bool {
			if !full {
				buf[cursor] = PackKV(k, v)
				cursor++
				if cursor == n {
					cursor, full = 0, true
				}
				return true
			}
			old := buf[cursor]
			if !yield(old.K, old.V) {
				return false
			}
			buf[cursor] = PackKV(k, v)
			cursor = (cursor + 1) % n
			return true
		})
	}
}

// SkipWhile skips values until f first returns false.
func SkipWhile[Seq ~func(yield func(V) bool), V any](f func(V) bool, seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		skipping := true
		seq(func(v V) bool {
			if skipping && f(v) {
				return true
			}
			skipping = false
			return yield(v)
		})
	}
}

// SkipWhile2 skips pairs until f first returns false.
func SkipWhile2[Seq ~func(yield func(K, V) bool), K, V any](
	f func(K, V) bool,
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		skipping := true
		seq(func(k K, v V) bool {
			if skipping && f(k, v) {
				return true
			}
			skipping = false
			return yield(k, v)
		})
	}
}

func (f Seq[V]) Skip(n int) Seq[V] {
	return Seq[V](Skip(n, f))
}

func (f Seq[V]) SkipLast(n int) Seq[V] {
	return Seq[V](SkipLast(n, f))
}

func (f Seq[V]) SkipWhile(predicate func(V) bool) Seq[V] {
	return Seq[V](SkipWhile(predicate, f))
}

func (f Seq2[K, V]) Skip(n int) Seq2[K, V] {
	return Seq2[K, V](Skip2(n, f))
}

func (f Seq2[K, V]) SkipLast(n int) Seq2[K, V] {
	return Seq2[K, V](SkipLast2(n, f))
}

func (f Seq2[K, V]) SkipWhile(predicate func(K, V) bool) Seq2[K, V] {
	return Seq2[K, V](SkipWhile2(predicate, f))
}
