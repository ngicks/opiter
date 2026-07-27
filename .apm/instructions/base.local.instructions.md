---
description: "Basic instructions for the project"
applyTo: "*"
---

### Rules

#### Naming convention

- To organize large number of files layed flatten, files have prefixes
- Prefixes may be combined

List of prefixres:

- `basic-` for basic things, e.g. `Seq[V]`, `Seq2[K, V]`, `KV[K, V]`
- `source-`: T -> iterator, e.g. Chan
- `filter-`: iterator -> iterator, e.g. Filter, Map
- `reducer-`: e.g. Reduce, Sum

#### Code convention

This module targets Go 1.27rc1 or higher. Generic methods are part of the language at
this version and must be used for `Seq` / `Seq2` method counterparts where a
filter or reducer needs method-specific type parameters. Do not place generic
methods behind a `GOEXPERIMENT` build guard.

Order generic parameters like the standard library's `slices` and `maps`
packages: put the larger inferred shape first, followed by the element types it
determines, even when the earlier constraint refers to later parameters. For
example, write `S ~[]E, E any`, `M ~map[K]V, K comparable, V any`, and
`Seq ~func(yield func(V) bool), V any`.

We might keep using desugared implementation rather than range-seq in filters, according to:

https://go-review.googlesource.com/c/go/+/745440

> // Using the desugared implementations instead of 'range seq' runs
> // about 2x faster than the code in the comments based on range loops.
> // The cost of range loops comes from the compiler heap-allocating a
> // control variable, and additional checks of well-behavedness w.r.t.
> // concurrency, panics, stopping when yield returns false, etc.
> // There's no value to these "middleware" iterators repeating these
> // checks: the user's range loop (the ultimate consumer) will keep
> // the underlying iterator 'seq' (the ultimate producer) honest.

Keep sugared range-over-func where:

- sources
- reducers

because they are "ultimate producer/consumer".

#### Testing

- Use `internal/testhelper/`
- If it does not suffice, expand it.
- Keep tests in implementation-aligned `*_test.go` files; do not centralize a
  category's tests in one large test file.
