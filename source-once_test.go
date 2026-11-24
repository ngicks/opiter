package opiter_test

import (
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestOnce(t *testing.T) {
	testhelper.TestSourceStateless(
		t,
		[]func() opiter.Seq[int]{func() opiter.Seq[int] { return opiter.Once(3) }},
		[]int{3},
		nil,
	)
	testhelper.TestSourceStateless2(
		t,
		[]func() opiter.Seq2[string, int]{func() opiter.Seq2[string, int] {
			return opiter.Once2("three", 3)
		}},
		[]opiter.KV[string, int]{{K: "three", V: 3}},
		nil,
	)
}
