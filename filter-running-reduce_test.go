package opiter_test

import (
	"iter"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/ngicks/opiter"
	"github.com/ngicks/opiter/internal/testhelper"
)

func TestRunningReduce(t *testing.T) {
	reducer := func(sum, v, _ int) int { return sum + v }
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.RunningReduce(reducer, 0, s) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.WrapSeq(s).RunningReduce(reducer, 0).Iter() },
	}
	expected := []int{3, 4, 8, 9, 14, 23, 25, 31}
	testhelper.TestFilterStateless(t, testhelper.SampleSeq(), filters, expected, nil)

	joined := opiter.WrapSeq(testhelper.SampleSeq()).RunningReduce(
		func(sum string, v int, _ int) string { return sum + strconv.Itoa(v) },
		"",
	)
	if got := strings.Join(slices.Collect(joined.Iter()), ","); got != "3,31,314,3141,31415,314159,3141592,31415926" {
		t.Fatalf("generic RunningReduce method = %q", got)
	}
}

func TestRunningReduceStateful(t *testing.T) {
	reducer := func(_ int, v int, _ int) int { return v }
	filters := []func(iter.Seq[int]) iter.Seq[int]{
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.RunningReduce(reducer, 0, s) },
		func(s iter.Seq[int]) iter.Seq[int] { return opiter.WrapSeq(s).RunningReduce(reducer, 0).Iter() },
	}
	testhelper.TestFilterStateful(t, testhelper.SampleSeq(), filters, testhelper.SampleValues, nil)
}
