package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestSeqIter(t *testing.T) {
	seq := opiter.WrapSeq(testhelper.SampleSeq())
	testhelper.TestSourceStateless(
		t,
		[]func() iter.Seq[int]{seq.Iter, seq.IntoIter},
		testhelper.SampleValues,
		nil,
	)
}

func TestSeq2Iter2(t *testing.T) {
	seq := opiter.WrapSeq2(testhelper.SampleSeq2())
	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{seq.Iter2, seq.IntoIter2},
		testhelper.SamplePairs,
		nil,
	)
}
