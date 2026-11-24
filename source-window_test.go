package opiter_test

import (
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestWindow(t *testing.T) {
	values := []int{1, 2, 3, 4}
	testhelper.TestSourceStateless(
		t,
		[]func() opiter.Seq[[]int]{func() opiter.Seq[[]int] { return opiter.Window(values, 2) }},
		[][]int{{1, 2}, {2, 3}, {3, 4}},
		nil,
	)
	testhelper.TestSourceStateless(
		t,
		[]func() opiter.Seq[[]int]{func() opiter.Seq[[]int] { return opiter.Window(values, 0) }},
		nil,
		nil,
	)
}
