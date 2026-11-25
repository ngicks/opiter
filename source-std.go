package opiter

import (
	"bytes"
	"maps"
	"slices"
	"strings"
)

// SlicesAll wraps [slices.All].
// It returns an iterator over index-value pairs in s.
func SlicesAll[S ~[]E, E any](s S) Seq2[int, E] {
	return Seq2[int, E](slices.All(s))
}

// SlicesBackward wraps [slices.Backward].
// It returns an iterator over index-value pairs in s, from the last index to the first.
func SlicesBackward[S ~[]E, E any](s S) Seq2[int, E] {
	return Seq2[int, E](slices.Backward(s))
}

// SlicesValues wraps [slices.Values].
// It returns an iterator over the elements of s.
func SlicesValues[S ~[]E, E any](s S) Seq[E] {
	return Seq[E](slices.Values(s))
}

// SlicesChunk wraps [slices.Chunk].
// It returns an iterator over consecutive subslices of up to n elements of s.
//
// It panics if n is less than 1, as [slices.Chunk] does.
func SlicesChunk[S ~[]E, E any](s S, n int) Seq[S] {
	return Seq[S](slices.Chunk(s, n))
}

// MapsAll wraps [maps.All].
// It returns an iterator over key-value pairs of m in unspecified order.
func MapsAll[M ~map[K]V, K comparable, V any](m M) Seq2[K, V] {
	return Seq2[K, V](maps.All(m))
}

// MapsKeys wraps [maps.Keys].
// It returns an iterator over keys of m in unspecified order.
func MapsKeys[M ~map[K]V, K comparable, V any](m M) Seq[K] {
	return Seq[K](maps.Keys(m))
}

// MapsValues wraps [maps.Values].
// It returns an iterator over values of m in unspecified order.
func MapsValues[M ~map[K]V, K comparable, V any](m M) Seq[V] {
	return Seq[V](maps.Values(m))
}

// StringsLines wraps [strings.Lines].
// It returns an iterator over the newline-terminated lines in s.
//
// The returned iterator is single-use, as the wrapped one is.
func StringsLines(s string) Seq[string] {
	return Seq[string](strings.Lines(s))
}

// StringsSplitSeq wraps [strings.SplitSeq].
// It returns an iterator over substrings of s separated by sep.
//
// The returned iterator is single-use, as the wrapped one is.
func StringsSplitSeq(s, sep string) Seq[string] {
	return Seq[string](strings.SplitSeq(s, sep))
}

// StringsSplitAfterSeq wraps [strings.SplitAfterSeq].
// It returns an iterator over substrings of s split after each instance of sep.
//
// The returned iterator is single-use, as the wrapped one is.
func StringsSplitAfterSeq(s, sep string) Seq[string] {
	return Seq[string](strings.SplitAfterSeq(s, sep))
}

// StringsFieldsSeq wraps [strings.FieldsSeq].
// It returns an iterator over substrings of s split around runs of whitespace.
func StringsFieldsSeq(s string) Seq[string] {
	return Seq[string](strings.FieldsSeq(s))
}

// StringsFieldsFuncSeq wraps [strings.FieldsFuncSeq].
// It returns an iterator over substrings of s split around runs of code points satisfying f.
func StringsFieldsFuncSeq(s string, f func(rune) bool) Seq[string] {
	return Seq[string](strings.FieldsFuncSeq(s, f))
}

// BytesLines wraps [bytes.Lines].
// It returns an iterator over the newline-terminated lines in s.
//
// The returned iterator is single-use, as the wrapped one is.
func BytesLines(s []byte) Seq[[]byte] {
	return Seq[[]byte](bytes.Lines(s))
}

// BytesSplitSeq wraps [bytes.SplitSeq].
// It returns an iterator over subslices of s separated by sep.
//
// The returned iterator is single-use, as the wrapped one is.
func BytesSplitSeq(s, sep []byte) Seq[[]byte] {
	return Seq[[]byte](bytes.SplitSeq(s, sep))
}

// BytesSplitAfterSeq wraps [bytes.SplitAfterSeq].
// It returns an iterator over subslices of s split after each instance of sep.
//
// The returned iterator is single-use, as the wrapped one is.
func BytesSplitAfterSeq(s, sep []byte) Seq[[]byte] {
	return Seq[[]byte](bytes.SplitAfterSeq(s, sep))
}

// BytesFieldsSeq wraps [bytes.FieldsSeq].
// It returns an iterator over subslices of s split around runs of whitespace.
func BytesFieldsSeq(s []byte) Seq[[]byte] {
	return Seq[[]byte](bytes.FieldsSeq(s))
}

// BytesFieldsFuncSeq wraps [bytes.FieldsFuncSeq].
// It returns an iterator over subslices of s split around runs of code points satisfying f.
func BytesFieldsFuncSeq(s []byte, f func(rune) bool) Seq[[]byte] {
	return Seq[[]byte](bytes.FieldsFuncSeq(s, f))
}
