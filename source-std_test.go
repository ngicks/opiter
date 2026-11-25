package opiter_test

import (
	"cmp"
	"maps"
	"slices"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
	"gotest.tools/v3/assert"
)

func TestSlicesAll(t *testing.T) {
	values := []string{"a", "b", "c"}
	testhelper.TestSourceStateless2(
		t,
		[]func() opiter.Seq2[int, string]{
			func() opiter.Seq2[int, string] { return opiter.SlicesAll(values) },
		},
		[]opiter.KV[int, string]{{K: 0, V: "a"}, {K: 1, V: "b"}, {K: 2, V: "c"}},
		nil,
	)
}

func TestSlicesBackward(t *testing.T) {
	values := []string{"a", "b", "c"}
	testhelper.TestSourceStateless2(
		t,
		[]func() opiter.Seq2[int, string]{
			func() opiter.Seq2[int, string] { return opiter.SlicesBackward(values) },
		},
		[]opiter.KV[int, string]{{K: 2, V: "c"}, {K: 1, V: "b"}, {K: 0, V: "a"}},
		nil,
	)
}

func TestSlicesValues(t *testing.T) {
	values := []string{"a", "b", "c"}
	testhelper.TestSourceStateless(
		t,
		[]func() opiter.Seq[string]{
			func() opiter.Seq[string] { return opiter.SlicesValues(values) },
		},
		values,
		nil,
	)
}

func TestSlicesChunk(t *testing.T) {
	type ints []int
	values := ints{1, 2, 3, 4, 5}
	testhelper.TestSourceStateless(
		t,
		[]func() opiter.Seq[ints]{
			func() opiter.Seq[ints] { return opiter.SlicesChunk(values, 2) },
		},
		[]ints{{1, 2}, {3, 4}, {5}},
		nil,
	)
}

func TestSlicesChunkPanicsOnNonPositiveSize(t *testing.T) {
	defer func() {
		assert.Assert(t, recover() != nil)
	}()
	opiter.SlicesChunk([]int{1}, 0)
}

// Map iteration order is unspecified,
// so the contract tests use a single-entry map
// and the multi-entry tests compare sorted results.

func TestMapsContract(t *testing.T) {
	m := map[string]int{"a": 1}
	testhelper.TestSourceStateless2(
		t,
		[]func() opiter.Seq2[string, int]{
			func() opiter.Seq2[string, int] { return opiter.MapsAll(m) },
		},
		[]opiter.KV[string, int]{{K: "a", V: 1}},
		nil,
	)
	testhelper.TestSourceStateless(
		t,
		[]func() opiter.Seq[string]{
			func() opiter.Seq[string] { return opiter.MapsKeys(m) },
		},
		[]string{"a"},
		nil,
	)
	testhelper.TestSourceStateless(
		t,
		[]func() opiter.Seq[int]{
			func() opiter.Seq[int] { return opiter.MapsValues(m) },
		},
		[]int{1},
		nil,
	)
}

func TestMapsAll(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := opiter.MapsAll(m).CollectSlice()
	slices.SortFunc(got, func(l, r opiter.KV[string, int]) int { return cmp.Compare(l.K, r.K) })
	assert.DeepEqual(
		t,
		[]opiter.KV[string, int]{{K: "a", V: 1}, {K: "b", V: 2}, {K: "c", V: 3}},
		got,
	)
	assert.DeepEqual(t, m, opiter.MapsAll(m).Collect(maps.Collect))
}

func TestMapsKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	assert.DeepEqual(t, []string{"a", "b", "c"}, opiter.MapsKeys(m).Collect(slices.Sorted))
}

func TestMapsValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	assert.DeepEqual(t, []int{1, 2, 3}, opiter.MapsValues(m).Collect(slices.Sorted))
}

func isComma(r rune) bool { return r == ',' }

func toBytes(ss ...string) [][]byte {
	out := make([][]byte, len(ss))
	for i, s := range ss {
		out[i] = []byte(s)
	}
	return out
}

func TestStringsSingleUse(t *testing.T) {
	for _, tc := range []struct {
		name     string
		source   func() opiter.Seq[string]
		expected []string
	}{
		{
			"Lines",
			func() opiter.Seq[string] { return opiter.StringsLines("a\nb\nc") },
			[]string{"a\n", "b\n", "c"},
		},
		{
			"SplitSeq",
			func() opiter.Seq[string] { return opiter.StringsSplitSeq("a,b,c", ",") },
			[]string{"a", "b", "c"},
		},
		{
			"SplitAfterSeq",
			func() opiter.Seq[string] { return opiter.StringsSplitAfterSeq("a,b,c", ",") },
			[]string{"a,", "b,", "c"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testhelper.TestSourceSingleUse(
				t,
				[]func() opiter.Seq[string]{tc.source},
				tc.expected,
				nil,
			)
		})
	}
}

func TestStringsFields(t *testing.T) {
	testhelper.TestSourceStateless(
		t,
		[]func() opiter.Seq[string]{
			func() opiter.Seq[string] { return opiter.StringsFieldsSeq(" a  b c ") },
			func() opiter.Seq[string] { return opiter.StringsFieldsFuncSeq(",a,,b,c", isComma) },
		},
		[]string{"a", "b", "c"},
		nil,
	)
}

func TestBytesSingleUse(t *testing.T) {
	for _, tc := range []struct {
		name     string
		source   func() opiter.Seq[[]byte]
		expected [][]byte
	}{
		{
			"Lines",
			func() opiter.Seq[[]byte] { return opiter.BytesLines([]byte("a\nb\nc")) },
			toBytes("a\n", "b\n", "c"),
		},
		{
			"SplitSeq",
			func() opiter.Seq[[]byte] { return opiter.BytesSplitSeq([]byte("a,b,c"), []byte(",")) },
			toBytes("a", "b", "c"),
		},
		{
			"SplitAfterSeq",
			func() opiter.Seq[[]byte] {
				return opiter.BytesSplitAfterSeq([]byte("a,b,c"), []byte(","))
			},
			toBytes("a,", "b,", "c"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testhelper.TestSourceSingleUse(
				t,
				[]func() opiter.Seq[[]byte]{tc.source},
				tc.expected,
				nil,
			)
		})
	}
}

func TestBytesFields(t *testing.T) {
	testhelper.TestSourceStateless(
		t,
		[]func() opiter.Seq[[]byte]{
			func() opiter.Seq[[]byte] { return opiter.BytesFieldsSeq([]byte(" a  b c ")) },
			func() opiter.Seq[[]byte] {
				return opiter.BytesFieldsFuncSeq([]byte(",a,,b,c"), isComma)
			},
		},
		toBytes("a", "b", "c"),
		nil,
	)
}
