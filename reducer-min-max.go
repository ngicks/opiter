package opiter

import (
	"cmp"

	"github.com/ngicks/option/opt"
)

// Min returns the minimum value of seq, or none if seq is empty.
func Min[Seq ~func(yield func(V) bool), V cmp.Ordered](seq Seq) opt.Option[V] {
	return MinFunc(cmp.Compare, seq)
}

// MinFunc returns the minimum value of seq using fn, or none if seq is empty.
func MinFunc[Seq ~func(yield func(V) bool), V any](
	fn func(x, y V) int,
	seq Seq,
) opt.Option[V] {
	var min opt.Option[V]
	for v := range seq {
		if min.IsNone() || fn(v, min.Value()) < 0 {
			min = opt.Some(v)
		}
	}
	return min
}

// MinFunc returns the minimum value of seq using fn, or none if seq is empty.
func (seq Seq[V]) MinFunc(fn func(x, y V) int) opt.Option[V] {
	return MinFunc(fn, seq)
}

// Max returns the maximum value of seq, or none if seq is empty.
func Max[Seq ~func(yield func(V) bool), V cmp.Ordered](seq Seq) opt.Option[V] {
	return MaxFunc(cmp.Compare, seq)
}

// MaxFunc returns the maximum value of seq using fn, or none if seq is empty.
func MaxFunc[Seq ~func(yield func(V) bool), V any](
	fn func(x, y V) int,
	seq Seq,
) opt.Option[V] {
	var max opt.Option[V]
	for v := range seq {
		if max.IsNone() || fn(v, max.Value()) > 0 {
			max = opt.Some(v)
		}
	}
	return max
}

// MaxFunc returns the maximum value of seq using fn, or none if seq is empty.
func (seq Seq[V]) MaxFunc(fn func(x, y V) int) opt.Option[V] {
	return MaxFunc(fn, seq)
}
