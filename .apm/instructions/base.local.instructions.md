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
