# Port selective set of go-iterator-helper into opiter

Port a curated, option-aware subset of `github.com/ngicks/go-iterator-helper/hiter`
into `github.com/ngicks/opiter`, reorganized under opiter's flat prefix-named file
layout and verified with `internal/testhelper`.

## Goal / success criteria

- opiter exposes the selected helpers as a flat root package, named per the
  `basic-` / `source-` / `filter-` / `reducer-` prefix convention
  (`.claude/rules/base.local.md`).
- Helpers that upstream expresses as `(V, bool)` / zero-value-on-empty return
  `opt.Option[V]` (`github.com/ngicks/option/opt`) — this is opiter's purpose
  ("option-aware iterator helpers", `doc.go`).
- Every ported source/filter/reducer passes the corresponding
  `internal/testhelper` contract suite (`TestSourceStateless`/`Stateful`,
  `TestFilterStateless`/`Stateful` (+`2`, `2To1`, `1To2`), `TestReducer`(2)).
- `go vet` and `gofmt` clean; no new dependencies beyond
  `github.com/ngicks/option` (already required).

## Scope and non-goals

- **In scope**: the upstream `hiter` root package, selectively (see inventory).
  Decided (D3): root only.
- **Non-goals**: upstream subpackages — `hiter/iterable`, `hiter/errbox`,
  `hiter/async`, `hiter/mapper`, `hiter/tee`, the stdlib-binding packages
  (`bufioiter`, `containeriter`, `cryptoiter`, `databaseiter`, `encodingiter`,
  `ioiter`, `iterreader`, `mathiter`, `reflectiter`, `stringsiter`), and
  `x/exp/xiter`. Add later on demand.
- `MergeSort*` / `ConcatSliceLike` are explicitly excluded (D4) even though
  plain `Merge` is in.
- Not a fork: code is re-authored to opiter conventions (desugared filter
  bodies, option returns), not copied verbatim. Upstream is the same author, so
  licensing is a non-issue.

## Context

- Source repo is cloned locally at
  `/home/watage/gitrepo/github.com/ngicks/go-iterator-helper`.
- opiter today has only the `basic-` layer: `basic-iter.go` (`Seq`, `Seq2`
  wrappers), `basic-func.go` (`FuncIterable`(2), `WrapFunc`(2)),
  `basic-interface.go` (`Iterable`(2), `IntoIterable`(2)), `basic-kv.go`
  (`KV`, `PackKV`, `Values2`, `AppendSeq2`, `Collect2`, `AssembleKV`,
  `DisassembleKV`, `KVPairs`).
