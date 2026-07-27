package opiter_test

import (
	"context"
	"iter"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestChan(t *testing.T) {
	testhelper.TestSourceStateful(
		t,
		[]func() iter.Seq[int]{func() iter.Seq[int] {
			ch := make(chan int, 4)
			for _, v := range []int{1, 2, 3, 4} {
				ch <- v
			}
			close(ch)
			return opiter.Chan(context.Background(), ch)
		}},
		[]int{1, 2, 3, 4},
		nil,
	)
}

func TestChanCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	ch := make(chan int, 1)
	ch <- 1
	if got := slices.Collect(opiter.Chan(ctx, ch)); len(got) != 0 {
		t.Errorf("canceled iterator yielded %v", got)
	}
}
