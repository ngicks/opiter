package opiter

import "iter"

// Flatten returns an iterator over each element of the slices from seq.
func Flatten[Seq ~func(yield func(S) bool), S ~[]E, E any](seq Seq) iter.Seq[E] {
	return func(yield func(E) bool) {
		seq(func(s S) bool {
			for _, v := range s {
				if !yield(v) {
					return false
				}
			}
			return true
		})
	}
}

// FlattenSeq returns an iterator over values from the iterators in seq.
func FlattenSeq[Seq ~func(yield func(Inner) bool), Inner ~func(yield func(V) bool), V any](
	seq Seq,
) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(inner Inner) bool {
			keepGoing := true
			inner(func(v V) bool {
				keepGoing = yield(v)
				return keepGoing
			})
			return keepGoing
		})
	}
}

// FlattenSeq2 returns an iterator over pairs from the iterators in seq.
func FlattenSeq2[
	Seq ~func(yield func(Inner) bool),
	Inner ~func(yield func(K, V) bool),
	K, V any,
](
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(inner Inner) bool {
			keepGoing := true
			inner(func(k K, v V) bool {
				keepGoing = yield(k, v)
				return keepGoing
			})
			return keepGoing
		})
	}
}

// FlattenF flattens the first slice component while repeating the second component.
func FlattenF[Seq ~func(yield func(S, V) bool), S ~[]E, E, V any](seq Seq) iter.Seq2[E, V] {
	return func(yield func(E, V) bool) {
		seq(func(s S, v V) bool {
			for _, e := range s {
				if !yield(e, v) {
					return false
				}
			}
			return true
		})
	}
}

// FlattenL flattens the second slice component while repeating the first component.
func FlattenL[Seq ~func(yield func(K, S) bool), S ~[]E, K, E any](seq Seq) iter.Seq2[K, E] {
	return func(yield func(K, E) bool) {
		seq(func(k K, s S) bool {
			for _, e := range s {
				if !yield(k, e) {
					return false
				}
			}
			return true
		})
	}
}

// FlattenSeqF flattens the first iterator component while repeating the second component.
func FlattenSeqF[
	Seq ~func(yield func(Inner, V) bool),
	Inner ~func(yield func(K) bool),
	K, V any,
](
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(inner Inner, v V) bool {
			keepGoing := true
			inner(func(k K) bool {
				keepGoing = yield(k, v)
				return keepGoing
			})
			return keepGoing
		})
	}
}

// FlattenSeqL flattens the second iterator component while repeating the first component.
func FlattenSeqL[
	Seq ~func(yield func(K, Inner) bool),
	Inner ~func(yield func(V) bool),
	K, V any,
](
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, inner Inner) bool {
			keepGoing := true
			inner(func(v V) bool {
				keepGoing = yield(k, v)
				return keepGoing
			})
			return keepGoing
		})
	}
}
