package testhelper

import (
	"iter"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
)

func TestFilterStateless[V any](
	t *testing.T,
	sources []func() iter.Seq[V],
	execpected []V,
	cmpOpt []goCmp.Option,
) {

}

func TestFilterStateful[V any](
	t *testing.T,
	sources []func() iter.Seq[V],
	execpected []V,
	cmpOpt []goCmp.Option,
) {
}
