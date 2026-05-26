# 743. Network Delay Time

> Topics: Graph, Shortest Path, Dijkstra, Heap, Priority Queue

## Problem Statement

LeetCode: [743. Network Delay Time](https://leetcode.com/problems/network-delay-time/)

## Problem Overview

You are given a directed weighted graph where:

- `times[i] = [u, v, w]`
  - `u → v` represents a directed edge
  - `w` is the travel time
- There are `n` nodes labeled from `1` to `n`
- A signal starts from node `k`

Return the time it takes for all nodes to receive the signal starting from node `k`.

If it is impossible for all nodes to receive the signal, return `-1`.

## Intuition

This problem asks us to compute the shortest time from one source node to all other nodes.

Since:
- The graph is weighted
- Edge weights are non-negative
- We need minimum distances

this becomes a classic Single Source Shortest Path problem.

The best algorithm for this scenario is Dijkstra’s Algorithm.

## Approach

### 1. Build the Graph

Convert the input edge list into an adjacency list.

Example:

```text
times = [[2,1,1],[2,3,1],[3,4,1]]

Graph:
2 → (1,1), (3,1)
3 → (4,1)
```

### 2. Maintain Shortest Distances

```text
dist[i] = shortest known distance from node k to node i
```

Initialize:
- dist[k] = 0
- all others = ∞

### 3. Use a Min Heap (Priority Queue)

Each heap entry:

```text
(distance, node)
```

This ensures the closest node is processed first.

### 4. Relax Neighboring Edges

For each edge:

```text
newDistance = currentDistance + edgeWeight
```

If better path found:

```text
newDistance < dist[neighbor]
```

update distance and push into heap.

This is called Edge Relaxation.

### 5. Compute Final Answer

If any node is unreachable → return -1

Otherwise:

```text
max(dist[])
```

This represents the farthest node (last node to receive the signal).

## Dry Run

### Input

```text
times = [[2,1,1],[2,3,1],[3,4,1]]
n = 4
k = 2
```

### Final Distances

```text
dist = [∞, 1, 0, 1, 2]
```

### Answer

```text
2
```

## Why Dijkstra Works

- Greedy expansion of nearest node first
- Non-negative weights guarantee correctness
- Once a node is processed, its shortest path is final

## Important Optimization

Skip outdated heap entries:

```text
if currentDistance > dist[node]
```

## Complexity Analysis

### Time Complexity
O((V + E) log V)

### Space Complexity
O(V + E)

## Key Takeaways

- This is a shortest path problem
- Use Dijkstra with min heap
- Answer is maximum of shortest distances
- Return -1 if unreachable

## Notes

- Classic Dijkstra shortest path problem
- Always use min-heap for optimal performance
