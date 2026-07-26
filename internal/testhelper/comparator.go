package testhelper

import (
	"errors"
	"reflect"

	goCmp "github.com/google/go-cmp/cmp"
)

var (
	CompareErrorsIs = goCmp.Comparer(func(e1, e2 error) bool {
		if e1 == nil || e2 == nil {
			return e1 == nil && e2 == nil
		} else {
			// github.com/google/go-cmp/cmp panics when it detects non-symmetric comparator
			// But by its nature, errors.Is is asymmetric.
			return errors.Is(e1, e2) || errors.Is(e2, e1)
		}
	})
	CompareErrorsAs = goCmp.Comparer(func(e1, e2 error) bool {
		if e1 == nil || e2 == nil {
			return e1 == nil && e2 == nil
		}
		// Same as above. Make it symmetric.
		as := func(err, proto error) bool {
			target := reflect.New(reflect.TypeOf(proto)).Interface()
			return errors.As(err, target)
		}
		return as(e1, e2) || as(e2, e1)
	})
	CompareReflectStructFieldSameType = goCmp.Comparer(func(i, j reflect.StructField) bool {
		return i.Name == j.Name && i.PkgPath == j.PkgPath
	})
	CompareReflectValue = goCmp.Comparer(func(i, j reflect.Value) (same bool) {
		defer func() {
			if recover() != nil {
				same = false
			}
		}()
		return i.Interface() == j.Interface()
	})
)
