package opiter_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestMap(t *testing.T) {
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(seq iter.Seq[int]) iter.Seq[int] { return opiter.Map(func(v int) int { return v * 2 }, seq) },
		func(seq iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(seq).Map(func(v int) int { return v * 2 }).Iter()
		},
	}
	expected := []int{6, 2, 8, 2, 10, 18, 4, 12}
	testhelper.TestFilterStateless(t, testhelper.SampleSeq(), filters, expected, nil)
	testhelper.TestFilterStateful(t, testhelper.SampleSeq(), filters, expected, nil)
}

func TestMap2(t *testing.T) {
	fn := func(k string, v int) (string, int) { return k + "!", v * 2 }
	filters := []func(iter.Seq2[string, int]) iter.Seq2[string, int]{
		func(seq iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.Map2(fn, seq) },
		func(seq iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.WrapSeq2(seq).Map(fn).Iter2() },
	}
	expected := []opiter.KV[string, int]{
		{"foo!", 6}, {"bar!", 2}, {"baz!", 8}, {"foo!", 2}, {"qux!", 10},
	}
	testhelper.TestFilterStateless2(t, testhelper.SampleSeq2(), filters, expected, nil)
	testhelper.TestFilterStateful2(t, testhelper.SampleSeq2(), filters, expected, nil)
}

func TestMapMethodsTransformTypes(t *testing.T) {
	got := opiter.WrapSeq(testhelper.SampleSeq()).Map(strconv.Itoa)
	if values := slices.Collect(got.Iter()); len(values) != len(testhelper.SampleValues) || values[0] != "3" {
		t.Fatalf("Map method = %v", values)
	}
	pairs := opiter.WrapSeq2(testhelper.SampleSeq2()).Map(func(k string, v int) (int, string) {
		return v, k
	})
	if values := opiter.Collect2(pairs.Iter2()); len(values) != len(testhelper.SamplePairs) || values[0] != (opiter.KV[int, string]{3, "foo"}) {
		t.Fatalf("Map method = %v", values)
	}
}
