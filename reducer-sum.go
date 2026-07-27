package opiter

// Summable is the constraint of types that support the + operator.
type Summable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128 |
		~string
}

// Sum returns the sum of the values in seq.
func Sum[Seq ~func(yield func(S) bool), S Summable](seq Seq) S {
	var sum S
	for v := range seq {
		sum += v
	}
	return sum
}

// SumOf returns the sum of the values selected from seq by selector.
func SumOf[Seq ~func(yield func(V) bool), V any, S Summable](
	selector func(V) S,
	seq Seq,
) S {
	var sum S
	for v := range seq {
		sum += selector(v)
	}
	return sum
}

// SumOf returns the sum of the values selected from seq by selector.
func (seq Seq[V]) SumOf[S Summable](selector func(V) S) S {
	return SumOf(selector, seq)
}
