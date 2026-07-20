package opiter

import "iter"

// Iterable wraps basic Iter method.
//
// Iter should always return pure / stateless iterators, which always generates same set of data.
type Iterable[V any] interface {
	Iter() iter.Seq[V]
}

// Iterable2 wraps basic Iter2 method.
//
// Iter2 should always return pure / stateless iterators, which always generates same set of pairs.
type Iterable2[K, V any] interface {
	Iter2() iter.Seq2[K, V]
}

// IntoIterable wraps basic IntoIter2 method.
//
// IntoIter might return non-pure / stateful iterators, which would also mutate internal state of
// implementation. Therefore calling the method or invoking the returned iterator multiple times
// might yield different data without replaying them.
type IntoIterable[V any] interface {
	IntoIter() iter.Seq[V]
}

// IntoIterable2 wraps basic IntoIter2 method.
//
// IntoIter2 might return non-pure / stateful iterators, which would also mutate internal state of
// implementation. Therefore calling the method or invoking the returned iterator multiple times
// might yield different data without replaying them.
type IntoIterable2[K, V any] interface {
	IntoIter2() iter.Seq2[K, V]
}
