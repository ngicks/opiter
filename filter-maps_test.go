package opiter_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestMapsBy(t *testing.T) {
	m := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	input := slices.Values([]string{"baz", "missing", "foo"})
	filters := []func(iter.Seq[string]) iter.Seq2[string, int]{
		func(keys iter.Seq[string]) iter.Seq2[string, int] {
			return opiter.MapsBy(m, keys)
		},
	}
	expected := []opiter.KV[string, int]{
		{K: "baz", V: 3},
		{K: "missing", V: 0},
		{K: "foo", V: 1},
	}
	testhelper.TestFilterStateless1To2(t, input, filters, expected, nil)
	testhelper.TestFilterStateful1To2(t, input, filters, expected, nil)
}
