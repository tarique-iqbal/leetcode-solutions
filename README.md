# LeetCode Solutions

A curated repository of accepted, optimized, and thoroughly documented Go solutions to LeetCode algorithmic problems.

Each solution includes:

* Clean and tested Go implementation
* Dedicated problem explanation
* Time and space complexity analysis
* Unit tests for validation

## Repository Structure

Solutions are organized by difficulty and isolated in problem-specific directories.

```text
.
├── Dockerfile
├── Makefile
├── go.mod
├── Easy
│   ├── binary-tree-inorder-traversal
│   │   ├── README.md
│   │   ├── solution.go
│   │   └── solution_test.go
│   └── squares-of-a-sorted-array
├── Medium
│   ├── binary-tree-level-order-traversal
│   ├── boats-to-save-people
│   ├── clone-graph
│   ├── container-with-most-water
│   ├── decode-string
│   │   ├── README.md
│   │   ├── solution.go
│   │   ├── solution_test.go
│   │   └── recursive/
│   ├── find-first-and-last-position-of-element-in-sorted-array
│   └── network-delay-time
└── README.md
```

## Progress

| Difficulty | Solved |
| ---------- | ------ |
| Easy       | 2      |
| Medium     | 7      |
| Hard       | 0      |
| **Total**  | **9**  |

## Problem Tracker

| Problem                                                 | Difficulty | Language | Notes                             |
| ------------------------------------------------------- | ---------- | -------- | --------------------------------- |
| Binary Tree Inorder Traversal                           | Easy       | Go       | DFS traversal                     |
| Squares of a Sorted Array                               | Easy       | Go       | Two pointers                      |
| Binary Tree Level Order Traversal                       | Medium     | Go       | BFS traversal                     |
| Boats to Save People                                    | Medium     | Go       | Greedy + two pointers             |
| Clone Graph                                             | Medium     | Go       | Graph traversal                   |
| Container With Most Water                               | Medium     | Go       | Two pointers                      |
| Decode String                                           | Medium     | Go       | Stack & recursive implementations |
| Find First and Last Position of Element in Sorted Array | Medium     | Go       | Binary search                     |
| Network Delay Time                                      | Medium     | Go       | Shortest path (Dijkstra)          |

## Testing

Run all tests:

```bash
make test
```

Or using Go directly:

```bash
go test ./...
```

## Documentation Standard

Each problem directory contains a dedicated `README.md` with:

### 1. Problem Statement

* Link to the original LeetCode problem.
* Brief description of the challenge.

### 2. Intuition & Approach

* Core observations.
* Data structures used.
* Algorithm selection rationale.
* Alternative approaches when relevant.

### 3. Complexity Analysis

* **Time Complexity** with justification.
* **Space Complexity** with justification.

### 4. Implementation Notes

* Edge cases handled.
* Language-specific considerations.
* Trade-offs and optimizations.

##  Goals

* Prioritize readability alongside performance.
* Keep explanations concise but educational.
* Build a reusable reference for common algorithmic patterns and data structures.
