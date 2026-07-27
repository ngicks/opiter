package opiter_test

import (
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestCycleVariants(t *testing.T) {
	expected := []int{1, 2, 1, 2, 1}
	seqs := []opiter.Seq[int]{
		opiter.WrapSeq(opiter.Limit(5, opiter.Cycle(slices.Values([]int{1, 2})))),
		opiter.WrapSeq(opiter.Limit(5, opiter.CycleBuffered(testhelper.OneShotSeq(1, 2)))),
		opiter.WrapSeq(slices.Values([]int{1, 2})).Cycle().Limit(5),
		opiter.WrapSeq(testhelper.OneShotSeq(1, 2)).CycleBuffered().Limit(5),
	}
	for i, seq := range seqs {
		if got := slices.Collect(seq.Iter()); !slices.Equal(got, expected) {
			t.Fatalf("seq[%d] = %v", i, got)
		}
	}
}

func TestCycleVariants2(t *testing.T) {
	input := opiter.Values2([]opiter.KV[int, int]{{1, 2}, {3, 4}})
	expected := []opiter.KV[int, int]{{1, 2}, {3, 4}, {1, 2}}
	seqs := []opiter.Seq2[int, int]{
		opiter.WrapSeq2(opiter.Limit2(3, opiter.Cycle2(input))),
		opiter.WrapSeq2(opiter.Limit2(3, opiter.CycleBuffered2(testhelper.OneShotSeq2(opiter.PackKV(1, 2), opiter.PackKV(3, 4))))),
		opiter.WrapSeq2(input).Cycle().Limit(3),
		opiter.WrapSeq2(testhelper.OneShotSeq2(opiter.PackKV(1, 2), opiter.PackKV(3, 4))).CycleBuffered().Limit(3),
	}
	for i, seq := range seqs {
		if got := opiter.Collect2(seq.Iter2()); !slices.Equal(got, expected) {
			t.Fatalf("seq[%d] = %v", i, got)
		}
	}
}
