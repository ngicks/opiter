package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestSkipVariants(t *testing.T) {
	type testCase struct {
		name     string
		filters  []func(iter.Seq[int]) iter.Seq[int]
		expected []int
	}
	cases := []testCase{
		{"skip", []func(iter.Seq[int]) iter.Seq[int]{func(s iter.Seq[int]) iter.Seq[int] { return opiter.Skip(2, s) }, func(s iter.Seq[int]) iter.Seq[int] { return opiter.WrapSeq(s).Skip(2).Iter() }}, []int{4, 1, 5, 9, 2, 6}},
		{"last", []func(iter.Seq[int]) iter.Seq[int]{func(s iter.Seq[int]) iter.Seq[int] { return opiter.SkipLast(2, s) }, func(s iter.Seq[int]) iter.Seq[int] { return opiter.WrapSeq(s).SkipLast(2).Iter() }}, []int{3, 1, 4, 1, 5, 9}},
		{"while", []func(iter.Seq[int]) iter.Seq[int]{func(s iter.Seq[int]) iter.Seq[int] { return opiter.SkipWhile(func(v int) bool { return v < 4 }, s) }, func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).SkipWhile(func(v int) bool { return v < 4 }).Iter()
		}}, []int{4, 1, 5, 9, 2, 6}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testhelper.TestFilterStateless(t, testhelper.SampleSeq(), tc.filters, tc.expected, nil)
		})
	}
}

func TestSkipVariants2(t *testing.T) {
	type testCase struct {
		name     string
		filters  []func(iter.Seq2[string, int]) iter.Seq2[string, int]
		expected []opiter.KV[string, int]
	}
	cases := []testCase{
		{"skip", []func(iter.Seq2[string, int]) iter.Seq2[string, int]{func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.Skip2(2, s) }, func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.WrapSeq2(s).Skip(2).Iter2() }}, testhelper.SamplePairs[2:]},
		{"last", []func(iter.Seq2[string, int]) iter.Seq2[string, int]{func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.SkipLast2(2, s) }, func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.WrapSeq2(s).SkipLast(2).Iter2() }}, testhelper.SamplePairs[:3]},
		{"while", []func(iter.Seq2[string, int]) iter.Seq2[string, int]{func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.SkipWhile2(func(_ string, v int) bool { return v < 4 }, s)
		}, func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).SkipWhile(func(_ string, v int) bool { return v < 4 }).Iter2()
		}}, testhelper.SamplePairs[2:]},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testhelper.TestFilterStateless2(t, testhelper.SampleSeq2(), tc.filters, tc.expected, nil)
		})
	}
}

func TestSkipVariantsStateful(t *testing.T) {
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.Skip(0, s) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.SkipWhile(func(int) bool { return false }, s) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.WrapSeq(s).Skip(0).Iter() },
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).SkipWhile(func(int) bool { return false }).Iter()
		},
	}
	testhelper.TestFilterStateful(t, testhelper.SampleSeq(), filters, testhelper.SampleValues, nil)
}

func TestSkipVariants2Stateful(t *testing.T) {
	filters := []func(iter.Seq2[string, int]) iter.Seq2[string, int]{
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.Skip2(0, s) },
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.SkipWhile2(func(string, int) bool { return false }, s)
		},
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.WrapSeq2(s).Skip(0).Iter2() },
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).SkipWhile(func(string, int) bool { return false }).Iter2()
		},
	}
	testhelper.TestFilterStateful2(t, testhelper.SampleSeq2(), filters, testhelper.SamplePairs, nil)
}
