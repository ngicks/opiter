package opiter

import "iter"

// Concat returns an iterator over the concatenation of seqs.
func Concat[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			stopped := false
			seq(func(v V) bool {
				if !yield(v) {
					stopped = true
					return false
				}
				return true
			})
			if stopped {
				return
			}
		}
	}
}

// Concat2 returns an iterator over the concatenation of seqs.
func Concat2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			stopped := false
			seq(func(k K, v V) bool {
				if !yield(k, v) {
					stopped = true
					return false
				}
				return true
			})
			if stopped {
				return
			}
		}
	}
}

// Concat returns an iterator over f followed by seqs.
func (f Seq[V]) Concat(seqs ...iter.Seq[V]) Seq[V] {
	return Seq[V](Concat(append([]iter.Seq[V]{f.Iter()}, seqs...)...))
}

// Concat returns an iterator over f followed by seqs.
func (f Seq2[K, V]) Concat(seqs ...iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](Concat2(append([]iter.Seq2[K, V]{f.Iter2()}, seqs...)...))
}
