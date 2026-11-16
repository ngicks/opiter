package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestWrapFunc(t *testing.T) {
	stateless := opiter.WrapFunc(testhelper.SampleSeq)
	testhelper.TestSourceStateless(
		t,
		[]func() iter.Seq[int]{stateless.Iter, stateless.IntoIter},
		testhelper.SampleValues,
		nil,
	)

	stateful := opiter.WrapFunc(func() iter.Seq[int] {
		return testhelper.OneShotSeq(testhelper.SampleValues...)
	})
	testhelper.TestSourceStateful(
		t,
		[]func() iter.Seq[int]{stateful.Iter, stateful.IntoIter},
		testhelper.SampleValues,
		nil,
	)
}

func TestWrapFunc2(t *testing.T) {
	stateless := opiter.WrapFunc2(testhelper.SampleSeq2)
	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{stateless.Iter2, stateless.IntoIter2},
		testhelper.SamplePairs,
		nil,
	)

	stateful := opiter.WrapFunc2(func() iter.Seq2[string, int] {
		return testhelper.OneShotSeq2(testhelper.SamplePairs...)
	})
	testhelper.TestSourceStateful2(
		t,
		[]func() iter.Seq2[string, int]{stateful.Iter2, stateful.IntoIter2},
		testhelper.SamplePairs,
		nil,
	)
}
