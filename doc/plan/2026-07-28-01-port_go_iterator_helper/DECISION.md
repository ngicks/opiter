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

### D6. Test layout: per category (resolved 2026-07-28)

`source_test.go`, `filter_test.go`, `reducer_test.go` at repo root. Rejected:
one test file per ported file — testhelper suites make entries short enough
that per-file tests would be mostly boilerplate.
