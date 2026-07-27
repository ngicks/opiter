# Status: port go-iterator-helper

State: **complete and verified** — the selected source, filter, and reducer
helpers, Go 1.27 fluent methods, implementation-aligned tests, examples, and
documentation are implemented.

## Checklist (mirrors PLAN.md steps)

- [x] 0. Resolve open questions 1–4 (2026-07-28)
- [x] 1. Sources (`source-empty.go`, `source-once.go`, `source-range.go`, `source-repeat.go`, `source-chan.go`, `source-maps.go`) + matching tests
- [x] 2. Core filters (`filter-concat.go`, `filter-map.go`, `filter-filter.go`, `filter-limit.go`, `filter-skip.go`) + matching tests
- [x] 3. Shape filters (`filter-flatten.go`, `filter-translate.go`, `filter-tap.go`, `filter-compact.go`, `filter-unique.go`, `filter-window.go`, `filter-step.go`, `filter-zip.go`, `filter-merge.go`, `filter-running-reduce.go`, `filter-cycle.go`) + matching tests
- [x] 4. Reducers (12 `reducer-*.go` files per PLAN inventory) + matching tests
- [x] 5. Short-circuit and infinite-source coverage (ad hoc; no testhelper expansion required)
- [x] 6. Examples + `doc.go` package doc
- [x] 7. Go 1.27 generic `Seq` / `Seq2` method counterparts where receiver constraints permit
- [x] 8. Per-implementation test layout and shape-first generic parameter ordering

## Next action

None.

## Verification

- `go test ./... -count=1`
- `go test -race ./... -count=1`
- `go vet ./...`
- Go 1.27rc1 `gofmt -l .`
- `git diff --check`
