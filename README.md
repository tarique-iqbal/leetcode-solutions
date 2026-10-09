# LeetCode Solutions

A curated collection of accepted, documented Go solutions to LeetCode problems, each with tests and a write-up of the approach and complexity.

Each problem directory includes:

* A Go implementation (`solution.go`)
* Table-driven unit tests (`solution_test.go`)
* A write-up with intuition, approach, and time/space complexity (`README.md`)

## Repository Structure

Solutions are organized by difficulty, with one directory per problem named after its LeetCode slug.

```text
leetcode-solutions/
├── Dockerfile
├── Makefile
├── go.mod
├── Easy/
│   ├── binary-tree-inorder-traversal/
│   │   ├── README.md
│   │   ├── solution.go
│   │   └── solution_test.go
│   ├── squares-of-a-sorted-array/
│   └── ...
├── Medium/
│   ├── boats-to-save-people/
│   ├── decode-string/
│   │   └── recursive/          # alternative implementation
│   ├── network-delay-time/
│   └── ...
└── README.md
```

Every problem is its own `package main`, so problems are fully independent of each other.

## Progress

| Difficulty | Solved |
| ---------- | ------ |
| Easy       | 2      |
| Medium     | 7      |
| Hard       | 0      |
| **Total**  | **9**  |

## Problem Tracker

| Problem                                                                                                         | Difficulty | Notes                             |
| --------------------------------------------------------------------------------------------------------------- | ---------- | --------------------------------- |
| [Binary Tree Inorder Traversal](Easy/binary-tree-inorder-traversal)                                             | Easy       | DFS traversal                     |
| [Squares of a Sorted Array](Easy/squares-of-a-sorted-array)                                                     | Easy       | Two pointers                      |
| [Binary Tree Level Order Traversal](Medium/binary-tree-level-order-traversal)                                   | Medium     | BFS traversal                     |
| [Boats to Save People](Medium/boats-to-save-people)                                                             | Medium     | Greedy + two pointers             |
| [Clone Graph](Medium/clone-graph)                                                                               | Medium     | Graph traversal                   |
| [Container With Most Water](Medium/container-with-most-water)                                                   | Medium     | Two pointers                      |
| [Decode String](Medium/decode-string)                                                                           | Medium     | Stack & recursive implementations |
| [Find First and Last Position of Element in Sorted Array](Medium/find-first-and-last-position-of-element-in-sorted-array) | Medium     | Binary search                     |
| [Network Delay Time](Medium/network-delay-time)                                                                 | Medium     | Shortest path (Dijkstra)          |

## Running Tests

Requires Go 1.24+ (see `go.mod`).

```bash
go test ./...                                            # all problems
go test ./Medium/decode-string/...                       # one problem, including subpackages
go test ./Medium/decode-string -run TestDecodeString -v  # a single test
```

Or run in a reproducible Docker environment:

```bash
make build        # build the image
make test         # run all tests in a clean container
make test-local   # run tests against your local working tree
make help         # list all targets
```

## Documentation Standard

Each problem's `README.md` follows the same structure:

1. **Problem Statement**: title as `# <number>. <Title>`, a `> Topics:` line, a link to the LeetCode problem, and a brief description with examples.
2. **Intuition & Approach**: core observations, data structures used, why the algorithm was chosen, and alternatives when relevant.
3. **Complexity Analysis**: time and space complexity, each with justification.
4. **Implementation Notes**: edge cases, Go-specific considerations, and trade-offs.

## Adding a Solution

1. Create `Easy/<slug>/` or `Medium/<slug>/` (or `Hard/<slug>/`) with `solution.go`, `solution_test.go`, and `README.md`.
2. Verify with `go test ./<Difficulty>/<slug>/...`.
3. Update the **Progress** counts and add a row to the **Problem Tracker** above.

## Goals

* Prioritize readability alongside performance.
* Keep explanations concise but educational.
* Build a reusable reference for common algorithmic patterns and data structures.
