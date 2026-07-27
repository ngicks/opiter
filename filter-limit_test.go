package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestLimitVariants(t *testing.T) {
	type testCase struct {
		name     string
		filter   func(iter.Seq[int]) iter.Seq[int]
		method   func(iter.Seq[int]) iter.Seq[int]
		expected []int
	}
	cases := []testCase{
		{"limit", func(s iter.Seq[int]) iter.Seq[int] { return opiter.Limit(3, s) }, func(s iter.Seq[int]) iter.Seq[int] { return opiter.WrapSeq(s).Limit(3).Iter() }, []int{3, 1, 4}},
		{"until", func(s iter.Seq[int]) iter.Seq[int] { return opiter.LimitUntil(func(v int) bool { return v < 5 }, s) }, func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).LimitUntil(func(v int) bool { return v < 5 }).Iter()
		}, []int{3, 1, 4, 1}},
		{"after", func(s iter.Seq[int]) iter.Seq[int] { return opiter.LimitAfter(func(v int) bool { return v < 5 }, s) }, func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).LimitAfter(func(v int) bool { return v < 5 }).Iter()
		}, []int{3, 1, 4, 1, 5}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			filters := []func(iter.Seq[int]) iter.Seq[int]{tc.filter, tc.method}
			testhelper.TestFilterStateless(t, testhelper.SampleSeq(), filters, tc.expected, nil)
		})
	}
}

func TestLimitVariants2(t *testing.T) {
	type testCase struct {
		name     string
		filter   func(iter.Seq2[string, int]) iter.Seq2[string, int]
		method   func(iter.Seq2[string, int]) iter.Seq2[string, int]
		expected []opiter.KV[string, int]
	}
	cases := []testCase{
		{"limit", func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.Limit2(2, s) }, func(s iter.Seq2[string, int]) iter.Seq2[string, int] { return opiter.WrapSeq2(s).Limit(2).Iter2() }, testhelper.SamplePairs[:2]},
		{"until", func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.LimitUntil2(func(_ string, v int) bool { return v < 4 }, s)
		}, func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).LimitUntil(func(_ string, v int) bool { return v < 4 }).Iter2()
		}, testhelper.SamplePairs[:2]},
		{"after", func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.LimitAfter2(func(_ string, v int) bool { return v < 4 }, s)
		}, func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).LimitAfter(func(_ string, v int) bool { return v < 4 }).Iter2()
		}, testhelper.SamplePairs[:3]},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			filters := []func(iter.Seq2[string, int]) iter.Seq2[string, int]{tc.filter, tc.method}
			testhelper.TestFilterStateless2(t, testhelper.SampleSeq2(), filters, tc.expected, nil)
		})
	}
}

func TestLimitVariantsStateful(t *testing.T) {
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.Limit(len(testhelper.SampleValues), s) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.LimitUntil(func(int) bool { return true }, s) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.LimitAfter(func(int) bool { return true }, s) },
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).Limit(len(testhelper.SampleValues)).Iter()
		},
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).LimitUntil(func(int) bool { return true }).Iter()
		},
		func(s iter.Seq[int]) iter.Seq[int] {
			return opiter.WrapSeq(s).LimitAfter(func(int) bool { return true }).Iter()
		},
	}
	testhelper.TestFilterStateful(t, testhelper.SampleSeq(), filters, testhelper.SampleValues, nil)
}

func TestLimitVariants2Stateful(t *testing.T) {
	filters := []func(iter.Seq2[string, int]) iter.Seq2[string, int]{
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.Limit2(len(testhelper.SamplePairs), s)
		},
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.LimitUntil2(func(string, int) bool { return true }, s)
		},
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.LimitAfter2(func(string, int) bool { return true }, s)
		},
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).Limit(len(testhelper.SamplePairs)).Iter2()
		},
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).LimitUntil(func(string, int) bool { return true }).Iter2()
		},
		func(s iter.Seq2[string, int]) iter.Seq2[string, int] {
			return opiter.WrapSeq2(s).LimitAfter(func(string, int) bool { return true }).Iter2()
		},
	}
	testhelper.TestFilterStateful2(t, testhelper.SampleSeq2(), filters, testhelper.SamplePairs, nil)
}
