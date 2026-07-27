package opiter

import "iter"

// RunningReduce yields every intermediate reducer result.
func RunningReduce[Seq ~func(yield func(V) bool), V, Sum any](
	reducer func(accumulator Sum, current V, i int) Sum,
	initial Sum,
	seq Seq,
) iter.Seq[Sum] {
	return func(yield func(Sum) bool) {
		sum := initial
		i := 0
		seq(func(v V) bool {
			sum = reducer(sum, v, i)
			i++
			return yield(sum)
		})
	}
}

func (f Seq[V]) RunningReduce[Sum any](reducer func(Sum, V, int) Sum, initial Sum) Seq[Sum] {
	return Seq[Sum](RunningReduce(reducer, initial, f))
}
