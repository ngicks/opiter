package opiter_test

import (
	"errors"
	"iter"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"github.com/ngicks/option/opt"
	"gotest.tools/v3/assert"
)

func TestTryReducers(t *testing.T) {
	errStop := errors.New("stop")
	input := opiter.Values2([]opiter.KV[int, error]{
		{K: 3},
		{K: 1},
		{K: 4, V: errStop},
		{K: 9},
	})
	type findResult struct {
		Value opt.Option[opiter.KV[int, int]]
		Err   error
	}
	testhelper.TestReducer2(
		t,
		input,
		[]func(iter.Seq2[int, error]) findResult{
			func(seq iter.Seq2[int, error]) findResult {
				value, err := opiter.TryFind(func(v int) bool { return v == 1 }, seq)
				return findResult{Value: value, Err: err}
			},
		},
		findResult{Value: opt.Some(opiter.PackKV(1, 1))},
		[]goCmp.Option{compareOption[opiter.KV[int, int]](), testhelper.CompareErrorsIs},
	)

	type valuesResult struct {
		Values []int
		Err    error
	}
	expected := valuesResult{Values: []int{3, 1}, Err: errStop}
	testhelper.TestReducer2(
		t,
		input,
		[]func(iter.Seq2[int, error]) valuesResult{
			func(seq iter.Seq2[int, error]) (result valuesResult) {
				result.Err = opiter.TryForEach(func(v int) {
					result.Values = append(result.Values, v)
				}, seq)
				return result
			},
			func(seq iter.Seq2[int, error]) valuesResult {
				values, err := opiter.TryCollect(seq)
				return valuesResult{Values: values, Err: err}
			},
			func(seq iter.Seq2[int, error]) valuesResult {
				values, err := opiter.TryAppendSeq([]int{}, seq)
				return valuesResult{Values: values, Err: err}
			},
			func(seq iter.Seq2[int, error]) valuesResult {
				values, err := opiter.TryReduce(func(values []int, v int) []int {
					return append(values, v)
				}, []int{}, seq)
				return valuesResult{Values: values, Err: err}
			},
		},
		expected,
		[]goCmp.Option{testhelper.CompareErrorsIs},
	)
	found, err := opiter.TryFind(func(v int) bool { return v == 4 }, input)
	assert.Assert(t, found.IsNone())
	assert.ErrorIs(t, err, errStop)
}

func TestTryMapReducers(t *testing.T) {
	errStop := errors.New("stop")
	input := opiter.Values2([]opiter.KV[opiter.KV[string, int], error]{
		{K: opiter.PackKV("a", 1)},
		{K: opiter.PackKV("b", 2)},
		{K: opiter.PackKV("c", 3), V: errStop},
		{K: opiter.PackKV("d", 4)},
	})
	type result struct {
		Map map[string]int
		Err error
	}
	expected := result{Map: map[string]int{"a": 1, "b": 2}, Err: errStop}
	testhelper.TestReducer2(
		t,
		input,
		[]func(iter.Seq2[opiter.KV[string, int], error]) result{
			func(seq iter.Seq2[opiter.KV[string, int], error]) result {
				m, err := opiter.TryMapsCollect(seq)
				return result{Map: m, Err: err}
			},
			func(seq iter.Seq2[opiter.KV[string, int], error]) result {
				m := make(map[string]int)
				err := opiter.TryInsert(m, seq)
				return result{Map: m, Err: err}
			},
		},
		expected,
		[]goCmp.Option{testhelper.CompareErrorsIs},
	)
}
