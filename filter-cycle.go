package opiter

import "iter"

// Cycle repeatedly calls a finite, reusable seq until the consumer stops.
func Cycle[Seq ~func(yield func(V) bool), V any](seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			atLeastOne, stopped := false, false
			seq(func(v V) bool {
				atLeastOne = true
				if !yield(v) {
					stopped = true
					return false
				}
				return true
			})
			if stopped || !atLeastOne {
				return
			}
		}
	}
}

// Cycle2 repeatedly calls a finite, reusable seq until the consumer stops.
func Cycle2[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for {
			atLeastOne, stopped := false, false
			seq(func(k K, v V) bool {
				atLeastOne = true
				if !yield(k, v) {
					stopped = true
					return false
				}
				return true
			})
			if stopped || !atLeastOne {
				return
			}
		}
	}
}

// CycleBuffered calls seq once, then repeats its buffered values until the consumer stops.
func CycleBuffered[Seq ~func(yield func(V) bool), V any](seq Seq) iter.Seq[V] {
	return func(yield func(V) bool) {
		var buf []V
		stopped := false
		seq(func(v V) bool {
			if !yield(v) {
				stopped = true
				return false
			}
			buf = append(buf, v)
			return true
		})
		if stopped || len(buf) == 0 {
			return
		}
		for {
			for _, v := range buf {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// CycleBuffered2 calls seq once, then repeats its buffered pairs until the consumer stops.
func CycleBuffered2[Seq ~func(yield func(K, V) bool), K, V any](seq Seq) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var buf []KV[K, V]
		stopped := false
		seq(func(k K, v V) bool {
			if !yield(k, v) {
				stopped = true
				return false
			}
			buf = append(buf, PackKV(k, v))
			return true
		})
		if stopped || len(buf) == 0 {
			return
		}
		for {
			for _, kv := range buf {
				if !yield(kv.K, kv.V) {
					return
				}
			}
		}
	}
}

func (seq Seq[V]) Cycle() Seq[V] {
	return Seq[V](Cycle(seq))
}

func (seq Seq[V]) CycleBuffered() Seq[V] {
	return Seq[V](CycleBuffered(seq))
}

func (seq Seq2[K, V]) Cycle() Seq2[K, V] {
	return Seq2[K, V](Cycle2(seq))
}

func (seq Seq2[K, V]) CycleBuffered() Seq2[K, V] {
	return Seq2[K, V](CycleBuffered2(seq))
}
