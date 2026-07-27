package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestTapVariants(t *testing.T) {
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.Tap(func(int) {}, s) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.WrapSeq(s).Tap(func(int) {}).Iter() },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.TapLast(func() {}, s) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.WrapSeq(s).TapLast(func() {}).Iter() },
	}
	testhelper.TestFilterStateless(t, testhelper.SampleSeq(), filters, testhelper.SampleValues, nil)
}

func TestTapVariants2(t *testing.T) {
	filters := []func(iter.Seq2[string, int]) iter.Seq2[string, int]{
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.Tap2(func(string, int) {}, s) },
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).Tap(func(string, int) {}).Iter2()
		},
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.TapLast2(func() {}, s) },
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).TapLast(func() {}).Iter2()
		},
	}
	testhelper.TestFilterStateless2(t, testhelper.SampleSeq2(), filters, testhelper.SamplePairs, nil)
}
