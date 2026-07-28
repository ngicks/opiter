package opiter

import "iter"

// Tap calls tap for every value and yields seq unchanged.
func Tap[Seq ~func(yield func(V) bool), V any](tap func(V), seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(v V) bool {
			tap(v)
			return yield(v)
		})
	}
}

// Tap2 calls tap for every pair and yields seq unchanged.
func Tap2[Seq ~func(yield func(K, V) bool), K, V any](
	tap func(K, V),
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			tap(k, v)
			return yield(k, v)
		})
	}
}

// TapLast calls tap after seq is consumed without an early break.
func TapLast[Seq ~func(yield func(V) bool), V any](tap func(), seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		completed := true
		seq(func(v V) bool {
			completed = yield(v)
			return completed
		})
		if completed {
			tap()
		}
	}
}

// TapLast2 calls tap after seq is consumed without an early break.
func TapLast2[Seq ~func(yield func(K, V) bool), K, V any](
	tap func(),
	seq Seq,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		completed := true
		seq(func(k K, v V) bool {
			completed = yield(k, v)
			return completed
		})
		if completed {
			tap()
		}
	}
}

func (seq Seq[V]) Tap(tap func(V)) Seq[V] {
	return Seq[V](Tap(tap, seq))
}

func (seq Seq[V]) TapLast(tap func()) Seq[V] {
	return Seq[V](TapLast(tap, seq))
}

func (seq Seq2[K, V]) Tap(tap func(K, V)) Seq2[K, V] {
	return Seq2[K, V](Tap2(tap, seq))
}

func (seq Seq2[K, V]) TapLast(tap func()) Seq2[K, V] {
	return Seq2[K, V](TapLast2(tap, seq))
}
