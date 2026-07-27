# Decisions: port go-iterator-helper

## Decided

### D1. Re-author, don't vendor or depend

Port by re-writing functions against the upstream reference rather than copying
files or importing `go-iterator-helper` as a dependency. Rationale: signatures
change (option returns), file organization changes (prefix naming), and opiter
wants zero extra deps. Rejected: re-export shim over an upstream dependency.

### D2. Follow opiter code conventions from `.claude/rules/base.local.md`

Desugared bodies in filters, sugared `range` in sources/reducers; flat root
package with `source-` / `filter-` / `reducer-` file prefixes; loose generic
seq constraints (`Seq ~func(yield func(V) bool)`) as in `basic-kv.go`.

### D3. Scope: `hiter` root only (resolved 2026-07-28)

Port only the flat `hiter` package. Rejected: also porting `hiter/iterable`
and/or `x/exp/xiter` — deferred until a concrete need appears.

### D4. Candidate cut: everything in, except `MergeSort*` (resolved 2026-07-28)

All [?] bundles accepted: `Compact*`, `Unique*`, `Window`, `Zip`, `Try*`,
`Equal*`, `Chan`, `Maps*`, `Cycle*`, `Step*`, `RunningReduce`, `Merge`(Func)(2),
`ReduceGroup`/`InsertReduceGroup`. Explicitly excluded: `MergeSort`,
`MergeSortFunc`, `MergeSortSliceLike`(Func), `ConcatSliceLike` (user: "We need
merge and ReduceGroup but MergeSort").

### D5. Signature rule: options everywhere, `Find*` packs index into KV (resolved 2026-07-28)

`(V, bool)` and zero-on-empty returns become `opt.Option[V]`. `Find*`'s
`(V, int)` becomes `opt.Option[KV[int, V]]`; `TryFind` becomes
`(opt.Option[KV[int, V]], error)`. Boolean predicates stay bool. Rejected:
`(opt.Option[V], int)` with -1 sentinel (half-option), dropping the index
(loses information callers may want).

### D6. Test layout: per implementation file (superseded 2026-07-28)

The original per-category decision was superseded during implementation.
Tests live beside their corresponding implementation as `source-*_test.go`,
`filter-*_test.go`, and `reducer-*_test.go`. Small shared test helpers may use
an explicitly named helper test file.

### D7. Go 1.27 generic fluent methods (resolved 2026-07-28)

Add `Seq` / `Seq2` method counterparts for filters and reducers wherever the
receiver constraints can faithfully express the operation. Methods may declare
their own type parameters because the module targets Go 1.27rc1; no
`GOEXPERIMENT` guard is used. Operations that require strengthening an existing
receiver parameter (`comparable`, ordered, or a concrete `error` position)
remain package functions.

When returning a constructed `Seq[...]` causes a compiler method-set
instantiation cycle, the fluent method returns the corresponding `iter.Seq`
instead.

### D8. Zip representation and constraints (resolved 2026-07-28)

`Zip` and `Zip2` yield `tuple.Tuple2` whose fields are `opt.Option` values,
rather than defining a package-specific zipped pair type. Their inputs use
loose `~func(yield ...)` constraints so named sequence types are accepted.
The package functions return `iter.Seq` because their yielded tuple shape
differs from both inputs and therefore cannot safely preserve either input's
named type. The fluent methods return `Seq[Z]`, using an inferred exact result
type parameter `Z` to avoid recursive method-set instantiation while retaining
the fluent wrapper.
