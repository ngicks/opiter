# Status: port go-iterator-helper

State: **ready to implement** — plan finalized, all open questions resolved
(DECISION.md D1–D6).

## Checklist (mirrors PLAN.md steps)

- [x] 0. Resolve open questions 1–4 (2026-07-28)
- [ ] 1. Sources (`source-empty.go`, `source-once.go`, `source-range.go`, `source-repeat.go`, `source-chan.go`, `source-maps.go`) + `source_test.go`
- [ ] 2. Core filters (`filter-concat.go`, `filter-map.go`, `filter-filter.go`, `filter-limit.go`, `filter-skip.go`) + `filter_test.go`
- [ ] 3. Shape filters (`filter-flatten.go`, `filter-translate.go`, `filter-tap.go`, `filter-compact.go`, `filter-unique.go`, `filter-window.go`, `filter-step.go`, `filter-zip.go`, `filter-merge.go`, `filter-running-reduce.go`, `filter-cycle.go`)
- [ ] 4. Reducers (12 `reducer-*.go` files per PLAN inventory) + `reducer_test.go`
- [ ] 5. testhelper expansion (short-circuit / infinite-source contracts) — only if needed
- [ ] 6. Examples + `doc.go` package doc

## Next action

Start step 1 (sources).
