package testhelper

import (
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
)

func testSource[V any, Iter ~func(yield func(V) bool)](
	t *testing.T,
	sources []func() Iter,
	execpected []V,
	cmpOpt []goCmp.Option,
	stateful bool,
) {
}

func testSource2[K, V any, Iter ~func(yield func(K, V) bool)](
	t *testing.T,
	sources []func() Iter,
	execpected []V,
	cmpOpt []goCmp.Option,
	stateful bool,
) {
}

func TestSourceStateless[V any, Iter ~func(yield func(V) bool)](
	t *testing.T,
	sources []func() Iter,
	execpected []V,
	cmpOpt []goCmp.Option,
) {
	testSource(t, sources, execpected, cmpOpt, false)
}

func TestSourceStateful[V any, Iter ~func(yield func(V) bool)](
	t *testing.T,
	sources []func() Iter,
	execpected []V,
	cmpOpt []goCmp.Option,
) {
	testSource(t, sources, execpected, cmpOpt, true)
}

func TestSourceStateless2[K, V any, Iter ~func(yield func(K, V) bool)](
	t *testing.T,
	sources []func() Iter,
	execpected []V,
	cmpOpt []goCmp.Option,
) {
	testSource2(t, sources, execpected, cmpOpt, false)
}

func TestSourceStateful2[K, V any, Iter ~func(yield func(K, V) bool)](
	t *testing.T,
	sources []func() Iter,
	execpected []V,
	cmpOpt []goCmp.Option,
) {
	testSource2(t, sources, execpected, cmpOpt, true)
}
