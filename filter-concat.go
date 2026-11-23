package opiter

import "iter"

// Concat returns an iterator over the concatenation of seqs.
func Concat[Seq ~func(yield func(V) bool), V any](seqs ...Seq) iter.Seq[V] {
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
func Concat2[Seq ~func(yield func(K, V) bool), K, V any](seqs ...Seq) iter.Seq2[K, V] {
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

// Concat returns an iterator over seq followed by seqs.
func (seq Seq[V]) Concat[Other ~func(yield func(V) bool)](seqs ...Other) Seq[V] {
	all := make([]Seq[V], 0, len(seqs)+1)
	all = append(all, seq)
	for _, other := range seqs {
		all = append(all, Seq[V](other))
	}
	return Seq[V](Concat(all...))
}

// Concat returns an iterator over seq followed by seqs.
func (seq Seq2[K, V]) Concat[Other ~func(yield func(K, V) bool)](seqs ...Other) Seq2[K, V] {
	all := make([]Seq2[K, V], 0, len(seqs)+1)
	all = append(all, seq)
	for _, other := range seqs {
		all = append(all, Seq2[K, V](other))
	}
	return Seq2[K, V](Concat2(all...))
}
