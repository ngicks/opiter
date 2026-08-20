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
