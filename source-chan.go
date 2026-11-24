package opiter

import (
	"context"
)

// Chan returns an iterator over ch. Cancelling ctx or closing ch stops
// iteration.
func Chan[V any](ctx context.Context, ch <-chan V) Seq[V] {
	return func(yield func(V) bool) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok || !yield(v) {
						return
					}
				}
			}
		}
	}
}
