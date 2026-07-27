package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestEmpty(t *testing.T) {
	testhelper.TestSourceStateless(
		t,
		[]func() iter.Seq[int]{opiter.Empty[int]},
		nil,
		nil,
	)
	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{opiter.Empty2[string, int]},
		nil,
		nil,
	)
}
