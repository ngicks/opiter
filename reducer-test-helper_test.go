package opiter_test

import (
	goCmp "github.com/google/go-cmp/cmp"
	"github.com/ngicks/option/opt"
)

func compareOption[V comparable]() goCmp.Option {
	return goCmp.Comparer(func(a, b opt.Option[V]) bool {
		return a.IsSome() == b.IsSome() && (a.IsNone() || a.Value() == b.Value())
	})
}
