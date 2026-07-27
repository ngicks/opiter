package opiter

import "iter"

// Step yields an infinite sequence beginning at initial and increasing by step.
func Step[N Numeric](initial, step N) iter.Seq[N] {
	return func(yield func(N) bool) {
		for n := initial; ; n += step {
			if !yield(n) {
				return
			}
		}
	}
}

// StepBy yields elements of values at indices beginning at initial and increasing by step.
func StepBy[S ~[]V, V any](initial, step int, values S) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		if initial < 0 {
			return
		}
		for i := initial; 0 <= i && i < len(values); i += step {
			if !yield(i, values[i]) {
				return
			}
		}
	}
}
