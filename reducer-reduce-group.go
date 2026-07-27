package opiter

// ReduceGroup reduces values independently for each unique key in seq.
func ReduceGroup[Seq ~func(yield func(K, V) bool), K comparable, V, Sum any](
	reducer func(Sum, V) Sum,
	initial Sum,
	seq Seq,
) map[K]Sum {
	return InsertReduceGroup(make(map[K]Sum), reducer, initial, seq)
}

// InsertReduceGroup is like ReduceGroup but inserts the results into m.
func InsertReduceGroup[
	M ~map[K]Sum,
	Seq ~func(yield func(K, V) bool),
	K comparable,
	V, Sum any,
](
	m M,
	reducer func(Sum, V) Sum,
	initial Sum,
	seq Seq,
) M {
	for k, v := range seq {
		if _, ok := m[k]; !ok {
			m[k] = initial
		}
		m[k] = reducer(m[k], v)
	}
	return m
}
