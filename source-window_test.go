package opiter_test

import (
	"slices"
	"testing"

	"github.com/ngicks/opiter"
)

func TestWindow(t *testing.T) {
	got := slices.Collect(opiter.Window([]int{1, 2, 3, 4}, 2))
	expected := [][]int{{1, 2}, {2, 3}, {3, 4}}
	if !slices.EqualFunc(got, expected, slices.Equal) {
		t.Fatalf("Window() = %v", got)
	}
}
