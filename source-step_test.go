package opiter_test

import (
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestStep(t *testing.T) {
	got := slices.Collect(opiter.Limit(4, opiter.Step(1, 2)))
	if !slices.Equal(got, []int{1, 3, 5, 7}) {
		t.Fatalf("Step() = %v", got)
	}
}

func TestStepBy(t *testing.T) {
	values := []string{"a", "b", "c", "d"}
	testhelper.TestSourceStateless2(
		t,
		[]func() opiter.Seq2[int, string]{
			func() opiter.Seq2[int, string] { return opiter.StepBy(1, 2, values) },
		},
		[]opiter.KV[int, string]{{K: 1, V: "b"}, {K: 3, V: "d"}},
		nil,
	)
	testhelper.TestSourceStateless2(
		t,
		[]func() opiter.Seq2[int, string]{
			func() opiter.Seq2[int, string] { return opiter.StepBy(-1, 2, values) },
		},
		nil,
		nil,
	)
}
