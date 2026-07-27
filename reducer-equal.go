package opiter

import "iter"

// Equal reports whether x and y yield equal values in the same order.
func Equal[X ~func(yield func(V) bool), Y ~func(yield func(V) bool), V comparable](
	x X,
	y Y,
) bool {
	return EqualFunc(func(a, b V) bool { return a == b }, x, y)
}

// Equal2 reports whether x and y yield equal key-value pairs in the same order.
func Equal2[
	X ~func(yield func(K, V) bool),
	Y ~func(yield func(K, V) bool),
	K, V comparable,
](x X, y Y) bool {
	return EqualFunc2(func(k1 K, v1 V, k2 K, v2 V) bool {
		return k1 == k2 && v1 == v2
	}, x, y)
}

// EqualFunc reports whether x and y yield values equal according to fn in the same order.
func EqualFunc[
	X ~func(yield func(V1) bool),
	Y ~func(yield func(V2) bool),
	V1, V2 any,
](fn func(V1, V2) bool, x X, y Y) bool {
	next, stop := iter.Pull(iter.Seq[V2](y))
	defer stop()
	for v1 := range x {
		v2, ok := next()
		if !ok || !fn(v1, v2) {
			return false
		}
	}
	_, ok := next()
	return !ok
}

// EqualFunc reports whether seq and other yield values equal according to fn in the same order.
func (seq Seq[V]) EqualFunc[U any](fn func(V, U) bool, other iter.Seq[U]) bool {
	return EqualFunc(fn, seq, other)
}

// EqualFunc2 reports whether x and y yield pairs equal according to fn in the same order.
func EqualFunc2[
	X ~func(yield func(K1, V1) bool),
	Y ~func(yield func(K2, V2) bool),
	K1, V1, K2, V2 any,
](
	fn func(K1, V1, K2, V2) bool,
	x X,
	y Y,
) bool {
	next, stop := iter.Pull2(iter.Seq2[K2, V2](y))
	defer stop()
	for k1, v1 := range x {
		k2, v2, ok := next()
		if !ok || !fn(k1, v1, k2, v2) {
			return false
		}
	}
	_, _, ok := next()
	return !ok
}

// EqualFunc2 reports whether seq and other yield pairs equal according to fn in the same order.
func (seq Seq2[K, V]) EqualFunc2[K2, V2 any](
	fn func(K, V, K2, V2) bool,
	other iter.Seq2[K2, V2],
) bool {
	return EqualFunc2(fn, seq, other)
}
