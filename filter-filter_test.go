package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestFilter(t *testing.T) {
	predicate := func(v int) bool { return v%2 == 0 }
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(seq iter.Seq[int]) iter.Seq[int] { return opiter.Filter(predicate, seq) },
		func(seq iter.Seq[int]) iter.Seq[int] { return opiter.WrapSeq(seq).Filter(predicate).Iter() },
	}
	testhelper.TestFilterStateless(t, testhelper.SampleSeq(), filters, []int{4, 2, 6}, nil)
	testhelper.TestFilterStateful(t, testhelper.SampleSeq(), filters, []int{4, 2, 6}, nil)
}

func TestFilter2(t *testing.T) {
	predicate := func(_ string, v int) bool { return v%2 == 1 }
	filters := []func(iter.Seq2[string, int]) iter.Seq2[string, int]{
		func(seq iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.Filter2(predicate, seq) },
		func(seq iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(seq).Filter(predicate).Iter2()
		},
	}
	expected := []opiter.KV[string, int]{{"foo", 3}, {"bar", 1}, {"foo", 1}, {"qux", 5}}
	testhelper.TestFilterStateless2(t, testhelper.SampleSeq2(), filters, expected, nil)
	testhelper.TestFilterStateful2(t, testhelper.SampleSeq2(), filters, expected, nil)
}
