package opiter

import (
	"iter"
)

var (
	_ Iterable2[any, any]     = Seq2[any, any](nil)
	_ IntoIterable2[any, any] = Seq2[any, any](nil)
)

// Seq wraps iter.Seq[V] to add generic function on it.
type Seq[V any] iter.Seq[V]

// WrapSeq wraps seq into Seq[V].
//
// This is here only for type inference; less typing, easy by auto-fill.
func WrapSeq[V any, Iter ~func(yield func(V) bool)](seq Iter) Seq[V] {
	return Seq[V](seq)
}

func (f Seq[V]) Iter() iter.Seq[V] {
	return iter.Seq[V](f)
}

func (f Seq[V]) IntoIter() iter.Seq[V] {
	return iter.Seq[V](f)
}

// Seq2 wraps iter.Seq2[K, V] to add generic function on it.
type Seq2[K, V any] iter.Seq2[K, V]

// WrapSeq2 wraps seq into Seq2[V].
//
// This is here only for type inference; less typing, easy by auto-fill.
func WrapSeq2[K, V any, Iter ~func(yield func(K, V) bool)](seq Iter) Seq2[K, V] {
	return Seq2[K, V](seq)
}

func (f Seq2[K, V]) Iter2() iter.Seq2[K, V] {
	return iter.Seq2[K, V](f)
}

func (f Seq2[K, V]) IntoIter2() iter.Seq2[K, V] {
	return iter.Seq2[K, V](f)
}
