package opiter_test

import (
	"math"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestRange(t *testing.T) {
	type testCase struct {
		name     string
		source   func() opiter.Seq[int]
		expected []int
	}
	tests := []testCase{
		{
			name:     "ascending",
			source:   func() opiter.Seq[int] { return opiter.Range(1, 5) },
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "descending",
			source:   func() opiter.Seq[int] { return opiter.Range(2, -2) },
			expected: []int{2, 1, 0, -1},
		},
		{
			name:   "equal",
			source: func() opiter.Seq[int] { return opiter.Range(1, 1) },
		},
		{
			name:     "exclusive",
			source:   func() opiter.Seq[int] { return opiter.RangeInclusive(1, 5, false, false) },
			expected: []int{2, 3, 4},
		},
		{
			name:     "inclusive descending",
			source:   func() opiter.Seq[int] { return opiter.RangeInclusive(2, -2, true, true) },
			expected: []int{2, 1, 0, -1, -2},
		},
		{
			name:     "exclusive descending",
			source:   func() opiter.Seq[int] { return opiter.RangeInclusive(2, -2, false, false) },
			expected: []int{1, 0, -1},
		},
		{
			name:     "equal included",
			source:   func() opiter.Seq[int] { return opiter.RangeInclusive(1, 1, true, true) },
			expected: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testhelper.TestSourceStateless(t, []func() opiter.Seq[int]{tt.source}, tt.expected, nil)
		})
	}
}

func TestRangeIntegerBounds(t *testing.T) {
	if got := opiter.RangeInclusive(uint8(math.MaxUint8-1), uint8(math.MaxUint8), true, true).
		CollectSlice(); !slices.Equal(
		got,
		[]uint8{math.MaxUint8 - 1, math.MaxUint8},
	) {
		t.Errorf("ascending at maximum: got %v", got)
	}
	if got := opiter.RangeInclusive(uint8(1), uint8(0), true, true).CollectSlice(); !slices.Equal(
		got,
		[]uint8{1, 0},
	) {
		t.Errorf("descending at minimum: got %v", got)
	}
}
