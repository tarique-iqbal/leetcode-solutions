# 133. Clone Graph

> Topics: Graph, BFS, DFS, Hash Table

## Problem Statement

LeetCode: 133. Clone Graph

You are given a reference to a node in a **connected undirected graph**.

Each node contains:

```text
Val        -> unique integer value
Neighbors  -> list of adjacent nodes
```

Return a **deep copy (clone)** of the graph.

A deep copy means:

* Every node must be recreated.
* Every edge relationship must be recreated.
* No node in the cloned graph should share memory with the original graph.

---

## Problem Overview

Given a graph like:

```text
1 -- 2
|    |
4 -- 3
```

We need to create:

```text
1' -- 2'
|      |
4' -- 3'
```

where:

```text
1 != 1'
2 != 2'
3 != 3'
4 != 4'
```

but the structure remains exactly the same.

---

## Key Challenge

Graphs can contain:

* Cycles
* Multiple paths to the same node

Example:

```text
1 -> 2
^    |
|____|
```

If we naively clone neighbors recursively or iteratively without tracking nodes:

```text
1 clones 2
2 clones 1
1 clones 2
2 clones 1
...
```

we enter an infinite loop.

Therefore we must remember:

```text
Original Node -> Cloned Node
```

using a hash map.

---

## Intuition

To clone a graph we need two things:

### 1. Create Every Node Exactly Once

Whenever we encounter a node for the first time:

```go
visited[node] = &Node{Val: node.Val}
```

This guarantees:

* one clone per original node
* no duplicate nodes

---

### 2. Rebuild Neighbor Connections

Suppose:

```text
1 -> [2,4]
```

After cloning:

```text
1' -> [2',4']
```

Whenever we process an edge:

```text
current ---- neighbor
```

we connect:

```text
clone(current) ---- clone(neighbor)
```

---

## Approach (BFS)

### Step 1: Handle Empty Graph

If no node exists:

```go
if node == nil {
    return nil
}
```

---

### Step 2: Create Clone of Starting Node

```go
visited := make(map[*Node]*Node)

visited[node] = &Node{
    Val: node.Val,
}
```

The map stores:

```text
Original -> Clone
```

Example:

```text
1 -> 1'
```

---

### Step 3: Start BFS Traversal

```go
queue := []*Node{node}
```

We traverse the graph level by level.

---

### Step 4: Process Current Node

```go
cur := queue[0]
queue = queue[1:]
```

For every neighbor:

```go
for _, nei := range cur.Neighbors
```

---

### Step 5: Clone Unvisited Neighbors

If a neighbor has not been cloned:

```go
if _, ok := visited[nei]; !ok {
    visited[nei] = &Node{
        Val: nei.Val,
    }

    queue = append(queue, nei)
}
```

This ensures each node is cloned only once.

---

### Step 6: Recreate the Edge

Original graph:

```text
cur ---- nei
```

Cloned graph:

```text
clone(cur) ---- clone(nei)
```

```go
visited[cur].Neighbors = append(
    visited[cur].Neighbors,
    visited[nei],
)
```

This rebuilds the graph structure.

---

### Step 7: Return Cloned Start Node

```go
return visited[node]
```

Since every reachable node was processed, the entire graph is cloned.

---

## Dry Run

### Input Graph

```text
1 -- 2
|    |
4 -- 3
```

---

### Initial State

```text
visited:
1 -> 1'

queue:
[1]
```

---

### Process Node 1

Neighbors:

```text
2, 4
```

Create clones:

```text
2 -> 2'
4 -> 4'
```

Connect:

```text
1' -> [2',4']
```

Queue:

```text
[2,4]
```

---

### Process Node 2

Neighbors:

```text
1,3
```

Clone:

```text
3 -> 3'
```

Connect:

```text
2' -> [1',3']
```

Queue:

```text
[4,3]
```

---

### Process Node 4

Neighbors:

```text
1,3
```

Connect:

```text
4' -> [1',3']
```

Queue:

```text
[3]
```

---

### Process Node 3

Neighbors:

```text
2,4
```

Connect:

```text
3' -> [2',4']
```

Queue:

```text
[]
```

---

### Final Cloned Graph

```text
1' -- 2'
|      |
4' -- 3'
```

All nodes are newly allocated.

---

## Why the Hash Map is Necessary

Consider:

```text
1 -- 2
|    |
4 -- 3
```

Node 3 can be reached through:

```text
1 -> 2 -> 3
```

and

```text
1 -> 4 -> 3
```

Without a map:

```text
3 could be cloned twice
```

resulting in:

```text
3'
3''
```

which breaks the graph structure.

The map guarantees:

```text
One original node
      ↓
One cloned node
```

throughout the traversal.

---

## Deep Copy Verification

Original:

```go
fmt.Printf("%p\n", n1)
```

Clone:

```go
fmt.Printf("%p\n", clone)
```

Addresses are different:

```text
Original Node 1 Address: 0xc000...
Cloned   Node 1 Address: 0xc000...
```

and

```go
n1 == clone
```

returns:

```text
false
```

proving that the clone is a separate graph in memory.

---

## Correctness Argument

For every node:

1. It is cloned exactly once when first discovered.
2. Every original edge is recreated between cloned nodes.
3. BFS visits every reachable node.
4. The hash map prevents duplicate cloning.

Therefore:

* every node exists in the clone,
* every edge exists in the clone,
* no original node is reused,

which produces a valid deep copy of the graph.

---

## Complexity Analysis

### Time Complexity

```text
O(V + E)
```

where:

* V = number of vertices
* E = number of edges

Each node is visited once and each edge is processed once.

---

### Space Complexity

```text
O(V)
```

for:

* visited map
* BFS queue

---

## BFS vs DFS

Both solutions are valid.

### BFS

```text
Queue
Level-order traversal
```

### DFS

```text
Recursion / Stack
Depth-first traversal
```

Both have:

```text
Time:  O(V + E)
Space: O(V)
```

The choice is mostly implementation preference.

---

## Key Takeaways

* Graph cloning is fundamentally a traversal problem.
* Use a hash map:

```text
Original Node -> Cloned Node
```

* The map prevents:

  * infinite loops
  * duplicate nodes
  * broken graph structure
* BFS and DFS are equally valid.
* Complexity is:

```text
Time  : O(V + E)
Space : O(V)
```

## Notes

* Classic Graph Traversal + Hash Map problem.
* One of the most important graph-copying patterns.
* Core lesson:

```text
When copying graphs with cycles,
always maintain:

Original -> Clone
mapping.
```
