package opiter

import "iter"

// Window yields overlapping subslices of size n.
func Window[S ~[]E, E any](s S, n int) iter.Seq[S] {
	return func(yield func(S) bool) {
		if n <= 0 {
			return
		}
		for start, end := 0, n; end <= len(s); start, end = start+1, end+1 {
			if !yield(s[start:end:end]) {
				return
			}
		}
	}
}

// WindowSeq yields moving windows of n values from seq.
func WindowSeq[Seq ~func(yield func(V) bool), V any](n int, seq Seq) iter.Seq[iter.Seq[V]] {
	return func(yield func(iter.Seq[V]) bool) {
		if n <= 0 {
			return
		}
		buf := make([]V, n)
		cursor, full := 0, false
		seq(func(v V) bool {
			if !full {
				buf[cursor] = v
				cursor++
				if cursor < n {
					return true
				}
				cursor, full = 0, true
				return yield(sliceRing(buf, cursor))
			}
			buf[cursor] = v
			cursor = (cursor + 1) % n
			return yield(sliceRing(buf, cursor))
		})
	}
}

func sliceRing[S ~[]E, E any](s S, start int) iter.Seq[E] {
	return func(yield func(E) bool) {
		for offset := range len(s) {
			if !yield(s[(start+offset)%len(s)]) {
				return
			}
		}
	}
}

func (f Seq[V]) Window(n int) iter.Seq[iter.Seq[V]] {
	return WindowSeq(n, f)
}
