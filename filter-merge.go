package opiter

import (
	"cmp"
	"iter"
)

// Merge merges two sequences of ordered values.
func Merge[V cmp.Ordered](x, y iter.Seq[V]) iter.Seq[V] {
	return MergeFunc(cmp.Compare[V], x, y)
}

// MergeFunc merges two sequences ordered by f.
func MergeFunc[V any](f func(V, V) int, x, y iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		next, stop := iter.Pull(y)
		defer stop()
		v2, ok2 := next()
		stopped := false
		x(func(v1 V) bool {
			for ok2 && f(v1, v2) > 0 {
				if !yield(v2) {
					stopped = true
					return false
				}
				v2, ok2 = next()
			}
			if !yield(v1) {
				stopped = true
				return false
			}
			return true
		})
		if stopped {
			return
		}
		for ok2 {
			if !yield(v2) {
				return
			}
			v2, ok2 = next()
		}
	}
}

// Merge2 merges two sequences of pairs ordered by key.
func Merge2[K cmp.Ordered, V any](x, y iter.Seq2[K, V]) iter.Seq2[K, V] {
	return MergeFunc2(cmp.Compare[K], x, y)
}

// MergeFunc2 merges two sequences of pairs ordered by f applied to their keys.
func MergeFunc2[K, V any](
	f func(K, K) int,
	x, y iter.Seq2[K, V],
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull2(y)
		defer stop()
		k2, v2, ok2 := next()
		stopped := false
		x(func(k1 K, v1 V) bool {
			for ok2 && f(k1, k2) > 0 {
				if !yield(k2, v2) {
					stopped = true
					return false
				}
				k2, v2, ok2 = next()
			}
			if !yield(k1, v1) {
				stopped = true
				return false
			}
			return true
		})
		if stopped {
			return
		}
		for ok2 {
			if !yield(k2, v2) {
				return
			}
			k2, v2, ok2 = next()
		}
	}
}

func (f Seq[V]) MergeFunc(order func(V, V) int, other iter.Seq[V]) Seq[V] {
	return Seq[V](MergeFunc(order, f.Iter(), other))
}

func (f Seq2[K, V]) MergeFunc(order func(K, K) int, other iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](MergeFunc2(order, f.Iter2(), other))
}
