// Package opiter defines option-aware iterator helpers.
//
// Helpers are grouped by how values flow through them:
//
//   - Sources create iterators from values, ranges, channels, and maps.
//   - Filters transform, combine, or select values from existing iterators.
//   - Reducers consume iterators to produce a result or side effect.
//
// Reducers that may not find a value return [opt.Option] instead of a zero
// value and a presence boolean. Package functions accept standard [iter.Seq]
// and [iter.Seq2] values; [Seq] and [Seq2] provide fluent counterparts for
// filters and reducers where Go's method type-parameter rules allow them.
package opiter
