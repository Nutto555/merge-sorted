# merge-sorted

A Go function that merges three already-sorted integer collections into one ascending slice,
without calling any sort function.

```go
func Merge(collection1, collection2, collection3 []int) []int
```

| Input | Order it arrives in |
|---|---|
| `collection1` | max to min (descending) |
| `collection2` | min to max (ascending) |
| `collection3` | min to max (ascending) |

```go
merge.Merge([]int{9, 7, 3}, []int{1, 4, 8}, []int{2, 5})
// [1 2 3 4 5 7 8 9]
```

## Setup

You need Go 1.22 or newer. The module has no third-party dependencies, so there is nothing to
install.

```sh
git clone https://github.com/Nutto555/merge-sorted.git
cd merge-sorted
go mod download     # reports nothing to download
```

## Run the code

```sh
go run ./cmd/merge
```

```
collection_1 (max to min): [9 7 3]
collection_2 (min to max): [1 4 8 10]
collection_3 (min to max): [2 5]
merged (ascending):        [1 2 3 4 5 7 8 9 10]
```

Pass your own collections as comma-separated integers:

```sh
go run ./cmd/merge -c1 "20,10,-5" -c2 "-7,0,30" -c3 "1,1,8"
```

## Run the unit tests

```sh
go test ./... -v
```

`make test` does the same with the race detector, and `make cover` reports coverage (100% of
`merge.go`).

## How it works

All three inputs are already sorted, so the smallest value that has not been used yet is
always at the front of one of them. `collection1` is in the opposite order, so its "front" is
its last element and it is read backwards.

The function keeps one position per collection. On every step it compares the three values at
those positions, appends the smallest to the result, and moves that one position on. A
collection that is used up is skipped. This is the merge step of merge sort, with three inputs
instead of two.

| | |
|---|---|
| Time | O(n), where n is the total number of values. Each value is appended once, after at most two comparisons |
| Space | O(n) for the result, which is allocated once at its final size. O(1) besides |

Appending everything and sorting would cost O(n log n) and ignore what is known about the
inputs.

## Tests

`merge_test.go` has:

- a table of cases: collections of different lengths; one, two or all three empty; `nil`
  inputs; single elements; duplicates inside one collection and across collections; all values
  identical; collections that do not overlap; negative numbers and zero
- a check that the inputs are not modified
- 500 rounds of generated inputs, checking two properties of the result: it is in ascending
  order, and it holds exactly the values that went in, each as many times as it went in
- a runnable example

The tests do not call a sort function either. The generated inputs are built already in
order, and the "same values" check counts occurrences.

## Assumptions

- "sorted from min(0) to max" is read as "the element at index 0 is the minimum", that is,
  ascending order. Negative numbers are handled like any other value.
- The inputs are trusted to be in the order the brief states, and are not checked. If one is
  not, the result is not sorted.
- `nil` and an empty collection are treated the same. The result is never `nil`.
- Duplicates are kept. The result has exactly as many values as the three inputs together.
- The inputs are not modified.

## Layout

```
merge.go            the Merge function
merge_test.go       unit tests
cmd/merge/main.go   command-line example
```

## AI assistance

This repository was written with Claude Code, Anthropic's coding agent.
