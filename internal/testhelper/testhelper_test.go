package testhelper

import (
	"fmt"
	"iter"
	"reflect"
	"slices"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSourceHelpers(t *testing.T) {
	TestSourceStateless(
		t,
		[]func() iter.Seq[int]{
			SampleSeq,
			func() iter.Seq[int] { return SampleSeq() },
		},
		SampleValues,
		nil,
	)
	TestSourceStateless(t, []func() iter.Seq[int]{func() iter.Seq[int] { return OneShotSeq[int]() }}, nil, nil)
	TestSourceStateful(
		t,
		[]func() iter.Seq[int]{func() iter.Seq[int] { return OneShotSeq(SampleValues...) }},
		SampleValues,
		nil,
	)

	TestSourceStateless2(t, []func() iter.Seq2[string, int]{SampleSeq2}, SamplePairs, nil)
	TestSourceStateful2(
		t,
		[]func() iter.Seq2[string, int]{func() iter.Seq2[string, int] { return OneShotSeq2(SamplePairs...) }},
		SamplePairs,
		nil,
	)
}

func TestFilterHelpers(t *testing.T) {
	double := func(s iter.Seq[int]) iter.Seq[int] {
		return func(yield func(int) bool) {
			s(func(v int) bool { return yield(v * 2) })
		}
	}
	even := func(s iter.Seq[int]) iter.Seq[int] {
		return func(yield func(int) bool) {
			s(func(v int) bool {
				if v%2 != 0 {
					return true
				}
				return yield(v)
			})
		}
	}
	reverse := func(s iter.Seq[int]) iter.Seq[int] {
		return func(yield func(int) bool) {
			// buffering severs the chain: the downstream range loop never
			// touches s, so this collection is s's ultimate consumer.
			vals := slices.Collect(s)
			for i := len(vals) - 1; i >= 0; i-- {
				if !yield(vals[i]) {
					return
				}
			}
		}
	}

	doubled := []int{6, 2, 8, 2, 10, 18, 4, 12}
	evens := []int{4, 2, 6}
	reversed := []int{6, 2, 9, 5, 1, 4, 1, 3}

	TestFilterStateless(t, SampleSeq(), []func(iter.Seq[int]) iter.Seq[int]{double}, doubled, nil)
	TestFilterStateless(t, SampleSeq(), []func(iter.Seq[int]) iter.Seq[int]{even}, evens, nil)
	// buffering filter: stateless only
	TestFilterStateless(t, SampleSeq(), []func(iter.Seq[int]) iter.Seq[int]{reverse}, reversed, nil)
	TestFilterStateful(t, SampleSeq(), []func(iter.Seq[int]) iter.Seq[int]{double}, doubled, nil)
	TestFilterStateful(t, SampleSeq(), []func(iter.Seq[int]) iter.Seq[int]{even}, evens, nil)

	swap := func(s iter.Seq2[string, int]) iter.Seq2[int, string] {
		return func(yield func(int, string) bool) {
			s(func(k string, v int) bool { return yield(v, k) })
		}
	}
	swapped := []opiter.KV[int, string]{
		{K: 3, V: "foo"},
		{K: 1, V: "bar"},
		{K: 4, V: "baz"},
		{K: 1, V: "foo"},
		{K: 5, V: "qux"},
	}
	TestFilterStateless2(t, SampleSeq2(), []func(iter.Seq2[string, int]) iter.Seq2[int, string]{swap}, swapped, nil)
	TestFilterStateful2(t, SampleSeq2(), []func(iter.Seq2[string, int]) iter.Seq2[int, string]{swap}, swapped, nil)

	join := func(s iter.Seq2[string, int]) iter.Seq[string] {
		return func(yield func(string) bool) {
			s(func(k string, v int) bool { return yield(fmt.Sprintf("%s:%d", k, v)) })
		}
	}
	joined := []string{"foo:3", "bar:1", "baz:4", "foo:1", "qux:5"}
	TestFilterStateless2To1(t, SampleSeq2(), []func(iter.Seq2[string, int]) iter.Seq[string]{join}, joined, nil)
	TestFilterStateful2To1(t, SampleSeq2(), []func(iter.Seq2[string, int]) iter.Seq[string]{join}, joined, nil)

	fan := func(s iter.Seq[int]) iter.Seq2[int, int] {
		return func(yield func(int, int) bool) {
			s(func(v int) bool { return yield(v, v*2) })
		}
	}
	fanned := make([]opiter.KV[int, int], len(SampleValues))
	for i, v := range SampleValues {
		fanned[i] = opiter.KV[int, int]{K: v, V: v * 2}
	}
	TestFilterStateless1To2(t, SampleSeq(), []func(iter.Seq[int]) iter.Seq2[int, int]{fan}, fanned, nil)
	TestFilterStateful1To2(t, SampleSeq(), []func(iter.Seq[int]) iter.Seq2[int, int]{fan}, fanned, nil)
}

func TestReducerHelpers(t *testing.T) {
	sum := func(s iter.Seq[int]) int {
		total := 0
		for v := range s {
			total += v
		}
		return total
	}
	TestReducer(t, SampleSeq(), []func(iter.Seq[int]) int{sum}, 31, nil)
	TestReducer(t, OneShotSeq[int](), []func(iter.Seq[int]) int{sum}, 0, nil)

	concatKeys := func(s iter.Seq2[string, int]) string {
		var out string
		for k := range s {
			out += k
		}
		return out
	}
	TestReducer2(t, SampleSeq2(), []func(iter.Seq2[string, int]) string{concatKeys}, "foobarbazfooqux", nil)
}

type codedErr struct{ code int }

func (e codedErr) Error() string { return fmt.Sprintf("code %d", e.code) }

// notDeepEqual inverts [is.DeepEqual]; gotest.tools has no negative counterpart.
func notDeepEqual(x, y any, opts ...goCmp.Option) is.Comparison {
	return func() is.Result {
		if is.DeepEqual(x, y, opts...)().Success() {
			return is.ResultFailure(fmt.Sprintf("values are unexpectedly deep equal: %v, %v", x, y))
		}
		return is.ResultSuccess
	}
}

func TestComparators(t *testing.T) {
	wrapped := fmt.Errorf("wrap: %w", ErrSample1)
	assert.DeepEqual(t, []error{wrapped}, []error{ErrSample1}, CompareErrorsIs)
	assert.Assert(
		t,
		notDeepEqual([]error{ErrSample1}, []error{ErrSample2}, CompareErrorsIs),
		"CompareErrorsIs: ErrSample1 should not equal ErrSample2",
	)
	assert.DeepEqual(t, []error{nil}, []error{nil}, CompareErrorsIs)
	assert.Assert(
		t,
		notDeepEqual([]error{ErrSample1}, []error{nil}, CompareErrorsIs),
		"CompareErrorsIs: non-nil should not equal nil",
	)

	assert.DeepEqual(
		t,
		[]error{fmt.Errorf("wrap: %w", codedErr{code: 1})},
		[]error{codedErr{code: 2}},
		CompareErrorsAs,
	)
	assert.Assert(
		t,
		notDeepEqual([]error{codedErr{code: 1}}, []error{ErrSample1}, CompareErrorsAs),
		"CompareErrorsAs: codedErr should not equal ErrSample1",
	)

	assert.DeepEqual(t, reflect.ValueOf(1), reflect.ValueOf(1), CompareReflectValue)
	assert.Assert(
		t,
		notDeepEqual(reflect.ValueOf(1), reflect.ValueOf(2), CompareReflectValue),
		"CompareReflectValue: 1 should not equal 2",
	)
}
