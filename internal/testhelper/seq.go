package testhelper

import (
	"iter"

	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/opiter"
)

// takeN consumes n values from seq, then breaks. The value delivered to the
// breaking iteration is returned as rej and is not included in got.
// Full collection is not takeN's job; use slices.Collect / opiter.Collect2.
// takeN ranges over seq, not calls it directly, because it is the ultimate
// consumer: the runtime's range checks keep the producer honest (e.g. panic
// when yield is called again after it returned false).
func takeN[V any](seq iter.Seq[V], n int) (got []V, rej V, rejected bool) {
	for v := range seq {
		if len(got) >= n {
			rej = v
			rejected = true
			break
		}
		got = append(got, v)
	}
	return
}

// takeN2 is [takeN] for iter.Seq2-shaped iterators.
func takeN2[K, V any](seq iter.Seq2[K, V], n int) (got []opiter.KV[K, V], rej opiter.KV[K, V], rejected bool) {
	for k, v := range seq {
		if len(got) >= n {
			rej = opiter.PackKV(k, v)
			rejected = true
			break
		}
		got = append(got, opiter.PackKV(k, v))
	}
	return
}

// diffValues compares slices treating nil and empty as equal at the top level.
func diffValues[V any](expected, got []V, cmpOpt []goCmp.Option) string {
	if len(expected) == 0 && len(got) == 0 {
		return ""
	}
	return goCmp.Diff(expected, got, cmpOpt...)
}
