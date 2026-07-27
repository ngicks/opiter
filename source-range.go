package opiter

import "iter"

// Numeric is a numeric type supported by [Range] and [RangeInclusive].
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Range returns an iterator over sequential values in the half-open interval
// [start, end). Values increase by one when start is less than end and decrease
// by one when start is greater than end.
func Range[T Numeric](start, end T) iter.Seq[T] {
	return rangeInclusive(start, end, true, false)
}

// RangeInclusive is like [Range] with control over whether each endpoint is
// included.
func RangeInclusive[T Numeric](start, end T, includeStart, includeEnd bool) iter.Seq[T] {
	return rangeInclusive(start, end, includeStart, includeEnd)
}

func rangeInclusive[T Numeric](start, end T, includeStart, includeEnd bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		start, end := start, end
		switch {
		default:
			if includeStart && includeEnd {
				yield(start)
			}
		case start < end:
			if !includeStart {
				start++
			}
			if !includeEnd {
				end--
			}
			for i := start; i <= end; i++ {
				if !yield(i) || i == end {
					return
				}
			}
		case start > end:
			if !includeStart {
				start--
			}
			if !includeEnd {
				end++
			}
			for i := start; i >= end; i-- {
				if !yield(i) || i == end {
					return
				}
			}
		}
	}
}
