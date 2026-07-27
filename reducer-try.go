package opiter

import "github.com/ngicks/option/opt"

// TryFind returns the first value satisfying fn and its zero-based index,
// stopping if seq yields a non-nil error.
func TryFind[Seq ~func(yield func(V, error) bool), V any](
	fn func(V) bool,
	seq Seq,
) (opt.Option[KV[int, V]], error) {
	i := 0
	for v, err := range seq {
		if err != nil {
			return opt.None[KV[int, V]](), err
		}
		if fn(v) {
			return opt.Some(PackKV(i, v)), nil
		}
		i++
	}
	return opt.None[KV[int, V]](), nil
}

// TryForEach calls fn with each value, stopping if seq yields a non-nil error.
func TryForEach[Seq ~func(yield func(V, error) bool), V any](fn func(V), seq Seq) error {
	for v, err := range seq {
		if err != nil {
			return err
		}
		fn(v)
	}
	return nil
}

// TryReduce combines values using fn, stopping if seq yields a non-nil error.
func TryReduce[Seq ~func(yield func(V, error) bool), V, Sum any](
	fn func(Sum, V) Sum,
	initial Sum,
	seq Seq,
) (Sum, error) {
	sum := initial
	for v, err := range seq {
		if err != nil {
			return sum, err
		}
		sum = fn(sum, v)
	}
	return sum, nil
}

// TryCollect collects values until seq yields a non-nil error.
func TryCollect[Seq ~func(yield func(V, error) bool), V any](seq Seq) ([]V, error) {
	return TryAppendSeq[[]V](nil, seq)
}

// TryAppendSeq appends values to s until seq yields a non-nil error.
func TryAppendSeq[S ~[]V, Seq ~func(yield func(V, error) bool), V any](
	s S,
	seq Seq,
) (S, error) {
	for v, err := range seq {
		if err != nil {
			return s, err
		}
		s = append(s, v)
	}
	return s, nil
}

// TryMapsCollect collects key-value pairs until seq yields a non-nil error.
func TryMapsCollect[Seq ~func(yield func(KV[K, V], error) bool), K comparable, V any](
	seq Seq,
) (map[K]V, error) {
	m := make(map[K]V)
	err := TryInsert(m, seq)
	return m, err
}

// TryInsert inserts key-value pairs into m until seq yields a non-nil error.
func TryInsert[
	M ~map[K]V,
	Seq ~func(yield func(KV[K, V], error) bool),
	K comparable,
	V any,
](
	m M,
	seq Seq,
) error {
	for kv, err := range seq {
		if err != nil {
			return err
		}
		m[kv.K] = kv.V
	}
	return nil
}
