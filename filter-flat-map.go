package opiter

import "iter"

// FlatMap returns an iterator over values from the iterators fn returns for each value from seq.
func FlatMap[
	Seq ~func(yield func(V) bool),
	Inner ~func(yield func(U) bool),
	V, U any,
](
	fn func(V) Inner,
	seq Seq,
) iter.Seq[U] {
	return func(yield func(U) bool) {
		seq(func(v V) bool {
			keepGoing := true
			fn(v)(func(u U) bool {
				keepGoing = yield(u)
				return keepGoing
			})
			return keepGoing
		})
	}
}

// FlatMap2 returns an iterator over pairs from the iterators fn returns for each pair from seq.
func FlatMap2[
	Seq ~func(yield func(K, V) bool),
	Inner ~func(yield func(K2, V2) bool),
	K, V, K2, V2 any,
](
	fn func(K, V) Inner,
	seq Seq,
) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		seq(func(k K, v V) bool {
			keepGoing := true
			fn(k, v)(func(k2 K2, v2 V2) bool {
				keepGoing = yield(k2, v2)
				return keepGoing
			})
			return keepGoing
		})
	}
}

// FlatMap returns an iterator over values from the iterators fn returns for each value from seq.
func (seq Seq[V]) FlatMap[Inner ~func(yield func(U) bool), U any](fn func(V) Inner) Seq[U] {
	return Seq[U](FlatMap(fn, seq))
}

// FlatMap returns an iterator over pairs from the iterators fn returns for each pair from seq.
func (seq Seq2[K, V]) FlatMap[Inner ~func(yield func(K2, V2) bool), K2, V2 any](
	fn func(K, V) Inner,
) Seq2[K2, V2] {
	return Seq2[K2, V2](FlatMap2(fn, seq))
}
