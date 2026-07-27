package opiter_test

import (
	"iter"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestRepeat(t *testing.T) {
	testhelper.TestSourceStateless(
		t,
		[]func() iter.Seq[int]{
			func() iter.Seq[int] { return opiter.Repeat(4, 3) },
			func() iter.Seq[int] { return opiter.RepeatFunc(func() int { return 4 }, 3) },
		},
		[]int{4, 4, 4},
		nil,
	)
	testhelper.TestSourceStateless2(
		t,
		[]func() iter.Seq2[string, int]{
			func() iter.Seq2[string, int] { return opiter.Repeat2("four", 4, 3) },
			func() iter.Seq2[string, int] {
				return opiter.RepeatFunc2(func() string { return "four" }, func() int { return 4 }, 3)
			},
		},
		[]opiter.KV[string, int]{
			{K: "four", V: 4},
			{K: "four", V: 4},
			{K: "four", V: 4},
		},
		nil,
	)
	testhelper.TestSourceStateless(
		t,
		[]func() iter.Seq[int]{func() iter.Seq[int] { return opiter.Repeat(4, 0) }},
		nil,
		nil,
	)
}

func TestRepeatInfiniteStops(t *testing.T) {
	count := 0
	opiter.Repeat(1, -1)(func(int) bool {
		count++
		return count < 5
	})
	if count != 5 {
		t.Errorf("Repeat yielded %d values, want 5", count)
	}

	count = 0
	opiter.RepeatFunc2(func() int { return count }, func() int { return count }, -1)(func(_, _ int) bool {
		count++
		return count < 5
	})
	if count != 5 {
		t.Errorf("RepeatFunc2 yielded %d pairs, want 5", count)
	}
}
