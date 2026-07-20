package opiter

import "iter"

var (
	_ Iterable[any]           = FuncIterable[any](nil)
	_ IntoIterable[any]       = FuncIterable[any](nil)
	_ Iterable2[any, any]     = FuncIterable2[any, any](nil)
	_ IntoIterable2[any, any] = FuncIterable2[any, any](nil)
)

type FuncIterable[V any] func() Seq[V]

func WrapFunc[V any, FuncIter ~func() Iter, Iter ~func(yield func(V) bool)](fn FuncIter) FuncIterable[V] {
	return FuncIterable[V](func() Seq[V] { return Seq[V](fn()) })
}

func (f FuncIterable[V]) Iter() iter.Seq[V] {
	return iter.Seq[V](f())
}

func (f FuncIterable[V]) IntoIter() iter.Seq[V] {
	return iter.Seq[V](f())
}

type FuncIterable2[K, V any] func() Seq2[K, V]

func WrapFunc2[K, V any, FuncIter ~func() Iter, Iter ~func(yield func(K, V) bool)](fn FuncIter) FuncIterable2[K, V] {
	return FuncIterable2[K, V](func() Seq2[K, V] { return Seq2[K, V](fn()) })
}

func (f FuncIterable2[K, V]) Iter2() iter.Seq2[K, V] {
	return iter.Seq2[K, V](f())
}

func (f FuncIterable2[K, V]) IntoIter2() iter.Seq2[K, V] {
	return iter.Seq2[K, V](f())
}