- `internal/testhelper` provides contract suites for the three categories and
  is the mandated test vehicle ("Use internal/testhelper; if it does not
  suffice, expand it").
- Code convention (from `.claude/rules/base.local.md`): filters use desugared
  `seq(func(v V) bool { ... })` bodies (per go.dev CL 745440 perf note);
  sources and reducers keep sugared `for ... range seq`.
- Generic signatures follow the existing loose style where useful:
  `Seq ~func(yield func(V) bool)` type params (as in `basic-kv.go`), so named
  seq types are accepted without conversion.
- Go 1.27 generic methods provide `Seq` / `Seq2` counterparts where receiver
  constraints can express the operation. Methods that construct a recursively
  nested `Seq` method set may return `iter.Seq` instead.

## Port inventory (final selection, D4)

Legend: **[P]** port, **[–]** skip (reason noted).

### Sources (`source-`)

| Upstream (file: funcs) | Verdict | Target file |
| --- | --- | --- |
| `empty.go`: `Empty`, `Empty2` | [P] | `source-empty.go` |
| `once.go`: `Once`, `Once2` | [P] | `source-once.go` |
| `range.go`: `Range`, `RangeInclusive` | [P] | `source-range.go` |
| `repeat.go`: `Repeat`, `Repeat2`, `RepeatFunc`, `RepeatFunc2` | [P] | `source-repeat.go` |
| `chan.go`: `Chan` (`ChanSend` is a sink — skip) | [P] | `source-chan.go` |
| `maps.go`: `MapsKeys`, `MapsSorted`, `MapsSortedFunc`, `MapsOverlay` | [P] | `source-maps.go` |
| `atter.go`, `nexter.go`, `iterable.go`, `permutation.go` | [–] niche / interface-bound | |

### Filters (`filter-`)

| Upstream | Verdict | Target file |
| --- | --- | --- |
| `basic_adapter.go`: `Concat`(2) | [P] | `filter-concat.go` |
| `basic_adapter.go`: `Map`(2) | [P] | `filter-map.go` |
| `basic_adapter.go`: `Filter`(2) | [P] | `filter-filter.go` |
| `basic_adapter.go`: `Limit`(2) + `limit.go`: `LimitUntil`(2), `LimitAfter`(2) | [P] | `filter-limit.go` |
| `skip.go`: `Skip`(2), `SkipLast`(2), `SkipWhile`(2) | [P] | `filter-skip.go` |
| `flatten.go`: `Flatten`, `FlattenSeq`(2), `FlattenF`, `FlattenL`, `FlattenSeqF`, `FlattenSeqL` | [P] | `filter-flatten.go` |
| `translate.go`: `Enumerate`, `Pairs`(2), `Omit`(2), `OmitF`, `OmitL`, `Unify`, `Divide`, `Transpose` | [P] | `filter-translate.go` |
| `tap.go`: `Tap`(2), `TapLast`(2) | [P] | `filter-tap.go` |
| `compact.go`: `Compact`(Func)(2) | [P] | `filter-compact.go` |
| `unique.go`: `Unique`(Func)(2) | [P] (buffering — stateless suite only) | `filter-unique.go` |
| `window.go`: `Window`, `WindowSeq` | [P] | `filter-window.go` |
| `step.go`: `Step`, `StepBy` | [P] | `filter-step.go` |
| `cycle.go`: `Cycle`(2), `CycleBuffered`(2) | [P] (infinite — needs ad-hoc tests) | `filter-cycle.go` |
| `basic_adapter.go`: `Zip`(2) | [P] | `filter-zip.go` |
| `basic_adapter.go`: `Merge`(Func)(2) | [P] | `filter-merge.go` |
| `running_reduce.go`: `RunningReduce` | [P] | `filter-running-reduce.go` |
| `merge_sort.go`: `MergeSort`(Func), `MergeSortSliceLike`(Func), `ConcatSliceLike` | [–] excluded by D4 | |
| `alternate.go`, `decorate.go`, `replace.go`, `check_each.go`, `group_id.go`, `assert.go` | [–] niche; add on demand | |

### Reducers (`reducer-`)

| Upstream | Verdict | Target file | Option-aware change (D5) |
| --- | --- | --- | --- |
| `basic_adapter.go`: `Reduce`(2) | [P] | `reducer-reduce.go` | — |
| `sum.go`: `Sum`, `SumOf` | [P] | `reducer-sum.go` | — (zero is the additive identity) |
| `min_max.go`: `Min`(Func), `Max`(Func) | [P] | `reducer-min-max.go` | `V` → `opt.Option[V]` |
| `first_last.go`: `First`(2), `Last`(2) | [P] | `reducer-first-last.go` | `(V, bool)` → `opt.Option[V]` |
| `nth.go`: `Nth`(2) | [P] | `reducer-nth.go` | `(V, bool)` → `opt.Option[V]` |
| `find.go`: `Find`(Func)(2), `FindLast`(Func)(2) | [P] | `reducer-find.go` | `(V, int)` → `opt.Option[KV[int, V]]` |
| `find.go`: `Contains`(Func)(2) | [P] | `reducer-contains.go` | — (bool) |
| `every_any.go`: `Every`(2), `Any`(2) | [P] | `reducer-every-any.go` | — |
| `for_each.go`: `ForEach`(2), `Discard`(2) | [P] | `reducer-for-each.go` | — |
| `for_each.go`: `TryFind`, `TryForEach`, `TryReduce`, `TryCollect`, `TryAppendSeq`, `TryMapsCollect`, `TryInsert` | [P] | `reducer-try.go` | `TryFind` `(v, idx, err)` → `(opt.Option[KV[int, V]], error)` |
| `basic_adapter.go`: `Equal`(Func)(2) | [P] | `reducer-equal.go` | — |
| `reduce_group.go`: `ReduceGroup`, `InsertReduceGroup` | [P] | `reducer-reduce-group.go` | — |
| `bytes.go`, `for_each.go` (`ForEachGo`(2)), `kv.go` (`ToKeyValue`/`FromKeyValue`) | [–] niche / stdlib-bound / already covered by `basic-kv.go` | |

## Approach

- Re-author function by function against upstream as reference; do not vendor
  files wholesale. Rewrite filter bodies desugared; keep sources/reducers
  sugared.
- **Signature rule (D5)**: every upstream `(V, bool)` and zero-value-on-empty
  return becomes `opt.Option[V]`; `Find*`-style `(V, int)` returns become
  `opt.Option[KV[int, V]]` (index in `K`, value in `V`); `TryFind` becomes
  `(opt.Option[KV[int, V]], error)`. Boolean predicates (`Contains`, `Every`,
  `Any`, `Equal`) stay bool.
- Keep upstream parameter order (`fn` first, `seq` last).
- Rejected alternative: importing `go-iterator-helper` as a dependency and
  re-exporting — rejected because opiter changes signatures and file
  organization, and wants zero extra deps.
- Rejected alternative: porting everything — contradicts "selective"; the
  stdlib-binding subpackages drag in deps and don't benefit from options.

## Implementation steps

Each step compiles and passes tests independently; files as in the inventory.
Tests are laid out per implementation file (D6, superseded): for example,
`source-range_test.go`, `filter-limit_test.go`, and `reducer-find_test.go`.

1. **Sources**: `source-empty.go`, `source-once.go`, `source-range.go`,
   `source-repeat.go`, `source-chan.go`, `source-maps.go`; matching
   `source-*_test.go` files
   via `TestSourceStateless`(2) / `TestSourceStateful`(2) (`Chan` is stateful;
   `Maps*` iteration order — sorted variants are deterministic, `MapsKeys` /
   `MapsOverlay` need order-insensitive comparison via `cmpOpt`).
2. **Core filters**: `filter-concat.go`, `filter-map.go`, `filter-filter.go`,
   `filter-limit.go`, `filter-skip.go`; matching `filter-*_test.go` files via
   `TestFilterStateless`/`Stateful` (+`2`, `2To1`, `1To2`).
3. **Shape filters**: `filter-flatten.go`, `filter-translate.go`,
   `filter-tap.go`, `filter-compact.go`, `filter-unique.go`,
   `filter-window.go`, `filter-step.go`, `filter-zip.go`, `filter-merge.go`,
   `filter-running-reduce.go`, `filter-cycle.go`. Buffering filters
   (`SkipLast`, `TapLast`, `Unique`, `Window`) run the stateless suite only,
   like `reverse` in `internal/testhelper/testhelper_test.go`; `Cycle` is
   infinite — test via `Limit(Cycle(...))` ad hoc.
4. **Reducers**: `reducer-reduce.go`, `reducer-sum.go`, `reducer-min-max.go`,
   `reducer-first-last.go`, `reducer-nth.go`, `reducer-find.go`,
   `reducer-contains.go`, `reducer-every-any.go`, `reducer-for-each.go`,
   `reducer-try.go`, `reducer-equal.go`, `reducer-reduce-group.go`; matching
   `reducer-*_test.go` files via `TestReducer`(2).
5. **testhelper expansion if needed**: e.g. a helper asserting a reducer
   short-circuits (stops consuming after decision) for `First`/`Find`/`Any`;
   an infinite-source guard for `Cycle`.
6. **Examples/docs**: port upstream example tests for ported symbols only,
   adapted to option returns (`example_*_test.go` at root). Update `doc.go`
   package doc with a category overview.
7. **Fluent methods**: add Go 1.27 generic `Seq` / `Seq2` method counterparts
   for filters and reducers wherever receiver constraints permit them. Keep
   package functions for operations that need stronger receiver constraints
   or specialized sequence shapes.

## Testing and verification

- `go test ./... -count=1`, `go vet ./...`, `gofmt -l .` after each step.
- Filters additionally exercised through the one-shot/no-reuse contracts the
  testhelper suites enforce; reducers through the single-use-iterator rule.
- Skills `go-check-outdated-patterns` and `go-review-checklist` after each
  editing step (session rule).

## Risks

- **Suite mismatch**: some helpers are neither pure source/filter/reducer
  (`Chan` needs a feeding goroutine; `Cycle` never terminates; `ForEachGo` was
  skipped for this reason) — testhelper contracts don't apply cleanly; handle
  via step 5 or ad-hoc tests.
- **Map iteration order**: `MapsKeys` / `MapsOverlay` yield in map order;
  contract suites compare ordered slices, so these need sorted expectations or
  a set-comparison `cmpOpt`.
- **Upstream drift**: upstream keeps evolving; the plan pins semantics to the
  local clone as of 2026-07-28 (`go-iterator-helper` HEAD on disk).

## Open questions

None — all resolved 2026-07-28; see DECISION.md D3–D6.
