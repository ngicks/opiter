package opiter

import (
	"cmp"
	"iter"
)

// Merge merges two sequences of ordered values.
func Merge[X ~func(yield func(V) bool), Y ~func(yield func(V) bool), V cmp.Ordered](
	x X,
	y Y,
) iter.Seq[V] {
	return MergeFunc(cmp.Compare[V], x, y)
}

// MergeFunc merges two sequences ordered by f.
func MergeFunc[X ~func(yield func(V) bool), Y ~func(yield func(V) bool), V any](
	f func(V, V) int,
	x X,
	y Y,
) iter.Seq[V] {
	return func(yield func(V) bool) {
		next, stop := iter.Pull(iter.Seq[V](y))
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
func Merge2[X ~func(yield func(K, V) bool), Y ~func(yield func(K, V) bool), K cmp.Ordered, V any](
	x X,
	y Y,
) iter.Seq2[K, V] {
	return MergeFunc2(cmp.Compare[K], x, y)
}

// MergeFunc2 merges two sequences of pairs ordered by f applied to their keys.
func MergeFunc2[X ~func(yield func(K, V) bool), Y ~func(yield func(K, V) bool), K, V any](
	f func(K, K) int,
	x X,
	y Y,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull2(iter.Seq2[K, V](y))
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

// MergeFunc merges seq and other ordered by order.
func (seq Seq[V]) MergeFunc[Other ~func(yield func(V) bool)](
	order func(V, V) int,
	other Other,
) Seq[V] {
	return Seq[V](MergeFunc(order, seq, other))
}

// MergeFunc merges seq and other ordered by order applied to their keys.
func (seq Seq2[K, V]) MergeFunc[Other ~func(yield func(K, V) bool)](
	order func(K, K) int,
	other Other,
) Seq2[K, V] {
	return Seq2[K, V](MergeFunc2(order, seq, other))
}
