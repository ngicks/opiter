package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestOnce(t *testing.T) {
	testhelper.TestSourceStateless(
		t,
		[]func() iter.Seq[int]{func() iter.Seq[int] { return opiter.Once(3) }},
		[]int{3},
		nil,
	)
	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{func() iter.Seq2[string, int] {
			return opiter.Once2("three", 3)
		}},
		[]opiter.KV[string, int]{{K: "three", V: 3}},
		nil,
	)
}
