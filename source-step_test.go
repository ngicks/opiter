package opiter_test

import (
	"slices"
	"testing"

	"github.com/ngicks/opiter"
)

func TestStep(t *testing.T) {
	got := slices.Collect(opiter.Limit(4, opiter.Step(1, 2)))
	if !slices.Equal(got, []int{1, 3, 5, 7}) {
		t.Fatalf("Step() = %v", got)
	}
}

func TestStepBy(t *testing.T) {
	got := opiter.Collect2(opiter.StepBy(1, 2, []string{"a", "b", "c", "d"}))
	if !slices.Equal(got, []opiter.KV[int, string]{{1, "b"}, {3, "d"}}) {
		t.Fatalf("StepBy() = %v", got)
	}
}
